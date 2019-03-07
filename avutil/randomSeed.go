package avutil

// #include <libavutil/random_seed.h>
import "C"

func RandomSeed() uint32 {
	return uint32(C.av_get_random_seed())
}
