package avutil

import (
	"encoding/hex"
	"log"
	"testing"
)

func TestMurMur3New(t *testing.T) {
	var m = MurMur3New()
	m.Update([]byte("test"))
	log.Println("murmur3 \"test\":", hex.EncodeToString(m.Final()))
	log.Println("murmur3 \"test\":", hex.EncodeToString(HashMurmur3([]byte("test"))))
}
