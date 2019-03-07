#include <libavutil/frame.h>

void frame_to_bytes(AVFrame *frame, int channels, int bytes_per_sample, uint8_t *data) {
  for (int i = 0; i < frame->nb_samples; i++)
    for (int ch = 0; ch < channels; ch++) {
      memcpy(data + i * channels * bytes_per_sample +
                 ch * bytes_per_sample,
             frame->data[ch] + bytes_per_sample * i, bytes_per_sample);
    }
}

int frame_samples(AVFrame *frame) { return frame->nb_samples; }
