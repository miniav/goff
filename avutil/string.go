package avutil

// #include <libavutil/avstring.h>
import "C"

// StrStart Return non-zero if pfx is a prefix of str.
// https://www.ffmpeg.org/doxygen/trunk/group__lavu__string.html#ga6135a12628e2b6a63c8d3d9b3a742b06
func StrStart(str, prefix string) (b bool, out string) {
	var s = C.CString(str)
	b = int(C.av_strstart(
		s,
		C.CString(prefix),
		&s,
	)) != 0
	out = C.GoString(s)
	return
}

// StrIStart Return non-zero if pfx is a prefix of str independent of case.
// https://www.ffmpeg.org/doxygen/trunk/group__lavu__string.html#gacd29ff1f7e62230a113c88fa10d3f5b9
func StrIStart(str, prefix string) (b bool, out string) {
	var s = C.CString(str)
	b = int(C.av_stristart(
		s,
		C.CString(prefix),
		&s,
	)) != 0
	out = C.GoString(s)
	return
}
