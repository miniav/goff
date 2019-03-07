package avformat

// #include <libavformat/avformat.h>
import "C"

// Version get the version
func Version() uint {
	return uint(C.avformat_version())
}

// License get the license
func License() string {
	return C.GoString(C.avformat_license())
}

// Configuration Return the build-time configuration.
func Configuration() string {
	return C.GoString(C.avformat_configuration())
}
