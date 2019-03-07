#include <libavcodec/avcodec.h>
#include <libavformat/avformat.h>
#include <libavformat/avio.h>
#include <libavutil/file.h>
#include <stdbool.h>
#include <stdint.h>
#include <stdio.h>

struct buffer_data {
  uint8_t *ptr;
  size_t size; ///< size left in the buffer
};

#define C_CAST(type, variable) ((type)variable)
#define REINTERPRET_CAST(type, variable) C_CAST(type, variable)
#define STATIC_CAST(type, variable) C_CAST(type, variable)

#define RAW_OUT_ON_PLANAR true

int printError(const char *prefix, int errorCode) {
  if (errorCode == 0) {
    return 0;
  } else {
    const size_t bufsize = 64;
    char buf[bufsize];

    if (av_strerror(errorCode, buf, bufsize) != 0) {
      strcpy(buf, "UNKNOWN_ERROR");
    }
    fprintf(stderr, "%s (%d: %s)\n", prefix, errorCode, buf);

    return errorCode;
  }
}

float getSample(const AVCodecContext *codecCtx, uint8_t *buffer,
                int sampleIndex) {
  int64_t val = 0;
  float ret = 0;
  int sampleSize = av_get_bytes_per_sample(codecCtx->sample_fmt);
  switch (sampleSize) {
  case 1:
    // 8bit samples are always unsigned
    val = REINTERPRET_CAST(uint8_t *, buffer)[sampleIndex];
    // make signed
    val -= 127;
    break;

  case 2:
    val = REINTERPRET_CAST(int16_t *, buffer)[sampleIndex];
    break;

  case 4:
    val = REINTERPRET_CAST(int32_t *, buffer)[sampleIndex];
    break;

  case 8:
    val = REINTERPRET_CAST(int64_t *, buffer)[sampleIndex];
    break;

  default:
    fprintf(stderr, "Invalid sample size %d.\n", sampleSize);
    return 0;
  }

  // Check which data type is in the sample.
  switch (codecCtx->sample_fmt) {
  case AV_SAMPLE_FMT_U8:
  case AV_SAMPLE_FMT_S16:
  case AV_SAMPLE_FMT_S32:
  case AV_SAMPLE_FMT_U8P:
  case AV_SAMPLE_FMT_S16P:
  case AV_SAMPLE_FMT_S32P:
    // integer => Scale to [-1, 1] and convert to float.
    ret = val / STATIC_CAST(float, ((1 << (sampleSize * 8 - 1)) - 1));
    break;

  case AV_SAMPLE_FMT_FLT:
  case AV_SAMPLE_FMT_FLTP:
    // float => reinterpret
    ret = *REINTERPRET_CAST(float *, &val);
    break;

  case AV_SAMPLE_FMT_DBL:
  case AV_SAMPLE_FMT_DBLP:
    // double => reinterpret and then static cast down
    ret = STATIC_CAST(float, *REINTERPRET_CAST(double *, &val));
    break;

  default:
    fprintf(stderr, "Invalid sample format %s.\n",
            av_get_sample_fmt_name(codecCtx->sample_fmt));
    return 0;
  }

  return ret;
}

static int read_packet(void *opaque, uint8_t *buf, int buf_size) {
  struct buffer_data *bd = (struct buffer_data *)opaque;
  buf_size = FFMIN(buf_size, bd->size);
  printf("%d, %d\n", buf_size, bd->size);

  if (!buf_size)
    return AVERROR_EOF;
  printf("ptr:%p size:%zu\n", bd->ptr, bd->size);

  /* copy internal buffer data to buf */
  memcpy(buf, bd->ptr, buf_size);
  bd->ptr += buf_size;
  bd->size -= buf_size;

  return buf_size;
}

int findAudioStream(const AVFormatContext *formatCtx) {
  int audioStreamIndex = -1;
  for (size_t i = 0; i < formatCtx->nb_streams; ++i) {
    // Use the first audio stream we can find.
    // NOTE: There may be more than one, depending on the file.
    if (formatCtx->streams[i]->codecpar->codec_type == AVMEDIA_TYPE_AUDIO) {
      audioStreamIndex = i;
      break;
    }
  }
  return audioStreamIndex;
}

FILE *outFile;

void handleFrame(const AVCodecContext *codecCtx, const AVFrame *frame) {
  if (av_sample_fmt_is_planar(codecCtx->sample_fmt) == 1) {
    // This means that the data of each channel is in its own buffer.
    // => frame->extended_data[i] contains data for the i-th channel.
    for (int s = 0; s < frame->nb_samples; ++s) {
      for (int c = 0; c < codecCtx->channels; ++c) {
        float sample = getSample(codecCtx, frame->extended_data[c], s);
        fwrite(&sample, sizeof(float), 1, outFile);
      }
    }
  } else {
    // This means that the data of each channel is in the same buffer.
    // => frame->extended_data[0] contains data of all channels.
    if (RAW_OUT_ON_PLANAR) {
      fwrite(frame->extended_data[0], 1, frame->linesize[0], outFile);
    } else {
      for (int s = 0; s < frame->nb_samples; ++s) {
        for (int c = 0; c < codecCtx->channels; ++c) {
          float sample = getSample(codecCtx, frame->extended_data[0],
                                   s * codecCtx->channels + c);
          fwrite(&sample, sizeof(float), 1, outFile);
        }
      }
    }
  }
}

int receiveAndHandle(AVCodecContext *codecCtx, AVFrame *frame) {
  int err = 0;
  // Read the packets from the decoder.
  // NOTE: Each packet may generate more than one frame, depending on the codec.
  while ((err = avcodec_receive_frame(codecCtx, frame)) == 0) {
    // Let's handle the frame in a function.
    handleFrame(codecCtx, frame);
    // Free any buffers and reset the fields to default values.
    av_frame_unref(frame);
  }
  return err;
}

int main(int argc, char *argv[]) {
  AVFormatContext *fmt_ctx = NULL;
  AVIOContext *avio_ctx = NULL;
  AVCodecContext *codecCtx = NULL;
  uint8_t *buffer = NULL, *avio_ctx_buffer = NULL;
  size_t buffer_size, avio_ctx_buffer_size = 1024;

  outFile = fopen("test.app.mp3.raw", "w+");
  char *input_filename = NULL;
  int ret = 0;
  struct buffer_data bd = {0};

  if (argc != 2) {
    fprintf(stderr,
            "usage: %s input_file\n"
            "API example program to show how to read from a custom buffer "
            "accessed through AVIOContext.\n",
            argv[0]);
    return 1;
  }
  input_filename = argv[1];

  /* slurp file content into buffer */
  ret = av_file_map(input_filename, &buffer, &buffer_size, 0, NULL);
  if (ret < 0)
    goto end;

  bd.ptr = buffer;
  bd.size = buffer_size;

  if (!(fmt_ctx = avformat_alloc_context())) {
    ret = AVERROR(ENOMEM);
    goto end;
  }

  avio_ctx_buffer = av_malloc(avio_ctx_buffer_size);
  if (!avio_ctx_buffer) {
    ret = AVERROR(ENOMEM);
    goto end;
  }
  avio_ctx = avio_alloc_context(avio_ctx_buffer, avio_ctx_buffer_size, 0, &bd,
                                &read_packet, NULL, NULL);
  if (!avio_ctx) {
    ret = AVERROR(ENOMEM);
    goto end;
  }
  fmt_ctx->pb = avio_ctx;

  ret = avformat_open_input(&fmt_ctx, NULL, NULL, NULL);
  if (ret < 0) {
    printf("%s\n", "got here");
    fprintf(stderr, "Could not open input\n");
    goto end;
  }

  ret = avformat_find_stream_info(fmt_ctx, NULL);
  if (ret < 0) {
    fprintf(stderr, "Could not find stream information\n");
    goto end;
  }

  av_dump_format(fmt_ctx, 0, input_filename, 0);

  int audioStreamIndex = findAudioStream(fmt_ctx);

  av_register_all();

  AVCodec *codec = avcodec_find_decoder(
      fmt_ctx->streams[audioStreamIndex]->codecpar->codec_id);
  if (codec == NULL) {
    fprintf(stderr, "Decoder not found. The codec is not supported.\n");
    goto end;
  }

  codecCtx = avcodec_alloc_context3(codec);
  if (codecCtx == NULL) {
    fprintf(stderr, "Could not allocate a decoding context.\n");
    goto end;
  }

  if ((ret = avcodec_parameters_to_context(
           codecCtx, fmt_ctx->streams[audioStreamIndex]->codecpar)) != 0) {
    goto end;
  }

  codecCtx->request_sample_fmt = av_get_alt_sample_fmt(codecCtx->sample_fmt, 0);

  if ((ret = avcodec_open2(codecCtx, codec, NULL)) != 0) {
    goto end;
  }

  AVFrame *frame = NULL;
  if ((frame = av_frame_alloc()) == NULL) {
    goto end;
  }

  AVPacket packet;
  av_init_packet(&packet);

  while ((ret = av_read_frame(fmt_ctx, &packet)) != AVERROR_EOF) {
    if (ret != 0) {
      printError("Read error.", ret);
      break; // Don't return, so we can clean up nicely.
    }
    if (packet.stream_index != audioStreamIndex) {
      // Free the buffers used by the frame and reset all fields.
      av_packet_unref(&packet);
      continue;
    }
    if ((ret = avcodec_send_packet(codecCtx, &packet)) == 0) {
      av_packet_unref(&packet);
    } else {
      printError("Send error.", ret);
      break;
    }

    if ((ret = receiveAndHandle(codecCtx, frame)) != AVERROR(EAGAIN)) {
      // Not EAGAIN => Something went wrong.
      printError("Receive error.", ret);
      break; // Don't return, so we can clean up nicely.
    }
  }

end:
  avcodec_close(codecCtx);
  avcodec_free_context(&codecCtx);
  avformat_close_input(&fmt_ctx);
  if (avio_ctx) {
    av_freep(&avio_ctx->buffer);
    av_freep(&avio_ctx);
  }
  av_file_unmap(buffer, buffer_size);

  if (ret < 0) {
    fprintf(stderr, "Error occurred: %s\n", av_err2str(ret));
    return 1;
  }

  return 0;
}
