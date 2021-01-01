package avutil

// #include <libavutil/aes.h>
import "C"

type AES C.struct_AVAES

//
//func AESNew(key []byte) (aes *AES) {
//	var ctx = C.av_aes_alloc()
//	C.av_aes_init(ctx, (*C.uint8_t)(unsafe.Pointer(&key[0])), 128, 0)
//	return
//}
//
//func (a *AES) AESEncrypt() {
//	C.av_aes_crypt(a)
//}
