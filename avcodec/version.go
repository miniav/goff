package avcodec

// #include <version.h>
import "C"

// Version get the version
func Version() uint {
	return uint(C.avcodec_version())
}

// License get the license
func License() string {
	return C.GoString(C.avcodec_license())
}

// Configuration Return the build-time configuration.
func Configuration() string {
	return C.GoString(C.avcodec_configuration())
}
