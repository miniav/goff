#include <libavcodec/avcodec.h>

void av_packet_free_personal(AVPacket *pkt) {
  av_packet_free(&pkt);
}

void av_packet_zero(AVPacket *pkt) {
  pkt->data = NULL;
  pkt->size = 0;
}
