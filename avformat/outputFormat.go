package avformat

// #include <libavformat/avformat.h>
import "C"

type OutputFormat C.AVOutputFormat

// RegisterOutputFormat ..
// func (o *OutputFormat) RegisterOutputFormat() {
// 	C.av_register_output_format((*C.AVOutputFormat)(unsafe.Pointer(o)))
// }
