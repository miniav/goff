package swresample

import (
	"log"
	"testing"
)

func TestVersion(t *testing.T) {
	log.Println(Version())
	log.Println(License())
}
