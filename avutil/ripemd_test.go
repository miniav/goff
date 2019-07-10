package avutil

import (
	"encoding/hex"
	"testing"
)

func TestRipemdNew(t *testing.T) {
	testRipemd(128, t)
	testRipemd(160, t)
	testRipemd(256, t)
	testRipemd(320, t)
}

func testRipemd(b uint, t *testing.T) {
	if ripemd, err := RipemdNew(b); err != nil {
		t.Fatal(err)
	} else {
		ripemd.Update([]byte("test"))
		t.Logf("ripemd %d \"test\": %s\n", b, hex.EncodeToString(ripemd.Final()))
	}
}
