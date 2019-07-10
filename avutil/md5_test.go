package avutil

import (
	"encoding/hex"
	"testing"
)

func TestHashMD5(t *testing.T) {
	t.Log("md5 \"test\":", hex.EncodeToString(HashMD5([]byte("test"))))

	md5 := MD5New()
	md5.Update([]byte("test"))
	t.Log("md5 \"test\":", hex.EncodeToString(md5.Final()))
}
