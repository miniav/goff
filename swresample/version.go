package swresample

// #include <libswresample/swresample.h>
import "C"

// Version get the version
func Version() uint {
	return uint(C.swresample_version())
}

// License get the license
func License() string {
	return C.GoString(C.swresample_license())
}
