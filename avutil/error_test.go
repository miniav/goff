package avutil

import (
	"log"
	"testing"
)

func TestStrError(t *testing.T) {
	log.Println("error string:", StrError(-22))
}
