package avutil

// #include <libavutil/sha512.h>
import "C"
import "unsafe"

// Sha512 ..
type Sha512 struct {
	sha512 *C.struct_AVSHA512
	bit    uint
}

// Sha512New ..
func Sha512New(b uint) (s *Sha512) {
	var bit uint
	if b == 224 || b == 256 || b == 384 || b == 512 {
		bit = b
	} else {
		bit = 512
	}
	var sha = C.av_sha512_alloc()
	C.av_sha512_init(sha, C.int(bit))
	s = new(Sha512)
	s.sha512 = sha
	s.bit = bit
	return
}

// Update ..
func (s *Sha512) Update(b []byte) {
	C.av_sha512_update(s.sha512, (*C.uint8_t)(unsafe.Pointer(&b[0])), C.uint(len(b)))
}

// Final ..
func (s *Sha512) Final() (o []byte) {
	o = make([]byte, s.bit/8)
	C.av_sha512_final(s.sha512, (*C.uint8_t)(unsafe.Pointer(&o[0])))
	return
}
