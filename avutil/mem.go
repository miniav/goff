package avutil

// #include <libavutil/mem.h>
import "C"
import "unsafe"

// Malloc 申请空间
func Malloc(size uint64) unsafe.Pointer {
	return C.av_malloc(C.ulong(size))
}
