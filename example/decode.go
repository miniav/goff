package main

/*
#cgo CFLAGS: -I${SRCDIR}/../drivers/include/ -Wno-deprecated-declarations
#cgo darwin,amd64 LDFLAGS: -L${SRCDIR}/../drivers/Darwin-lib
#cgo linux,amd64 LDFLAGS: -L${SRCDIR}/../drivers/Linux-lib
#cgo LDFLAGS: -lavformat -lavcodec -lavutil -lswresample -lswscale

#include <decode.h>
*/
import "C"
import (
	"io/ioutil"
	"log"
	"os"
	"unsafe"

	"github.com/tosone/goff1/avcodec"
	"github.com/tosone/goff1/avutil"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	data, _ := ioutil.ReadFile("../test.mp3")
	var p *C.uchar
	ctx := avcodec.NewContext(avcodec.FindCodec(avcodec.CodecIDMP3))
	parser := avcodec.NewParserContext(int(avcodec.CodecIDMP3))
	log.Println(ctx.Open2(avcodec.FindCodec(avcodec.CodecIDMP3), nil))
	var used int
	var out int
	var f *os.File
	var err error
	if f, err = os.Create("test.raw"); err != nil {
		log.Fatal(err)
	}
	for {
		if code := int(C.decode(
			(*C.AVCodecParserContext)(unsafe.Pointer(parser)),
			(*C.AVCodecContext)(unsafe.Pointer(ctx)),
			(*C.uchar)(unsafe.Pointer(&data[0])),
			C.int(len(data)),
			(*C.int)(unsafe.Pointer(&used)),
			(*C.int)(unsafe.Pointer(&out)),
			(*C.uchar)(unsafe.Pointer(p)),
		)); code < 0 {
			log.Fatal(avutil.StrError(code))
		}
		data = data[used:]
		if _, err = f.Write(C.GoBytes(unsafe.Pointer(p), C.int(out))); err != nil {
			log.Fatal(err)
		}
	}
	f.Close()
}
