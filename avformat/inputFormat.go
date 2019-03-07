package avformat

// #include <libavformat/avformat.h>
import "C"
import "unsafe"

type InputFormat C.AVInputFormat

// RegisterInputFormat ..
func (i *InputFormat) RegisterInputFormat() {
	C.av_register_input_format((*C.AVInputFormat)(unsafe.Pointer(i)))
}
