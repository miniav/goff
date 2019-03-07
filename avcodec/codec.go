package avcodec

// #include <codec.h>
import "C"

type Codec C.AVCodec

// FindCodec ..
func FindCodec(id CodecID) *Codec {
	return (*Codec)(C.avcodec_find_decoder(C.enum_AVCodecID(id)))
}

func (c *Codec) GetID() int {
	return int(C.get_id_from_codec((*C.AVCodec)(c)))
}
