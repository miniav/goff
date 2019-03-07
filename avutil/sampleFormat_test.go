package avutil

import (
	"log"
	"testing"
)

func TestBytesPerSample(t *testing.T) {
	// log.Println("Byte per sample format none:", SampleFmtNone.BytesPerSample())
	// log.Println("Byte per sample format u8:", SampleFmtU8.BytesPerSample())
	// log.Println("Byte per sample format s16:", SampleFmtS16.BytesPerSample())
	// log.Println("Byte per sample format s32:", SampleFmtS32.BytesPerSample())
	// log.Println("Byte per sample format float:", SampleFmtFLT.BytesPerSample())
	// log.Println("Byte per sample format double:", SampleFmtDBL.BytesPerSample())
	// log.Println("Byte per sample format u8 planar:", SampleFmtU8P.BytesPerSample())
	// log.Println("Byte per sample format s16 planar:", SampleFmtS16P.BytesPerSample())
	// log.Println("Byte per sample format s32 planar:", SampleFmtS32P.BytesPerSample())
	// log.Println("Byte per sample format float planar:", SampleFmtFLTP.BytesPerSample())
	// log.Println("Byte per sample format double planar:", SampleFmtDBLP.BytesPerSample())
}

func TestSampleFmtIsPlanar(t *testing.T) {
	log.Printf("Sample format is planar none: %v\n", SampleFmtNone.IsPlanar())
	log.Printf("Sample format is planar u8: %v\n", SampleFmtU8.IsPlanar())
	log.Printf("Sample format is planar s16: %v\n", SampleFmtS16.IsPlanar())
	log.Printf("Sample format is planar s32: %v\n", SampleFmtS32.IsPlanar())
	log.Printf("Sample format is planar float: %v\n", SampleFmtFLT.IsPlanar())
	log.Printf("Sample format is planar double: %v\n", SampleFmtDBL.IsPlanar())
	log.Printf("Sample format is planar u8 planar: %v\n", SampleFmtU8P.IsPlanar())
	log.Printf("Sample format is planar s16 planar: %v\n", SampleFmtS16P.IsPlanar())
	log.Printf("Sample format is planar s32 planar: %v\n", SampleFmtS32P.IsPlanar())
	log.Printf("Sample format is planar float planar: %v\n", SampleFmtFLTP.IsPlanar())
	log.Printf("Sample format is planar double planar: %v\n", SampleFmtDBLP.IsPlanar())
}
