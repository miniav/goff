package avutil

import (
	"encoding/hex"
	"log"
	"testing"
)

func TestHMacNew(t *testing.T) {
	testHMac(HMacMD5, "md5")
	testHMac(HMacSHA1, "sha1")
	testHMac(HMacSHA224, "sha224")
	testHMac(HMacSHA256, "sha256")
	testHMac(HMacSHA384, "sha384")
	testHMac(HMacSHA512, "sha512")
}

func testHMac(t HMacType, info string) {
	hmac := HMacNew(t, "test")
	hmac.Update([]byte("test"))
	log.Printf("hmac %s \"test\": %s\n", info, hex.EncodeToString(hmac.Final()))
	log.Printf("hmac %s \"test\": %s\n", info, hex.EncodeToString(HMacCalc(t, "test", []byte("test"))))
}
