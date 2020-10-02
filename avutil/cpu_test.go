package avutil

import (
	"testing"
)

func TestCPUCount(t *testing.T) {
	t.Logf("CPU Count: %d\n", CPUCount())
}

func TestCPUMaxAlign(t *testing.T) {
	t.Logf("CPU max align: %d", CPUMaxAlign())
}