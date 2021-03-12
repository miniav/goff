package avutil

// #include <libavutil/crc.h>
import "C"
import (
	"fmt"
	"unsafe"
)

// CRC
type CRC struct {
	crc *C.AVCRC
	t   CRCType
}

// CRCType
type CRCType struct {
	t    C.AVCRCId
	le   int
	bits int
	poly uint32
}

var (
	// CRC8ATM
	CRC8ATM = CRCType{C.AV_CRC_8_ATM, 0, 8, 0x07}
	// CRC8EBU
	CRC8EBU = CRCType{C.AV_CRC_8_EBU, 0, 8, 0x1D}
	// CRC16ANSI
	CRC16ANSI = CRCType{C.AV_CRC_16_ANSI, 0, 16, 0x8005}
	// CRC16ANSILE
	CRC16ANSILE = CRCType{C.AV_CRC_16_ANSI_LE, 1, 16, 0xA001}
	// CRC16CCITT
	CRC16CCITT = CRCType{C.AV_CRC_16_CCITT, 0, 16, 0x1021}
	// CRC24IEEE
	CRC24IEEE = CRCType{C.AV_CRC_24_IEEE, 0, 24, 0x864CFB}
	// CRC32IEEE
	CRC32IEEE = CRCType{C.AV_CRC_32_IEEE, 0, 32, 0x04C11DB7}
	// CRC32IEEELE
	CRC32IEEELE = CRCType{C.AV_CRC_32_IEEE_LE, 1, 32, 0xEDB88320}
)

// CRCNew
func CRCNew(t CRCType) (c *CRC, err error) {
	var crc = C.av_crc_get_table(t.t)
	if code := int(C.av_crc_init(crc, C.int(t.le), C.int(t.bits), C.uint32_t(t.poly), C.sizeof_AVCRC*1024)); code < 0 {
		err = fmt.Errorf("init crc error code: %d", code)
		return
	}
	c = new(CRC)
	c.crc = crc
	c.t = t
	return
}

// Update
func (c *CRC) Update(init uint32, b []byte) uint32 {
	return uint32(C.av_crc(c.crc, C.uint32_t(init), (*C.uint8_t)(unsafe.Pointer(&b[0])), C.size_t(len(b))))
}
