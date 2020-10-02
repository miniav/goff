package avutil

// #include <libavutil/hash.h>
import "C"
import (
	"bytes"
	"fmt"
	"unsafe"
)

type Hash C.struct_AVHashContext

// HashNew ..
func HashNew(method string) (h *Hash, err error) {
	var ctx = &C.struct_AVHashContext{}
	if code := int(C.av_hash_alloc(&ctx, C.CString(method))); code < 0 {
		err = fmt.Errorf("init %s with code: %d", method, code)
		return
	}
	C.av_hash_init(ctx)
	h = ctx
	return
}

// Method ..
func Method(i int) (s string) {
	return C.GoString(C.av_hash_names(C.int(i)))
}

// Method ..
func (h *Hash) Method() string {
	return C.GoString(C.av_hash_get_name((*C.struct_AVHashContext)(h)))
}

// Update ..
func (h *Hash) Update(b []byte) {
	C.av_hash_update((*C.struct_AVHashContext)(h), (*C.uint8_t)(unsafe.Pointer(&b[0])), C.int(len(b)))
}

// Final ..
func (h *Hash) Final() (o []byte) {
	o = make([]byte, int(C.av_hash_get_size((*C.struct_AVHashContext)(h))))
	C.av_hash_final((*C.struct_AVHashContext)(h), (*C.uint8_t)(unsafe.Pointer(&o[0])))
	return
}

// FinalHex ..
func (h *Hash) FinalHex() (o []byte) {
	size := 2*int(C.av_hash_get_size((*C.struct_AVHashContext)(h))) + 1
	o = make([]byte, size)
	C.av_hash_final_hex((*C.struct_AVHashContext)(h), (*C.uint8_t)(unsafe.Pointer(&o[0])), C.int(size))
	o = bytes.Trim(o, "\x00")
	return
}

// FinalBase64 ..
func (h *Hash) FinalBase64() (o []byte) {
	size := ((int(C.av_hash_get_size((*C.struct_AVHashContext)(h))))+2)/3*4 + 1
	o = make([]byte, size)
	C.av_hash_final_b64((*C.struct_AVHashContext)(h), (*C.uint8_t)(unsafe.Pointer(&o[0])), C.int(size))
	o = bytes.Trim(o, "\x00")
	return
}

// Destroy ..
func (h *Hash) Destroy() {
	ctx := (*C.struct_AVHashContext)(h)
	C.av_hash_freep(&ctx)
}
