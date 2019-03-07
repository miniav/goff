package avutil

// #include <libavutil/ripemd.h>
import "C"
import (
	"errors"
	"unsafe"
)

type Ripemd struct {
	ripemd *C.struct_AVRIPEMD
	bit    uint
}

// RipemdNew ..
func RipemdNew(b uint) (r *Ripemd, err error) {
	var bit uint
	if b == 128 || b == 160 || b == 256 || b == 320 {
		bit = b
	} else {
		bit = 128
	}
	var ripemd = C.av_ripemd_alloc()
	if int(C.av_ripemd_init(ripemd, C.int(bit))) != 0 {
		err = errors.New("init ripemd error")
		return
	}
	r = new(Ripemd)
	r.ripemd = ripemd
	r.bit = bit
	return
}

// Update ..
func (r *Ripemd) Update(b []byte) {
	C.av_ripemd_update(r.ripemd, (*C.uint8_t)(unsafe.Pointer(&b[0])), C.uint(len(b)))
}

// Final ..
func (r *Ripemd) Final() (o []byte) {
	o = make([]byte, r.bit/8)
	C.av_ripemd_final(r.ripemd, (*C.uint8_t)(unsafe.Pointer(&o[0])))
	return
}
