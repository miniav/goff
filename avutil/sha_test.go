package avutil

import (
	"encoding/hex"
	"testing"
)

func TestShaNew(t *testing.T) {
	testSha(160, t)
	testSha(224, t)
	testSha(256, t)
}

func testSha(b uint, t *testing.T) {
	var s = ShaNew(b)
	s.Update([]byte("test"))
	t.Logf("sha%d \"test\": %s\n", b, hex.EncodeToString(s.Final()))
}
