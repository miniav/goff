package avutil

import (
	"encoding/hex"
	"log"
	"testing"
)

func TestRipemdNew(t *testing.T) {
	testRipemd(128)
	testRipemd(160)
	testRipemd(256)
	testRipemd(320)
}

func testRipemd(b uint) {
	if ripemd, err := RipemdNew(b); err != nil {
		log.Fatal(err)
	} else {
		ripemd.Update([]byte("test"))
		log.Printf("ripemd %d \"test\": %s\n", b, hex.EncodeToString(ripemd.Final()))
	}
}
