package avformat

// #cgo CFLAGS: -I${SRCDIR}/../drivers/include/ -Wno-deprecated-declarations
// #cgo darwin,amd64 LDFLAGS: -L${SRCDIR}/../drivers/Darwin-lib
// #cgo linux,amd64 LDFLAGS: -L${SRCDIR}/../drivers/Linux-lib
// #cgo LDFLAGS: -lavformat -lavcodec -lavutil -lswresample -lswscale
import "C"
