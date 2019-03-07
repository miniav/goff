package avformat

import (
	"log"
	"testing"
)

func TestVersion(t *testing.T) {
	log.Println("version:", Version())
}

func TestLicense(t *testing.T) {
	log.Println("license:", License())
}

func TestConfiguration(t *testing.T) {
	log.Println("configuration:", Configuration())
}
