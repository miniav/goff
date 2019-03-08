package avcodec

// #include <context.h>
import "C"
import (
	"errors"
	"unsafe"

	"github.com/tosone/goff/avutil"
)

type Context C.AVCodecContext

// NewContext ..
func NewContext(codec *Codec) *Context {
	return (*Context)(C.avcodec_alloc_context3((*C.AVCodec)(codec)))
}

// AddParameters ..
func (c *Context) AddParameters(para *Parameters) (err error) {
	if code := int(C.avcodec_parameters_to_context(
		(*C.AVCodecContext)(c),
		(*C.AVCodecParameters)(para),
	)); code < 0 {
		err = errors.New(avutil.StrError(code))
	}
	return
}

// Open2 ..
func (c *Context) Open2(codec *Codec, dict *avutil.AVDictionary) (err error) {
	var c_dict = (*C.AVDictionary)(unsafe.Pointer(dict))
	var d_dict = (**C.AVDictionary)(C.malloc(C.ulong(unsafe.Sizeof(c_dict))))
	*d_dict = c_dict
	if code := int(C.avcodec_open2(
		(*C.AVCodecContext)(c),
		(*C.AVCodec)(codec),
		d_dict,
	)); code < 0 {
		err = errors.New(avutil.StrError(code))
	}
	defer C.free(unsafe.Pointer(d_dict))
	return
}

// ReceiveFrame ..
func (c *Context) ReceiveFrame(frame *avutil.Frame) (err error) {
	if code := int(C.avcodec_receive_frame(
		(*C.AVCodecContext)(unsafe.Pointer(c)),
		(*C.AVFrame)(unsafe.Pointer(frame)),
	)); code < 0 {
		err = avutil.Error(code)
	}
	return
}

// SendPacket ..
func (c *Context) SendPacket(pkt *Packet) (err error) {
	if code := int(C.avcodec_send_packet(
		(*C.AVCodecContext)(unsafe.Pointer(c)),
		(*C.AVPacket)(unsafe.Pointer(pkt)),
	)); code < 0 {
		err = avutil.Error(code)
	}
	return
}

// SampleFmt ..
func (c *Context) SampleFmt() avutil.SampleFmt {
	return avutil.SampleFmt(C.get_sample_format((*C.AVCodecContext)(unsafe.Pointer(c))))
}

func (c *Context) Channels() int {
	return int(C.get_channels((*C.AVCodecContext)(unsafe.Pointer(c))))
}
