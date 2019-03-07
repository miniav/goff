package swresample

// #include <libswresample/swresample.h>
import "C"

type Frame C.struct_AVFrame
type Class C.struct_AVClass
type AvSampleFormat C.enum_AVSampleFormat

//Get the Class for Context.
func SwrGetClass() *Class {
	return (*Class)(C.swr_get_class())
}

func SwresampleConfiguration() string {
	return C.GoString(C.swresample_configuration())
}
