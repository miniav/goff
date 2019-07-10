package avcodec

import (
	"testing"
)

func TestFindCodec(t *testing.T) {
	var codec = FindCodec(CodecIDMPEG1VIDEO)
	t.Log("2CodecIDMPEG1VIDEO:", codec.GetID())
}
