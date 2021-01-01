#include <libavcodec/avcodec.h>
#include <libavutil/avutil.h>
#include <libavutil/dict.h>
#include <libavutil/frame.h>
#include <libavutil/dict.h>

enum AVSampleFormat get_sample_format(AVCodecContext *dec_ctx) {
  return dec_ctx->sample_fmt;
}

int get_channels(AVCodecContext *dec_ctx) { return dec_ctx->channels; }
