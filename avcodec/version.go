package avcodec

/*
#include <libavcodec/avcodec.h>
#include <libavutil/avutil.h>
#include <libavdevice/avdevice.h>
#include <libavfilter/avfilter.h>
#include <libavformat/avformat.h>
#include <libswresample/swresample.h>
#include <libswscale/swscale.h>
*/
import "C"
import "log"

// Version ..
func Version() {
	log.Printf("%+v", C.avcodec_version())
	log.Printf("%+v", C.swresample_version())
	log.Printf("%+v", C.avutil_version())
	log.Printf("%+v", C.avdevice_version())
	log.Printf("%+v", C.avfilter_version())
	log.Printf("%+v", C.swscale_version())
	log.Printf("%+v", C.avformat_version())
}

