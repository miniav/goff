package avutil

import (
	"encoding/hex"
	"testing"
)

func TestSha512New(t *testing.T) {
	testSha512(224, "512/224", t)
	testSha512(256, "512/256", t)
	testSha512(384, "384", t)
	testSha512(512, "512", t)
}

func testSha512(b uint, info string, t *testing.T) {
	var s = Sha512New(b)
	s.Update([]byte("test"))
	t.Logf("sha%s \"test\": %s\n", info, hex.EncodeToString(s.Final()))
}
