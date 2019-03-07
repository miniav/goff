package test1

// #include <test.h>
import "C"
import "log"

func Test() {
	log.Println(C.a)
}
