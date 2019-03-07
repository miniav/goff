package avutil

import (
	"log"
	"testing"
)

func TestVersion(t *testing.T) {
	log.Println("version:", Version())
	log.Println("license:", License())
}
