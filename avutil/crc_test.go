package avutil

import (
	"log"
	"testing"
)

func TestCRCNew(t *testing.T) {
	testCRC(CRC8ATM, "8 atm", t)
	testCRC(CRC8EBU, "8 ebu", t)
	testCRC(CRC16ANSI, "16 ansi", t)
	testCRC(CRC16ANSILE, "16 ansi le", t)
	testCRC(CRC16CCITT, "16 ccitt", t)
	testCRC(CRC24IEEE, "24 ieee", t)
	testCRC(CRC32IEEE, "32 ieee", t)
	testCRC(CRC32IEEELE, "32 ieee le", t)
}

func testCRC(t CRCType, info string, test *testing.T) {
	if crc, err := CRCNew(t); err != nil {
		log.Fatal(err)
	} else {
		crc.Update(0, []byte("test"))
		test.Logf("crc %s \"test\": %d", info, crc.Update(0, []byte("test")))
	}
}
