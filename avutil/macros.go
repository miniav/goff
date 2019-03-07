package avutil

// #include <libavutil/avutil.h>
import "C"

var (
	NOPTS_VALUE = int64(C.AV_NOPTS_VALUE)
	TIME_BASE   = C.AV_TIME_BASE
	TIME_BASE_Q = C.AV_TIME_BASE_Q
)
