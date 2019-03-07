#include <libavcodec/avcodec.h>

int get_id_from_codec(AVCodec *codec) { return codec->id; }
