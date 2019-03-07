package postproc

// #include <libpostproc/postprocess.h>
import "C"

// Version get the version
func Version() uint {
	return uint(C.postproc_version())
}

// License get the license
func License() string {
	return C.GoString(C.postproc_license())
}

// Configuration Return the libpostproc build-time configuration.
func Configuration() string {
	return C.GoString(C.postproc_configuration())
}
