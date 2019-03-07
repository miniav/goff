package avformat

import (
	"log"
	"os"
	"testing"

	"github.com/tosone/goff1/avcodec"
	"github.com/tosone/goff1/avutil"
)

//func TestNewAVIOContext(t *testing.T) {
//	log.SetFlags(log.LstdFlags | log.Lshortfile)
//	Input()
//	RegisterAll()
//	_fmt := NewAVFormatContext()
//	buf := avutil.Malloc(4096)
//	_io := NewAVIOContext(buf, 4096)
//	_fmt.SetAVFormatAVIOContext(_io)
//	log.Println(_fmt.OpenInput("", nil, nil))
//	log.Printf("after open input: %+v, %+v\n", _fmt, *_fmt.streams)
//	log.Println(((*_fmt.streams).codec).codec_id)
//	_fmt.FindStreamInfo(nil)
//	log.Println("----------------------")
//	_fmt.DumpFormat()
//	log.Println("----------------------")
//	idx, _ := _fmt.AudioStreamIndex()
//	codecID := _fmt.AudioStreamCodecID(idx)
//	codec := avcodec.FindCodec(codecID)
//	_codecID := codec.GetID()
//	avcodec.NewParserContext(_codecID)
//	codecCtx := avcodec.NewContext(codec)
//	codecCtx.Open2(codec, nil)
//	packet := avcodec.NewPacket()
//	packet.InitPacket()
//
//	var err error
//	var f *os.File
//	if f, err = os.Create("test.raw"); err != nil {
//		log.Fatal(err)
//	}
//
//	for {
//		log.Println("got here")
//		if err = _fmt.ReadFrame(packet); err == avutil.Error(avutil.AVERROR_EOF) {
//			break
//		} else if err != nil {
//			log.Fatal(err)
//		}
//		if err := codecCtx.SendPacket(packet); err != nil {
//			log.Fatal(err)
//		}
//		for {
//			var frame = avutil.NewFrame()
//			if err := codecCtx.ReceiveFrame(frame); err != nil {
//				break
//			}
//			sf := codecCtx.SampleFmt()
//			var len int
//			if len, err = sf.BytesPerSample(); err != nil {
//				log.Fatal(err)
//			} else {
//				if _, err = f.Write(frame.ToBytes(codecCtx.Channels(), len)); err != nil {
//					log.Fatal(err)
//				}
//			}
//		}
//	}
//
//	f.Close()
//
//	_fmt.CloseInput()
//}

//func TestNewAVIOContext3(t *testing.T) {
//	log.SetFlags(log.LstdFlags | log.Lshortfile)
//	data, _ := ioutil.ReadFile("../test.mp3")
//	packet := avcodec.NewPacket()
//	codecctx := avcodec.NewContext(avcodec.FindCodec(avcodec.CodecIDMP3))
//	codecctx.Open2(avcodec.FindCodec(avcodec.CodecIDMP3), nil)
//	parser := avcodec.NewParserContext(int(avcodec.CodecIDMP3))
//	var err error
//	var f *os.File
//	if f, err = os.Create("test.raw"); err != nil {
//		log.Fatal(err)
//	}
//	var used int
//	for {
//		log.Printf("%+v\n", packet)
//		used, err = parser.Parser2(codecctx, packet, data, avutil.NOPTS_VALUE, avutil.NOPTS_VALUE, 0)
//		data = data[used:]
//		log.Printf("%+v\n", packet)
//		if err := codecctx.SendPacket(packet); err != nil {
//			log.Fatal(err)
//		}
//		for {
//			var frame = avutil.NewFrame()
//			if err := codecctx.ReceiveFrame(frame); err != nil {
//				break
//			}
//			sf := codecctx.SampleFmt()
//			var len int
//			if len, err = sf.BytesPerSample(); err != nil {
//				log.Fatal(err)
//			} else {
//				if _, err = f.Write(frame.ToBytes(codecctx.Channels(), len)); err != nil {
//					log.Fatal(err)
//				}
//			}
//		}
//	}
//	f.Close()
//}

func TestNewAVIOContext2(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	_fmt := NewAVFormatContext()
	log.Println(_fmt.OpenInput("../test.ogg", nil, nil))
	log.Printf("after open input: %+v, %+v\n", _fmt, *_fmt.streams)
	log.Println(((*_fmt.streams).codec).codec_id)
	_fmt.FindStreamInfo(nil)
	_fmt.DumpFormat()
	idx, _ := _fmt.AudioStreamIndex()
	codecID := _fmt.AudioStreamCodecID(idx)
	codec := avcodec.FindCodec(codecID)
	_codecID := codec.GetID()
	avcodec.NewParserContext(_codecID)
	codecCtx := avcodec.NewContext(codec)
	codecCtx.Open2(codec, nil)
	packet := avcodec.NewPacket()
	packet.InitPacket()
	var frame = avutil.NewFrame()

	var err error
	var f *os.File
	if f, err = os.Create("test.raw"); err != nil {
		log.Fatal(err)
	}

	for {
		if err = _fmt.ReadFrame(packet); err == avutil.Error(avutil.AVERROR_EOF) {
			if err := codecCtx.SendPacket(packet); err != nil {
				log.Fatal(err)
			}
			for {
				if err := codecCtx.ReceiveFrame(frame); err != nil {
					break
				}
				sf := codecCtx.SampleFmt()
				var len int
				if len, err = sf.BytesPerSample(); err != nil {
					log.Fatal(err)
				} else {
					if _, err = f.Write(frame.ToBytes(codecCtx.Channels(), len)); err != nil {
						log.Fatal(err)
					}
				}
			}
			break
		} else if err != nil {
			log.Fatal(err)
		}
		if err := codecCtx.SendPacket(packet); err == avutil.Error(avutil.EAGAIN) {
			continue
		} else if err != nil {
			log.Fatal(err)
		}
		for {
			if err := codecCtx.ReceiveFrame(frame); err == avutil.Error(avutil.EAGAIN) {
				break
			} else if err != nil {
				log.Fatal(err)
			}
			sf := codecCtx.SampleFmt()
			var len int
			if len, err = sf.BytesPerSample(); err != nil {
				log.Fatal(err)
			} else {
				if _, err = f.Write(frame.ToBytes(codecCtx.Channels(), len)); err != nil {
					log.Fatal(err)
				}
			}
		}
	}

	f.Close()
	_fmt.CloseInput()
}
