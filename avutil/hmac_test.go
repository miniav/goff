package avutil

import (
	"encoding/hex"
	"testing"
)

func TestHMacNew(t *testing.T) {
	testHMac(HMacMD5, "md5", t)
	testHMac(HMacSHA1, "sha1", t)
	testHMac(HMacSHA224, "sha224", t)
	testHMac(HMacSHA256, "sha256", t)
	testHMac(HMacSHA384, "sha384", t)
	testHMac(HMacSHA512, "sha512", t)
}

func testHMac(t HMacType, info string, test *testing.T) {
	hmac := HMacNew(t, "test")
	hmac.Update([]byte("test"))
	test.Logf("hmac %s \"test\": %s\n", info, hex.EncodeToString(hmac.Final()))
	test.Logf("hmac %s \"test\": %s\n", info, hex.EncodeToString(HMacCalc(t, "test", []byte("test"))))
}
