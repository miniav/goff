package swresample

// #include <libswresample/swresample.h>
// #include <libavutil/opt.h>
// #include <stdlib.h>
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/miniav/goff/avutil"
)

type Context C.struct_SwrContext

// NewContext Allocate SwrContext.
// If you use this function you will need to set the parameters (manually or with swr_alloc_set_opts()) before calling swr_init().
// Returns NULL on error, allocated context otherwise
func NewContext() (ctx *Context, err error) {
	if ctx = (*Context)(C.swr_alloc()); ctx == nil {
		err = fmt.Errorf("cannot allocate context")
	}
	return
}

// NewContextOpt Allocate SwrContext if needed and set/reset common parameters.
// This function does not require s to be allocated with swr_alloc(). On the
// other hand, swr_alloc() can use swr_alloc_set_opts() to set the parameters
// on the allocated context.
func NewContextOpt(ctx *Context, outCh int64, outSampleFmt avutil.SampleFmt, outSampleRate int,
	inCh int64, inSampleFmt avutil.SampleFmt, inSampleRate int) (newCtx *Context, err error) {
	if ctx = (*Context)(C.swr_alloc_set_opts(
		(*C.struct_SwrContext)(ctx),
		C.int64_t(outCh),
		C.enum_AVSampleFormat(outSampleFmt),
		C.int(outSampleRate),
		C.int64_t(inCh),
		C.enum_AVSampleFormat(inSampleFmt),
		C.int(inSampleRate),
		0, nil,
	)); ctx == nil {
		err = fmt.Errorf("cannot allocate context")
	}
	return
}

func (s *Context) SetOptString(name, value string, flag int) (err error) {
	ret := int(C.av_opt_set(unsafe.Pointer((*C.struct_SwrContext)(s)), C.CString(name), C.CString(value), C.int(flag)))
	if ret == 0 {
		err = nil
	}
	return
}

func (s *Context) SetOptInt(name string, value int64, flag int) (err error) {
	ret := int(C.av_opt_set_int(unsafe.Pointer((*C.struct_SwrContext)(s)), C.CString(name), C.int64_t(value), C.int(flag)))
	if ret == 0 {
		err = nil
	}
	return
}

func (s *Context) SetOptSampleFmt(name string, fmt avutil.SampleFmt, flag int) (err error) {
	ret := int(C.av_opt_set_sample_fmt(unsafe.Pointer((*C.struct_SwrContext)(s)), C.CString(name), C.enum_AVSampleFormat(fmt), C.int(flag)))
	if ret == 0 {
		err = nil
	}
	return
}

//Initialize context after user parameters have been set.
func (s *Context) SwrInit() int {
	return int(C.swr_init((*C.struct_SwrContext)(s)))
}

//Check whether an swr context has been initialized or not.
func (s *Context) SwrIsInitialized() int {
	return int(C.swr_is_initialized((*C.struct_SwrContext)(s)))
}

//Allocate Context if needed and set/reset common parameters.
func (s *Context) SwrAllocSetOpts(ocl int64, osf AvSampleFormat, osr int, icl int64, isf AvSampleFormat, isr, lo, lc int) *Context {
	return (*Context)(C.swr_alloc_set_opts((*C.struct_SwrContext)(s), C.int64_t(ocl), (C.enum_AVSampleFormat)(osf), C.int(osr), C.int64_t(icl), (C.enum_AVSampleFormat)(isf), C.int(isr), C.int(lo), unsafe.Pointer(&lc)))
}

//Context destructor functions. Free the given Context and set the pointer to NULL.
func (s *Context) SwrFree() {
	C.swr_free((**C.struct_SwrContext)(unsafe.Pointer(s)))
}

//Closes the context so that swr_is_initialized() returns 0.
func (s *Context) SwrClose() {
	C.swr_close((*C.struct_SwrContext)(s))
}

//Core conversion functions. Convert audio
func (s *Context) SwrConvert(out **uint8, oc int, in **uint8, ic int) int {
	return int(C.swr_convert((*C.struct_SwrContext)(s), (**C.uint8_t)(unsafe.Pointer(out)), C.int(oc), (**C.uint8_t)(unsafe.Pointer(in)), C.int(ic)))
}

//Convert the next timestamp from input to output timestamps are in 1/(in_sample_rate * out_sample_rate) units.
func (s *Context) SwrNextPts(pts int64) int64 {
	return int64(C.swr_next_pts((*C.struct_SwrContext)(s), C.int64_t(pts)))
}

//Low-level option setting functions
//These functons provide a means to set low-level options that is not possible with the AvOption API.
//Activate resampling compensation ("soft" compensation).
func (s *Context) SwrSetCompensation(sd, cd int) int {
	return int(C.swr_set_compensation((*C.struct_SwrContext)(s), C.int(sd), C.int(cd)))
}

//Set a customized input channel mapping.
func (s *Context) SwrSetChannelMapping(cm *int) int {
	return int(C.swr_set_channel_mapping((*C.struct_SwrContext)(s), (*C.int)(unsafe.Pointer(cm))))
}

//Set a customized remix matrix.
func (s *Context) SwrSetMatrix(m *int, t int) int {
	return int(C.swr_set_matrix((*C.struct_SwrContext)(s), (*C.double)(unsafe.Pointer(m)), C.int(t)))
}

//Sample handling functions. Drops the specified number of output samples.
func (s *Context) SwrDropOutput(c int) int {
	return int(C.swr_drop_output((*C.struct_SwrContext)(s), C.int(c)))
}

//Injects the specified number of silence samples.
func (s *Context) SwrInjectSilence(c int) int {
	return int(C.swr_inject_silence((*C.struct_SwrContext)(s), C.int(c)))
}

//Gets the delay the next input sample will experience relative to the next output sample.
func (s *Context) SwrGetDelay(b int64) int64 {
	return int64(C.swr_get_delay((*C.struct_SwrContext)(s), C.int64_t(b)))
}

//Frame based API. Convert the samples in the input Frame and write them to the output Frame.
func (s *Context) SwrConvertFrame(o, i *Frame) int {
	return int(C.swr_convert_frame((*C.struct_SwrContext)(s), (*C.struct_AVFrame)(o), (*C.struct_AVFrame)(i)))
}

//Configure or reconfigure the Context using the information provided by the AvFrames.
func (s *Context) SwrConfigFrame(o, i *Frame) int {
	return int(C.swr_config_frame((*C.struct_SwrContext)(s), (*C.struct_AVFrame)(o), (*C.struct_AVFrame)(i)))
}
