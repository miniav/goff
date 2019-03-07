package test2

// #include <test.h>
import "C"
import "log"

func Test() {
	log.Println(C.a)
}
