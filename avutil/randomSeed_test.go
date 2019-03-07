package avutil

import (
	"log"
	"testing"
)

func TestRandomSeed(t *testing.T) {
	log.Println("random seed:", RandomSeed())
}
