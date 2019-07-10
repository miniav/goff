package avcodec

import (
	"testing"
)

func TestVersion(t *testing.T) {
	t.Log("version:", Version())
}

func TestLicense(t *testing.T) {
	t.Log("license:", License())
}

func TestConfiguration(t *testing.T) {
	t.Log("configuration:", Configuration())
}
