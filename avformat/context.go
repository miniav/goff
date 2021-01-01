package avformat

// #include <context.h>
import "C"
import (
	"errors"
	"log"
	"unsafe"

	"github.com/tosone/goff/avcodec"
	"github.com/tosone/goff/avutil"
)

type Context C.AVFormatContext

// NewAVFormatContext ..
func NewAVFormatContext() *Context {
	return (*Context)(C.avformat_alloc_context())
}

// OpenInput ..
func (c *Context) OpenInput(url string, inFmt *InputFormat, options *avutil.Dictionary) error {
	var c_context = (*C.AVFormatContext)(unsafe.Pointer(c))
	var d_context = (**C.AVFormatContext)(C.malloc(C.ulong(unsafe.Sizeof(c_context))))
	*d_context = c_context
	defer C.free(unsafe.Pointer(d_context))

	var c_options = (*C.AVDictionary)(unsafe.Pointer(options))
	var d_options = (**C.AVDictionary)(C.malloc(C.ulong(unsafe.Sizeof(c_options))))
	*d_options = c_options
	defer C.free(unsafe.Pointer(d_options))

	if code := int(C.avformat_open_input(
		d_context,
		C.CString(url),
		(*C.AVInputFormat)(inFmt),
		d_options,
	)); code < 0 {
		return errors.New(avutil.StrError(code))
	}
	return nil
}

// CloseInput ..
func (c *Context) CloseInput() {
	var c_context = (*C.AVFormatContext)(unsafe.Pointer(c))
	var d_context = (**C.AVFormatContext)(C.malloc(C.ulong(unsafe.Sizeof(c_context))))
	*d_context = c_context
	defer C.free(unsafe.Pointer(d_context))
	C.avformat_close_input(d_context)
}

// FindStreamInfo ..
func (c *Context) FindStreamInfo(options *avutil.Dictionary) error {
	var c_options = (*C.AVDictionary)(unsafe.Pointer(options))
	var d_options = (**C.AVDictionary)(C.malloc(C.ulong(unsafe.Sizeof(c_options))))
	*d_options = c_options
	defer C.free(unsafe.Pointer(d_options))
	if code := int(C.avformat_find_stream_info(
		(*C.AVFormatContext)(unsafe.Pointer(c)),
		d_options,
	)); code < 0 {
		return errors.New(avutil.StrError(code))
	}
	return nil
}

// DumpFormat ..
func (c *Context) DumpFormat() {
	C.av_dump_format((*C.AVFormatContext)(c), C.int(0), nil, 0)
}

// AudioStreamIndex ..
func (c *Context) AudioStreamIndex() (index int, err error) {
	if index = int(C.find_audio_stream_index((*C.AVFormatContext)(c))); index < 0 {
		err = errors.New("audio stream index not found")
	}
	return
}

// AudioStreamCodecID ..
func (c *Context) AudioStreamCodecID(index int) avcodec.CodecID {
	return avcodec.CodecID(C.get_audio_stream_codec_id((*C.AVFormatContext)(c), C.int(index)))
}

// ReadFrame ..
func (c *Context) ReadFrame(pkt *avcodec.Packet) (err error) {
	if code := int(C.av_read_frame((*C.AVFormatContext)(unsafe.Pointer(c)), (*C.AVPacket)(unsafe.Pointer(pkt)))); code < 0 {
		err = avutil.Error(code)
	}
	log.Printf("packet size: %d", (*C.AVPacket)(unsafe.Pointer(pkt)).size)
	log.Println(C.GoBytes(unsafe.Pointer((*C.AVPacket)(unsafe.Pointer(pkt)).data), (*C.AVPacket)(unsafe.Pointer(pkt)).size)[:10])
	return
}
