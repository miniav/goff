package avutil

import (
	"log"
	"testing"
)

func TestStrStart(t *testing.T) {
	log.Println(StrStart("pretosone", "pre"))
}

func TestStrIStart(t *testing.T) {
	log.Println(StrIStart("PrEtosone", "pre"))
}
