package avcodec

// #include <packet.h>
import "C"
import (
	"errors"
	"unsafe"

	"github.com/miniav/goff/avutil"
)

type Packet C.AVPacket

func NewPacket() *Packet {
	return (*Packet)(C.av_packet_alloc())
}

// InitPacket ..
func (p *Packet) InitPacket() {
	C.av_init_packet((*C.AVPacket)(unsafe.Pointer(p)))
}

// NewPacket ..
func (p *Packet) NewPacket(size int) {
	C.av_new_packet((*C.AVPacket)(unsafe.Pointer(p)), C.int(size))
}

// NewPacket ..
func (p *Packet) Zero() {
	C.av_packet_zero((*C.AVPacket)(unsafe.Pointer(p)))
}

// FreePacket ..
func (p *Packet) FreePacket() {
	C.av_packet_free_personal((*C.AVPacket)(unsafe.Pointer(p)))
}

// Shrink ..
func (p *Packet) Shrink(size int) {
	C.av_shrink_packet((*C.AVPacket)(unsafe.Pointer(p)), C.int(size))
}

// Grow ..
func (p *Packet) Grow(size int) {
	C.av_grow_packet((*C.AVPacket)(unsafe.Pointer(p)), C.int(size))
}

// FromData ..
func (p *Packet) FromData(data []byte) (err error) {
	if len(data) == 0 {
		err = errors.New("zero data")
		return
	}
	if code := int(C.av_packet_from_data(
		(*C.AVPacket)(unsafe.Pointer(p)),
		(*C.uchar)(unsafe.Pointer(&data[0])),
		C.int(len(data)),
	)); code < 0 {
		err = errors.New(avutil.StrError(code))
	}
	return
}

// Unref Clean the packet data
func (p *Packet) Unref() {
	C.av_packet_unref((*C.AVPacket)(unsafe.Pointer(p)))
}
