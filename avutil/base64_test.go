package avutil

import (
	"log"
	"testing"
)

func TestBase64Encode(t *testing.T) {
	var err error

	var str = "test"

	var base []byte
	if base, err = Base64Encode([]byte(str)); err != nil {
		log.Fatal(err)
	}
	var out []byte
	if out, err = Base64Decode(base); err != nil {
		log.Fatal(err)
	}
	log.Println("base64 origin:", str, len(str), []byte(str))
	log.Println("base64 result:", string(out), len(string(out)), out)
}
