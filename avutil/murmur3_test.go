package avutil

import (
	"encoding/hex"
	"testing"
)

func TestMurMur3New(t *testing.T) {
	var m = MurMur3New()
	m.Update([]byte("test"))
	t.Log("murmur3 \"test\":", hex.EncodeToString(m.Final()))
	t.Log("murmur3 \"test\":", hex.EncodeToString(HashMurmur3([]byte("test"))))
}
