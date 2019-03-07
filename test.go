package main

/*
#include <stdio.h>
#include <string.h>

void fill_array(char *s) {
    strcpy(s, "cobbliu");
}
void fill_2d_array(char **arr, int columeSize) {
    strcpy((char*)(arr+0*sizeof(char)*columeSize), "hello");
    strcpy((char*)(arr+1*sizeof(char)*columeSize/sizeof(char*)), "cgo");
}
*/
import "C"
import "fmt"
import "unsafe"

func main() {
	var dir = make([]byte, 10)
	C.fill_array((*C.char)(unsafe.Pointer(&dir[0])))
	fmt.Println(string(dir[:]))

	// var dirs [4][16]byte
	dirs := make([][]byte, 4)
	for i := 0; i < 4; i++ {
		dirs[i] = make([]byte, 16)
	}

	C.fill_2d_array((**C.char)(unsafe.Pointer(&dirs[0][0])), C.int(16))
	fmt.Println(dirs)
}
