package avutil

// #include <libavutil/time.h>
import "C"

// GetTime ..
func GetTime() int64 {
	return int64(C.av_gettime())
}

// GetRelativeTime ..
func GetRelativeTime() int64 {
	return int64(C.av_gettime_relative())
}

// GetRelativeTimeIsMonotonic Indicates with a boolean result if the av_gettime_relative() time source is monotonic.
func GetRelativeTimeIsMonotonic() bool {
	return int(C.av_gettime_relative_is_monotonic()) == 1
}

// Sleep sleep for a period of time.  Although the duration is expressed in microseconds, the actual delay may be
// rounded to the precision of the system timer.
func Sleep(t uint32) int {
	return int(C.av_usleep(C.uint(t)))
}
