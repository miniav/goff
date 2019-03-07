package avutil

// #include <libavutil/avutil.h>
import "C"

// Version get the version
func Version() uint {
	return uint(C.avutil_version())
}

// License get the license
func License() string {
	return C.GoString(C.avutil_license())
}
