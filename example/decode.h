#include <libavcodec/avcodec.h>
#include <libavutil/frame.h>
#include <libavutil/mem.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int decode(AVCodecParserContext *s, AVCodecContext *ctx, uint8_t *buf,
           int buf_size, int *used, int *data_size, uint8_t *data) {
  int ret       = 0;
  AVPacket *pkt = av_packet_alloc();
  *used         = av_parser_parse2(s, ctx, &pkt->data, &pkt->size, buf, buf_size, AV_NOPTS_VALUE, AV_NOPTS_VALUE, 0);
  ret           = avcodec_send_packet(ctx, pkt);
  if (ret < 0) {
    printf("%s\n", "got here");
    return ret;
  }
  int bytes_per_sample = av_get_bytes_per_sample(ctx->sample_fmt);
  while (ret >= 0) {
    AVFrame *frame = av_frame_alloc();
    ret            = avcodec_receive_frame(ctx, frame);
    if (ret < 0) {
      return ret;
    }
    int temp_size = frame->nb_samples * ctx->channels * bytes_per_sample;
    uint8_t *temp = (uint8_t *)malloc(temp_size * sizeof(uint8_t));
    for (int i = 0; i < frame->nb_samples; i++) {
      for (int ch = 0; ch < ctx->channels; ch++) {
        memcpy(temp + i * ctx->channels * bytes_per_sample + ch * bytes_per_sample,
               frame->data[ch] + bytes_per_sample * i, bytes_per_sample);
      }
    }
    uint8_t *new_data = (uint8_t *)malloc((*data_size + temp_size) * sizeof(uint8_t));
    memcpy(new_data, data, *data_size);
    memcpy(new_data + *data_size, temp, temp_size);
    *data_size += temp_size;
    data = new_data;
  }
  return 0;
}
