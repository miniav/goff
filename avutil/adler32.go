package avutil

// #include <libavutil/adler32.h>
import "C"
import "unsafe"

// HashAdler32 ..
func HashAdler32(init uint32, b []byte) (h uint32) {
	h = uint32(C.av_adler32_update(C.ulong(init), (*C.uint8_t)(unsafe.Pointer(&b[0])), C.uint(len(b))))
	return
}
