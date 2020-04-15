package avutil

// #include <common.hpp>
import "C"

func FFMax(a, b int) int {
	return int(C.ffmax(C.int(a), C.int(b)))
}

func FFMax3(a, b, c int) int {
	return int(C.ffmax3(C.int(a), C.int(b), C.int(c)))
}

func FFMin(a, b int) int {
	return int(C.ffmin(C.int(a), C.int(b)))
}

func FFMin3(a, b, c int) int {
	return int(C.ffmin3(C.int(a), C.int(b), C.int(c)))
}
