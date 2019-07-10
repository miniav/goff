package avutil

import (
	"testing"
)

func TestMalloc(t *testing.T) {
	t.Logf("malloc space: %p\n", Malloc(1024))
}
