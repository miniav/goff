package avutil

import (
	"testing"
)

func TestHashAdler32(t *testing.T) {
	t.Log("adler32 \"test\":", HashAdler32(0, []byte("test")))
}
