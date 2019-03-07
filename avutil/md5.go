package avutil

// #include <libavutil/md5.h>
import "C"
import (
	"unsafe"
)

const md5Size = 16

// MD5 ..
type MD5 C.struct_AVMD5

// HashMD5 ..
func HashMD5(b []byte) (o []byte) {
	o = make([]byte, md5Size)
	C.av_md5_sum((*C.uint8_t)(unsafe.Pointer(&o[0])), (*C.uint8_t)(unsafe.Pointer(&b[0])), C.int(len(b)))
	return
}

// MD5New ..
func MD5New() (m *MD5) {
	var md5 = C.av_md5_alloc()
	C.av_md5_init(md5)
	m = (*MD5)(md5)
	return
}

// Update ..
func (m *MD5) Update(b []byte) {
	C.av_md5_update((*C.struct_AVMD5)(m), (*C.uint8_t)(unsafe.Pointer(&b[0])), C.int(len(b)))
}

// Final ..
func (m *MD5) Final() (o []byte) {
	o = make([]byte, md5Size)
	C.av_md5_final((*C.struct_AVMD5)(m), (*C.uint8_t)(unsafe.Pointer(&o[0])))
	return
}
