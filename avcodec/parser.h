#include <libavcodec/avcodec.h>

int av_parser_parse2_personal(AVCodecParserContext *s, AVCodecContext *avctx, AVPacket *packet, const uint8_t *buf, int buf_size, int64_t pts, int64_t dts, int64_t pos) {
  return av_parser_parse2(s, avctx, &packet->data, &packet->size, buf, buf_size, pts, dts, pos);
}
