package avutil

// #include <libavutil/avutil.h>
// #include <libavutil/opt.h>
import "C"

type SampleFmt C.enum_AVSampleFormat

// SampleFmtNone None
var SampleFmtNone = SampleFmt(C.AV_SAMPLE_FMT_NONE)

// SampleFmtU8 unsigned 8 bits
var SampleFmtU8 = SampleFmt(C.AV_SAMPLE_FMT_U8)

// SampleFmtS16 signed 16 bits
var SampleFmtS16 = SampleFmt(C.AV_SAMPLE_FMT_S16)

// SampleFmtS32 signed 32 bits
var SampleFmtS32 = SampleFmt(C.AV_SAMPLE_FMT_S32)

// SampleFmtFLT float
var SampleFmtFLT = SampleFmt(C.AV_SAMPLE_FMT_FLT)

// SampleFmtDBL double
var SampleFmtDBL = SampleFmt(C.AV_SAMPLE_FMT_DBL)

// SampleFmtU8P unsigned 8 bits, planar
var SampleFmtU8P = SampleFmt(C.AV_SAMPLE_FMT_U8P)

// SampleFmtS16P signed 16 bits, planar
var SampleFmtS16P = SampleFmt(C.AV_SAMPLE_FMT_S16P)

// SampleFmtS32P signed 32 bits, planar
var SampleFmtS32P = SampleFmt(C.AV_SAMPLE_FMT_S32P)

// SampleFmtFLTP float, planar
var SampleFmtFLTP = SampleFmt(C.AV_SAMPLE_FMT_FLTP)

// SampleFmtDBLP double, planar
var SampleFmtDBLP = SampleFmt(C.AV_SAMPLE_FMT_DBLP)

// SampleFmtNB Number of sample formats. DO NOT USE if linking dynamically.
var SampleFmtNB = SampleFmt(C.AV_SAMPLE_FMT_NB)

// SampleFmtIsPlanar ..
func (f SampleFmt) IsPlanar() (isPlanar bool) {
	if code := int(C.av_sample_fmt_is_planar(C.enum_AVSampleFormat(f))); code == 1 {
		isPlanar = true
	}
	return
}

func (f SampleFmt) BytesPerSample() (len int, err error) {
	if len = int(C.av_get_bytes_per_sample(C.enum_AVSampleFormat(f))); len < 0 {
		err = Error(len)
	}
	return
}
