package avutil

// #include <frame.h>
import "C"
import "unsafe"

type Frame C.AVFrame

func NewFrame() (frame *Frame) {
	return (*Frame)(C.av_frame_alloc())
}

// Unref ..
func (f *Frame) Unref() {
	C.av_frame_unref((*C.AVFrame)(unsafe.Pointer(f)))
}

// ToBytes ..
func (f *Frame) ToBytes(channels, bytesPerSample int) (b []byte) {
	samples := int(C.frame_samples((*C.AVFrame)(unsafe.Pointer(f))))
	c := make([]byte, samples*channels*bytesPerSample)
	C.frame_to_bytes((*C.AVFrame)(unsafe.Pointer(f)), C.int(channels), C.int(bytesPerSample), (*C.uchar)(unsafe.Pointer(&c[0])))
	b = c[:]
	return
}
