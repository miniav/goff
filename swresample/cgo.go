package swresample

/*
#cgo CFLAGS: -I${SRCDIR}/../install/include/
#cgo darwin LDFLAGS: -L${SRCDIR}/../install/darwin-lib
#cgo linux LDFLAGS: -L${SRCDIR}/../install/linux-lib
#cgo LDFLAGS: -lswresample -lavutil -lm
*/
import "C"

//  play -t raw -r 16k -b 16 -c 1 -e unsigned-integer s.pcm
