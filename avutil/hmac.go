package avutil

// #include <libavutil/hmac.h>
import "C"
import (
	"bytes"
	"unsafe"
)

// HMacType ..
type HMacType C.enum_AVHMACType

// HMac ..
type HMac C.AVHMAC

const hMacMaxSize = 64

var (
	// HMacMD5 ..
	HMacMD5 = HMacType(C.AV_HMAC_MD5)
	// HMacSHA1 ..
	HMacSHA1 = HMacType(C.AV_HMAC_SHA1)
	// HMacSHA224 ..
	HMacSHA224 = HMacType(C.AV_HMAC_SHA224)
	// HMacSHA256 ..
	HMacSHA256 = HMacType(C.AV_HMAC_SHA256)
	// HMacSHA384 ..
	HMacSHA384 = HMacType(C.AV_HMAC_SHA384)
	// HMacSHA512 ..
	HMacSHA512 = HMacType(C.AV_HMAC_SHA512)
)

func HMacCalc(t HMacType, k string, b []byte) (o []byte) {
	ctx := C.av_hmac_alloc(C.enum_AVHMACType(t))
	key := []byte(k)
	o = make([]byte, hMacMaxSize)
	C.av_hmac_calc(
		ctx,
		(*C.uint8_t)(unsafe.Pointer(&b[0])),
		C.uint(len(b)),
		(*C.uint8_t)(unsafe.Pointer(&key[0])),
		C.uint(len(key)),
		(*C.uint8_t)(unsafe.Pointer(&o[0])),
		C.uint(hMacMaxSize),
	)
	o = bytes.Trim(o, "\x00")
	return
}

// HMacNew ..
func HMacNew(t HMacType, k string) (h *HMac) {
	var ctx = C.av_hmac_alloc(C.enum_AVHMACType(t))
	key := []byte(k)
	C.av_hmac_init(ctx, (*C.uint8_t)(unsafe.Pointer(&key[0])), C.uint(len(key)))
	h = (*HMac)(ctx)
	return
}

func (h *HMac) Update(b []byte) {
	C.av_hmac_update((*C.struct_AVHMAC)(h), (*C.uint8_t)(unsafe.Pointer(&b[0])), C.uint(len(b)))
}

func (h *HMac) Final() (o []byte) {
	o = make([]byte, hMacMaxSize)
	C.av_hmac_final((*C.struct_AVHMAC)(h), (*C.uint8_t)(unsafe.Pointer(&o[0])), C.uint(hMacMaxSize))
	o = bytes.Trim(o, "\x00")
	return
}

// HMacDestroy ..
func (h *HMac) HMacDestroy() {
	C.av_hmac_free((*C.struct_AVHMAC)(h))
}
