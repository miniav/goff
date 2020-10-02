package avutil

import "testing"

func TestGetTime(t *testing.T) {
	t.Logf("GetTime: %d", GetTime())
}

func TestGetRelativeTime(t *testing.T) {
	t.Logf("GetRelativeTime: %d", GetRelativeTime())
}

func TestGetRelativeTimeIsMonotonic(t *testing.T) {
	t.Logf("GetRelativeTimeIsMonotonic: %v", GetRelativeTimeIsMonotonic())
}

func TestSleep(t *testing.T) {
	t.Logf("Sleep for a while: %d", Sleep(1000*3))
}
