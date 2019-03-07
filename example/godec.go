package main

// #include <godec.h>
import "C"
import (
	"io/ioutil"
	"log"
	"os"
	"unsafe"
)

func main() {
	var buffer = C.init_buffer()
	bytes, _ := ioutil.ReadFile("godec.h.gch")

	f, _ := os.Create("godec.out")
	for len(bytes) != 0 {
		var write = int(C.buffer_append(&buffer, (*C.uchar)(unsafe.Pointer(&bytes[0])), C.int(len(bytes))))
		bytes = bytes[write:]
		log.Println(write)
		for {
			var bytes = make([]byte, 1024)
			var read = int(C.buffer_read(&buffer, (*C.uchar)(unsafe.Pointer(&bytes[0])), C.int(len(bytes))))
			if read == 0 {
				break
			}
			f.Write(bytes[:read])
		}
	}
	f.Close()
}
