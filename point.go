package main

// #include <stdio.h>
// #include <stdlib.h>
//
// void print2(int **a) { printf("%d %p %p\n", **a, *a, a); }
// void print1(int *a) {
//   printf("%d %p %p\n", *a, a, &a);
//   print2(&a);
// }
import "C"
import (
	"unsafe"
)

func main() {
	var a = 269
	var cpointer_a = (*C.int)(unsafe.Pointer(&a))
	C.print1(cpointer_a)

	var b = (**C.int)(C.malloc(C.ulong(unsafe.Sizeof(cpointer_a))))
	*b = cpointer_a
	C.print2(b)
}
