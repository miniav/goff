package avutil

import (
	"log"
	"testing"
)

func TestMalloc(t *testing.T) {
	log.Printf("malloc space: %p\n", Malloc(1024))
}
