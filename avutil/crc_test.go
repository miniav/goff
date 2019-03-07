package avutil

import (
	"log"
	"testing"
)

func TestCRCNew(t *testing.T) {
	testCRC(CRC8ATM, "8 atm")
	testCRC(CRC8EBU, "8 ebu")
	testCRC(CRC16ANSI, "16 ansi")
	testCRC(CRC16ANSILE, "16 ansi le")
	testCRC(CRC16CCITT, "16 ccitt")
	testCRC(CRC24IEEE, "24 ieee")
	testCRC(CRC32IEEE, "32 ieee")
	testCRC(CRC32IEEELE, "32 ieee le")
}

func testCRC(t CRCType, info string) {
	if crc, err := CRCNew(t); err != nil {
		log.Fatal(err)
	} else {
		crc.Update(0, []byte("test"))
		log.Printf("crc %s \"test\": %d", info, crc.Update(0, []byte("test")))
	}
}
