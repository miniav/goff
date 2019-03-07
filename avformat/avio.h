#include <libavformat/avformat.h>
#include <libavutil/dict.h>
#include <libavutil/log.h>
#include <pthread.h>
#include <stdbool.h>
#include <stdint.h>
#include <stdio.h>

typedef struct buffer_data {
  uint8_t *ptr;
  size_t size;
} buffer_data;

pthread_mutex_t buffer_in_mutex  = PTHREAD_MUTEX_INITIALIZER;
pthread_mutex_t buffer_out_mutex = PTHREAD_MUTEX_INITIALIZER;

buffer_data buffer_in, buffer_out;

int read_packet(void *opaque, uint8_t *buf, int buf_size) {
  printf("read packet: %d %lu \n", buf_size, buffer_in.size);
  pthread_mutex_lock(&buffer_in_mutex);
  int out;
  if (buffer_in.size <= buf_size) {
    memcpy(buf, buffer_in.ptr, buffer_in.size);
    buffer_in.size -= buffer_in.size;
    out = buffer_in.size;
  } else {
    memcpy(buf, buffer_in.ptr, buf_size);
    buffer_in.size -= buf_size;
    out = buf_size;
  }
  pthread_mutex_unlock(&buffer_in_mutex);
  return out;
}

int input_buffer(uint8_t *buf, int buf_size) {
  pthread_mutex_lock(&buffer_in_mutex);
  buffer_in.ptr = malloc(buf_size * sizeof(uint8_t));
  memcpy(buffer_in.ptr, buf, buf_size);
  buffer_in.size += buf_size;
  printf("input buffer: %d %lu\n", buf_size, buffer_in.size);
  pthread_mutex_unlock(&buffer_in_mutex);
  return buffer_in.size;
}

AVIOContext *avio_alloc_read_context(unsigned char *buffer, int buffer_size) {
  av_log_set_level(AV_LOG_TRACE);
  return avio_alloc_context(buffer, buffer_size, 0, NULL, &read_packet, NULL,
                            NULL);
}

void set_avfmt_avio(AVFormatContext *fmt_ctx, AVIOContext *avio_ctx) {
  fmt_ctx->pb = avio_ctx;
}

// void decode(AVCodecParserContext *s, AVCodecContext *ctx, uint8_t *buf,
//             int buf_size, int *used, int *data_size, uint8_t *data) {
//   AVPacket *pkt = av_packet_alloc();
//   *used = av_parser_parse2(s, ctx, &pkt->data, &pkt->size, buf, buf_size,
//                            AV_NOPTS_VALUE, AV_NOPTS_VALUE, 0);
//   avcodec_send_packet(ctx, pkt);
//   AVFrame *frame = av_frame_alloc();
//   avcodec_receive_frame(ctx, frame);
//   int bytes_per_sample = av_get_bytes_per_sample(ctx->sample_fmt);
//   *data_size = frame->nb_samples * ctx->channels * bytes_per_sample;
//   data = (uint8_t *)malloc(data_size * sizeof(uint8_t));
//   for (int i = 0; i < frame->nb_samples; i++)
//     for (int ch = 0; ch < ctx->channels; ch++) {
//       memcpy(data + i * ctx->channels * bytes_per_sample +
//                  ch * bytes_per_sample,
//              frame->data[ch] + bytes_per_sample * i, bytes_per_sample);
//     }
// }
