package avutil

// #include <error.h>
import "C"
import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"unsafe"
)

type errorCode struct {
	err  error
	code int
}

var (
	EAGAIN                    = int(C.averror(C.EAGAIN))
	AVERROR_BSF_NOT_FOUND     = int(C.AVERROR_BSF_NOT_FOUND)
	AVERROR_BUG               = int(C.AVERROR_BUG)
	AVERROR_BUFFER_TOO_SMALL  = int(C.AVERROR_BUFFER_TOO_SMALL)
	AVERROR_DECODER_NOT_FOUND = int(C.AVERROR_DECODER_NOT_FOUND)
	AVERROR_DEMUXER_NOT_FOUND = int(C.AVERROR_DEMUXER_NOT_FOUND)
	AVERROR_ENCODER_NOT_FOUND = int(C.AVERROR_ENCODER_NOT_FOUND)
	AVERROR_EOF               = int(C.AVERROR_EOF)
	AVERROR_EXIT              = int(C.AVERROR_EXIT)
	AVERROR_EXTERNAL          = int(C.AVERROR_EXTERNAL)
	AVERROR_FILTER_NOT_FOUND  = int(C.AVERROR_FILTER_NOT_FOUND)
	AVERROR_INVALIDDATA       = int(C.AVERROR_INVALIDDATA)
	AVERROR_MUXER_NOT_FOUND   = int(C.AVERROR_MUXER_NOT_FOUND)
	AVERROR_OPTION_NOT_FOUND  = int(C.AVERROR_OPTION_NOT_FOUND)
	AVERROR_PATCHWELCOME      = int(C.AVERROR_PATCHWELCOME)
	//AVERROR_PROTOCOL_NOT_FOUN = int(C.AVERROR_PROTOCOL_NOT_FOUN)
	AVERROR_STREAM_NOT_FOUND  = int(C.AVERROR_STREAM_NOT_FOUND)
	AVERROR_BUG2              = int(C.AVERROR_BUG2)
	AVERROR_UNKNOWN           = int(C.AVERROR_UNKNOWN)
	AVERROR_EXPERIMENTAL      = int(C.AVERROR_EXPERIMENTAL)
	AVERROR_INPUT_CHANGED     = int(C.AVERROR_INPUT_CHANGED)
	AVERROR_OUTPUT_CHANGED    = int(C.AVERROR_OUTPUT_CHANGED)
	AVERROR_HTTP_BAD_REQUEST  = int(C.AVERROR_HTTP_BAD_REQUEST)
	AVERROR_HTTP_UNAUTHORIZED = int(C.AVERROR_HTTP_UNAUTHORIZED)
	AVERROR_HTTP_FORBIDDEN    = int(C.AVERROR_HTTP_FORBIDDEN)
	AVERROR_HTTP_NOT_FOUND    = int(C.AVERROR_HTTP_NOT_FOUND)
	AVERROR_HTTP_OTHER_4XX    = int(C.AVERROR_HTTP_OTHER_4XX)
	AVERROR_HTTP_SERVER_ERROR = int(C.AVERROR_HTTP_SERVER_ERROR)
)

type AVERRORS = map[int]error

var AVERROR = AVERRORS{
	EAGAIN:                    errors.New("again"),
	AVERROR_BSF_NOT_FOUND:     errors.New(StrError(int(C.AVERROR_BSF_NOT_FOUND))),
	AVERROR_BUG:               errors.New(StrError(int(C.AVERROR_BUG))),
	AVERROR_BUFFER_TOO_SMALL:  errors.New(StrError(int(C.AVERROR_BUFFER_TOO_SMALL))),
	AVERROR_DECODER_NOT_FOUND: errors.New(StrError(int(C.AVERROR_DECODER_NOT_FOUND))),
	AVERROR_DEMUXER_NOT_FOUND: errors.New(StrError(int(C.AVERROR_DEMUXER_NOT_FOUND))),
	AVERROR_ENCODER_NOT_FOUND: errors.New(StrError(int(C.AVERROR_ENCODER_NOT_FOUND))),
	AVERROR_EOF:               errors.New(StrError(int(C.AVERROR_EOF))),
	AVERROR_EXIT:              errors.New(StrError(int(C.AVERROR_EXIT))),
	AVERROR_EXTERNAL:          errors.New(StrError(int(C.AVERROR_EXTERNAL))),
	AVERROR_FILTER_NOT_FOUND:  errors.New(StrError(int(C.AVERROR_FILTER_NOT_FOUND))),
	AVERROR_INVALIDDATA:       errors.New(StrError(int(C.AVERROR_INVALIDDATA))),
	AVERROR_MUXER_NOT_FOUND:   errors.New(StrError(int(C.AVERROR_MUXER_NOT_FOUND))),
	AVERROR_OPTION_NOT_FOUND:  errors.New(StrError(int(C.AVERROR_OPTION_NOT_FOUND))),
	AVERROR_PATCHWELCOME:      errors.New(StrError(int(C.AVERROR_PATCHWELCOME))),
	//AVERROR_PROTOCOL_NOT_FOUN: errors.New(StrError(int(C.AVERROR_PROTOCOL_NOT_FOUN))),
	AVERROR_STREAM_NOT_FOUND:  errors.New(StrError(int(C.AVERROR_STREAM_NOT_FOUND))),
	AVERROR_BUG2:              errors.New(StrError(int(C.AVERROR_BUG2))),
	AVERROR_UNKNOWN:           errors.New(StrError(int(C.AVERROR_UNKNOWN))),
	AVERROR_EXPERIMENTAL:      errors.New(StrError(int(C.AVERROR_EXPERIMENTAL))),
	AVERROR_INPUT_CHANGED:     errors.New(StrError(int(C.AVERROR_INPUT_CHANGED))),
	AVERROR_OUTPUT_CHANGED:    errors.New(StrError(int(C.AVERROR_OUTPUT_CHANGED))),
	AVERROR_HTTP_BAD_REQUEST:  errors.New(StrError(int(C.AVERROR_HTTP_BAD_REQUEST))),
	AVERROR_HTTP_UNAUTHORIZED: errors.New(StrError(int(C.AVERROR_HTTP_UNAUTHORIZED))),
	AVERROR_HTTP_FORBIDDEN:    errors.New(StrError(int(C.AVERROR_HTTP_FORBIDDEN))),
	AVERROR_HTTP_NOT_FOUND:    errors.New(StrError(int(C.AVERROR_HTTP_NOT_FOUND))),
	AVERROR_HTTP_OTHER_4XX:    errors.New(StrError(int(C.AVERROR_HTTP_OTHER_4XX))),
	AVERROR_HTTP_SERVER_ERROR: errors.New(StrError(int(C.AVERROR_HTTP_SERVER_ERROR))),
}

func init() {
	// log.Println(AVERROR)
	var code = errorCode{fmt.Errorf("123"), EAGAIN}
	code.err = Error(code.code)
	fmt.Println(code)
}

// StrError get error description
func StrError(e int) (str string) {
	o := make([]byte, 64)
	C.av_strerror(C.int(e), (*C.char)(unsafe.Pointer(&o[0])), C.size_t(len(o)))
	o = bytes.Trim(o, "\x00")
	str = string(o)
	return
}

// Error ..
func Error(e int) (err error) {
	var ok bool
	if err, ok = AVERROR[e]; !ok {
		log.Println(fmt.Sprintf("no such a error code: %d, %s", e, StrError(e)))
		err = errors.New(StrError(e))
	}
	return
}
