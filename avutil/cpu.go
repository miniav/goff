package avutil

// #include <libavutil/cpu.h>
// #include <stddef.h>
import "C"

// CPUCount the number of logical CPU cores present.
func CPUCount() int {
	return int(C.av_cpu_count())
}

func CPUMaxAlign() uint64 {
	return uint64(C.av_cpu_max_align())
}
