package avutil

// #include <base64.h>
import "C"
import (
	"bytes"
	"fmt"
	"unsafe"
)

// Base64Decode 编码
func Base64Decode(b []byte) (o []byte, err error) {
	var outSize = C.int((len(b)) * 3 / 4)
	var gout = make([]byte, outSize)
	var out = (*C.uint8_t)(unsafe.Pointer(&gout[0]))
	if flag := int(C.av_base64_decode(out, (*C.char)(unsafe.Pointer(&b[0])), outSize)); flag < 0 {
		err = fmt.Errorf("cannot decode: %s", string(b))
		return
	}

	o = bytes.Trim(gout, "\x00")
	return
}

// Base64Encode 解码
func Base64Encode(b []byte) (o []byte, err error) {
	var outSize = C.int((((len(b))+2)/3*4 + 1))
	var gout = make([]byte, outSize)
	var out = (*C.char)(unsafe.Pointer(&gout[0]))
	if flag := C.av_base64_encode(out, outSize, (*C.uint8_t)(unsafe.Pointer(&b[0])), C.int(len(b))); flag == nil {
		err = fmt.Errorf("cannot encode this to base64: %s", string(b))
		return
	}
	o = gout[:]
	return
}
