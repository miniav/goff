#include <libavformat/avformat.h>
#include <libavutil/dict.h>

int find_audio_stream_index(AVFormatContext *c) {
  int audioStream = -1;
  for (int i = 0; i < c->nb_streams; i++) {
    if (c->streams[i]->codecpar->codec_type == AVMEDIA_TYPE_AUDIO) {
      audioStream = i;
      break;
    }
  }
  return audioStream;
}

enum AVCodecID get_audio_stream_codec_id(AVFormatContext *c, int index) {
  return c->streams[index]->codecpar->codec_id;
}
