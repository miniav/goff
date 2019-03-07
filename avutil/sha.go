package avutil

// #include <libavutil/sha.h>
import "C"
import "unsafe"

// Sha ..
type Sha struct {
	sha *C.struct_AVSHA
	bit uint
}

// ShaNew ..
func ShaNew(b uint) (s *Sha) {
	var bit uint
	if b == 224 || b == 160 || b == 256 {
		bit = b
	} else {
		bit = 512
	}
	var sha = C.av_sha_alloc()
	C.av_sha_init(sha, C.int(bit))
	s = new(Sha)
	s.sha = sha
	s.bit = bit
	return
}

// Update ..
func (s *Sha) Update(b []byte) {
	C.av_sha_update(s.sha, (*C.uint8_t)(unsafe.Pointer(&b[0])), C.uint(len(b)))
}

// Final ..
func (s *Sha) Final() (o []byte) {
	o = make([]byte, s.bit/8)
	C.av_sha_final(s.sha, (*C.uint8_t)(unsafe.Pointer(&o[0])))
	return
}
