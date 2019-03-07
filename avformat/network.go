package avformat

// #include <libavformat/avformat.h>
import "C"

// NetworkInit ..
func NetworkInit() {
	C.avformat_network_init()
}

// NetworkDeInit ..
func NetworkDeInit() {
	C.avformat_network_deinit()
}
