package avutil

import (
	"testing"
)

func TestFFMax(t *testing.T) {
	t.Log(FFMax(1, 2))
}

func TestFFMax3(t *testing.T) {
	t.Log(FFMax3(1, 2, 3))
}

func TestFFMin(t *testing.T) {
	t.Log(FFMin(1, 2))
}

func TestFFMin3(t *testing.T) {
	t.Log(FFMin3(1, 2, 3))
}
