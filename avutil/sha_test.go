package avutil

import (
	"encoding/hex"
	"log"
	"testing"
)

func TestShaNew(t *testing.T) {
	testSha(160)
	testSha(224)
	testSha(256)
}

func testSha(b uint) {
	var s = ShaNew(b)
	s.Update([]byte("test"))
	log.Printf("sha%d \"test\": %s\n", b, hex.EncodeToString(s.Final()))
}
