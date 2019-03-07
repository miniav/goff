package avutil

import (
	"encoding/base64"
	"encoding/hex"
	"log"
	"testing"
)

func TestHashNew(t *testing.T) {
	for i := 0; i < 20; i++ {
		if method := Method(i); method == "" {
			break
		} else {
			if hash, err := HashNew(method); err != nil {
				log.Fatal(err)
			} else {
				hash.Update([]byte("test"))
				result := hash.Final()
				log.Printf("hash %s hex \"test\": %s\n", method, hex.EncodeToString(result))
				log.Printf("hash %s base64 \"test\": %s\n", method, base64.StdEncoding.EncodeToString(result))
				hash.Destroy()
			}
			if hash, err := HashNew(method); err != nil {
				log.Fatal(err)
			} else {
				hash.Update([]byte("test"))
				log.Printf("hash %s hex \"test\": %s\n", method, string(hash.FinalHex()))
				hash.Destroy()
			}
			if hash, err := HashNew(method); err != nil {
				log.Fatal(err)
			} else {
				hash.Update([]byte("test"))
				log.Printf("hash %s base64 \"test\": %s\n", method, string(hash.FinalBase64()))
				hash.Destroy()
			}
		}
	}
}
