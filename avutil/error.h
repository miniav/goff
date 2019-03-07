#include <libavutil/common.h>
#include <libavutil/error.h>

int averror(int e) { return AVERROR(e); }
