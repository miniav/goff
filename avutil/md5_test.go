package avutil

import (
	"encoding/hex"
	"log"
	"testing"
)

func TestHashMD5(t *testing.T) {
	log.Println("md5 \"test\":", hex.EncodeToString(HashMD5([]byte("test"))))

	md5 := MD5New()
	md5.Update([]byte("test"))
	log.Println("md5 \"test\":", hex.EncodeToString(md5.Final()))
}
