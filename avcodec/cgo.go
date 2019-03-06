package avcodec

/*
#cgo CFLAGS: -I${SRCDIR}/../install/include
#cgo darwin LDFLAGS: -L${SRCDIR}/../install/darwin-lib
#cgo linux LDFLAGS: -L${SRCDIR}/../install/linux-lib
#cgo LDFLAGS: -lswscale -lavformat -lavdevice -lavfilter -lavcodec -lavutil -lswresample -lpostproc -lm
*/
import "C"
