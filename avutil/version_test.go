package avutil

import (
	"testing"
)

func TestVersion(t *testing.T) {
	t.Log("version:", Version())
	t.Log("license:", License())
}
