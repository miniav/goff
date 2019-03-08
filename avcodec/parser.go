package avcodec

// #include <parser.h>
import "C"
import (
	"errors"
	"unsafe"

	"github.com/tosone/goff/avutil"
)

type ParserContext C.AVCodecParserContext

type Parser C.AVCodecParser

// NewParserContext ..
func NewParserContext(codecID int) *ParserContext {
	return (*ParserContext)(C.av_parser_init(C.int(codecID)))
}

// Parser2 ..
func (p *ParserContext) Parser2(ctx *Context, packet *Packet, buf []byte, pts, dts, pos int64) (read int, err error) {
	if len(buf) == 0 {
		err = errors.New("input buffer is null")
		return
	}
	var c_dict = (*C.AVDictionary)(unsafe.Pointer(packet))
	var d_dict = (**C.AVDictionary)(C.malloc(C.ulong(unsafe.Sizeof(c_dict))))
	*d_dict = c_dict
	if read = int(C.av_parser_parse2_personal(
		(*C.AVCodecParserContext)(p),
		(*C.AVCodecContext)(ctx),
		(*C.AVPacket)(packet),
		(*C.uchar)(unsafe.Pointer(&buf[0])),
		C.int(len(buf)),
		C.long(pts),
		C.long(dts),
		C.long(pos),
	)); read < 0 {
		err = errors.New(avutil.StrError(read))
	}
	return
}

// Close ..
func (p *ParserContext) Close() {
	C.av_parser_close((*C.AVCodecParserContext)(p))
}
