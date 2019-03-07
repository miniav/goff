#include <stdbool.h>
#include <stdint.h>
#include <stdio.h>

#include <libavcodec/avcodec.h>
#include <libavformat/avformat.h>
#include <libavformat/avio.h>
#include <libavutil/file.h>

uint8_t *buffer    = NULL;
size_t buffer_size = 0;

static int read_packet(void *opaque, uint8_t *buf, int buf_size) {
  buf_size = FFMIN(buf_size, buffer_size);
  if (!buf_size) { return AVERROR_EOF; }
  memcpy(buf, buffer, buf_size);
  buffer += buf_size;
  buffer_size -= buf_size;
  return buf_size;
}

#define AUDIO_INBUF_SIZE 9162
#define AUDIO_REFILL_THRESH 4096

AVFormatContext *fmt_ctx = NULL;

void free_fmt_ctx() {
  if (fmt_ctx != NULL) {
    avformat_close_input(&fmt_ctx);
  }
}

AVIOContext *avio_ctx = NULL;

void free_avio_ctx() {
  if (avio_ctx != NULL) {
    av_freep(&avio_ctx->buffer);
    av_freep(&avio_ctx);
  }
}

AVCodec *codec = NULL;

void free_codec() {
  if (codec != NULL) {
    av_free(codec);
  }
}

AVCodecContext *codec_ctx = NULL;

void free_codec_ctx() {
  if (codec_ctx != NULL) {
    av_free(codec_ctx);
  }
}

AVFrame *frame = NULL;

void free_frame() {
  if (frame != NULL) {
    av_frame_free(&frame);
  }
}

int main(int argc, char *argv[]) {
  uint8_t inbuf[AUDIO_INBUF_SIZE + AV_INPUT_BUFFER_PADDING_SIZE];
  FILE *out_file = fopen("test.raw", "w+");
  int ret        = 0;

  if (argc != 2) {
    fprintf(stderr,
            "usage: %s input_file\n"
            "API example program to show how to read from a custom buffer accessed through AVIOContext.\n",
            argv[0]);
    return 1;
  }

  char *input_filename = argv[1];

  FILE *f = fopen(input_filename, "rb");
  if (f == NULL) {
    fprintf(stderr, "Could not open %s\n", input_filename);
    return -1;
  }

  buffer      = inbuf;
  buffer_size = fread(inbuf, 1, AUDIO_INBUF_SIZE, f);

  fmt_ctx = avformat_alloc_context();
  atexit(free_fmt_ctx);
  if (fmt_ctx == NULL) {
    return AVERROR(ENOMEM);
  }

  size_t avio_ctx_buffer_size = 1024;

  uint8_t *avio_ctx_buffer = av_malloc(avio_ctx_buffer_size);
  if (avio_ctx_buffer == NULL) {
    return AVERROR(ENOMEM);
  }

  avio_ctx = avio_alloc_context(avio_ctx_buffer, avio_ctx_buffer_size, 0, NULL, &read_packet, NULL, NULL);
  atexit(free_avio_ctx);
  if (avio_ctx == NULL) {
    return AVERROR(ENOMEM);
  }

  fmt_ctx->pb = avio_ctx;

  if ((ret = avformat_open_input(&fmt_ctx, NULL, NULL, NULL)) < 0) {
    return ret;
  }

  if ((ret = avformat_find_stream_info(fmt_ctx, NULL)) < 0) {
    return ret;
  }

  av_dump_format(fmt_ctx, 0, input_filename, 0);

  int audio_stream_index = -1;
  for (int i = 0; i < fmt_ctx->nb_streams; ++i) {
    if (fmt_ctx->streams[i]->codecpar->codec_type == AVMEDIA_TYPE_AUDIO) {
      audio_stream_index = i;
      break;
    }
  }
  if (audio_stream_index == -1) {
    return -1;
  }

  av_register_all();

  codec = avcodec_find_decoder(fmt_ctx->streams[audio_stream_index]->codecpar->codec_id);
  atexit(free_codec);
  if (codec == NULL) {
    return -1;
  }

  codec_ctx = avcodec_alloc_context3(codec);
  atexit(free_codec_ctx);
  if (codec_ctx == NULL) {
    return -1;
  }

  if ((ret = avcodec_parameters_to_context(codec_ctx, fmt_ctx->streams[audio_stream_index]->codecpar)) < 0) {
    return ret;
  }

  codec_ctx->request_sample_fmt = av_get_alt_sample_fmt(codec_ctx->sample_fmt, 0);

  if ((ret = avcodec_open2(codec_ctx, codec, NULL)) < 0) {
    return ret;
  }

  frame = av_frame_alloc();
  atexit(free_frame);
  if (frame == NULL) {
    return -1;
  }

  AVPacket packet;
  av_init_packet(&packet);

  int bytes_per_sample = av_get_bytes_per_sample(codec_ctx->sample_fmt);

  while ((ret = av_read_frame(fmt_ctx, &packet)) != AVERROR_EOF) {
    if (ret < 0) {
      break;
    }
    if (packet.stream_index != audio_stream_index) {
      av_packet_unref(&packet);
      continue;
    }
    if ((ret = avcodec_send_packet(codec_ctx, &packet)) == 0) {
      av_packet_unref(&packet);
    } else {
      break;
    }

    while ((ret = avcodec_receive_frame(codec_ctx, frame)) == 0) {
      for (int i = 0; i < frame->nb_samples; i++) {
        for (int ch = 0; ch < codec_ctx->channels; ch++) {
          fwrite(frame->data[ch] + bytes_per_sample * i, 1, bytes_per_sample, out_file);
        }
      }
      av_frame_unref(frame);
    }
    if (buffer_size < AUDIO_REFILL_THRESH) {
      memmove(inbuf, buffer, buffer_size);
      buffer  = inbuf;
      int len = fread(buffer + buffer_size, 1, AUDIO_INBUF_SIZE - buffer_size, f);
      if (len > 0) { buffer_size += len; }
    }
  }
  if (ret == AVERROR_EOF) {
    ret = 0;
  }
  return ret;
}
