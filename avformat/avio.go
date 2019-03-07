package avformat

// #include <avio.h>
import "C"
import (
	"io/ioutil"
	"unsafe"
)

type AVIOContext = C.AVIOContext

// ProtocolName Find protocol name by url
func ProtocolName(url string) (protocol string) {
	return C.GoString(C.avio_find_protocol_name(C.CString(url)))
}

// NewAVIOContext ..
func NewAVIOContext(buffer unsafe.Pointer, bufferSize int) *AVIOContext {
	return (*AVIOContext)(C.avio_alloc_read_context(
		(*C.uchar)(buffer),
		C.int(bufferSize),
	))
}

// SetAVFormatAVIOContext ..
func (c *Context) SetAVFormatAVIOContext(io *AVIOContext) {
	C.set_avfmt_avio((*C.AVFormatContext)(c), (*C.AVIOContext)(io))
}

func Input() {
	data, _ := ioutil.ReadFile("./test.mp3")
	C.input_buffer((*C.uchar)(unsafe.Pointer(&data[0])), C.int(len(data)))
}
