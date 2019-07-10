package avutil

import (
	"testing"
)

func TestStrStart(t *testing.T) {
	t.Log(StrStart("pretosone", "pre"))
}

func TestStrIStart(t *testing.T) {
	t.Log(StrIStart("PrEtosone", "pre"))
}
