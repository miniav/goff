package avutil

// #include <libavutil/murmur3.h>
import "C"
import (
	"math/rand"
	"time"
	"unsafe"
)

const murmur3Size = 16

type Murmur3 C.struct_AVMurMur3

// HashMurmur3 ..
func HashMurmur3(b []byte) []byte {
	var m = MurMur3New()
	m.Update(b)
	return m.Final()
}

// MurMur3New ..
func MurMur3New() (m *Murmur3) {
	rand.Seed(time.Now().UTC().UnixNano())
	seed := rand.Uint64()
	var mur = C.av_murmur3_alloc()
	C.av_murmur3_init(mur)
	C.av_murmur3_init_seeded(mur, C.uint64_t(seed))
	m = (*Murmur3)(mur)
	return
}

// Update ..
func (m *Murmur3) Update(b []byte) {
	C.av_murmur3_update((*C.struct_AVMurMur3)(m), (*C.uint8_t)(unsafe.Pointer(&b[0])), C.int(len(b)))
}

// Final ..
func (m *Murmur3) Final() (o []byte) {
	o = make([]byte, murmur3Size)
	C.av_murmur3_final((*C.struct_AVMurMur3)(m), (*C.uint8_t)(unsafe.Pointer(&o[0])))
	return
}
