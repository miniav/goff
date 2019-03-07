package avutil

import (
	"encoding/hex"
	"log"
	"testing"
)

func TestSha512New(t *testing.T) {
	testSha512(224, "512/224")
	testSha512(256, "512/256")
	testSha512(384, "384")
	testSha512(512, "512")
}

func testSha512(b uint, info string) {
	var s = Sha512New(b)
	s.Update([]byte("test"))
	log.Printf("sha%s \"test\": %s\n", info, hex.EncodeToString(s.Final()))
}
