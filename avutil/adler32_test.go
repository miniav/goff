package avutil

import (
	"log"
	"testing"
)

func TestHashAdler32(t *testing.T) {
	log.Println("adler32 \"test\":", HashAdler32(0, []byte("test")))
}
