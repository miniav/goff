package main

import (
	"fmt"
	"unsafe"
)

//#include <stdlib.h>
import "C"

var xs = []string{
	"1",
	"",
	"squirrel",
	"2",
	"two",
}

func main() {
	fmt.Println("Test conversion")
	fmt.Printf("\tN\t=\t%20s\t%20s\n", "Go", "C")
	//if I delete this loop everything works as intended
	for i, x := range xs {
		cvt := C.GoString(C.CString(x))
		//always equal
		fmt.Printf("\t%d\t%t\t%20q\t%20q\n", i, x == cvt, x, cvt)
	}

	//create an array meant to be passed to C
	//and a view into it to update the data
	size := C.size_t(unsafe.Sizeof((*C.char)(nil)))
	length := C.size_t(len(xs))
	arr := (**C.char)(C.malloc(length * size))
	view := (*[1 << 30]*C.char)(unsafe.Pointer(arr))[0:len(xs):len(xs)]

	fmt.Println("To C", &view[0], arr)
	fmt.Printf("\tN\t=\t%20s\t%20s\t&C\n", "Go", "C")
	for i, x := range xs {
		view[i] = C.CString(x)
		// (deleting the rest of this loop has no effect on the issue, but it's interesting that it still works at this point)
		g := C.GoString(view[i])
		//always equal
		fmt.Printf("\t%d\t%t\t%20q\t%20q\t%v\n", i, x == g, x, g, view[i])
	}

	//Everything is fine so far, but reading the values back out
	//(whether in Go or C) does not work
	fmt.Println("From C", &view[0], arr)
	fmt.Printf("\tN\t=\t%20s\t%20s\t&C\n", "Go", "C")
	for i, x := range view {
		g := C.GoString(x)
		//only the same after the first few entries
		fmt.Printf("\t%d\t%t\t%20q\t%20q\t%v\n", i, xs[i] == g, xs[i], g, x)
	}
}
