package avutil

// #include <stdlib.h>
// #include <libavutil/dict.h>
import "C"
import "unsafe"

type (
	Dictionary      C.struct_AVDictionary
	DictionaryEntry C.struct_AVDictionaryEntry
)

const (
	AV_DICT_MATCH_CASE      = int(C.AV_DICT_MATCH_CASE)
	AV_DICT_IGNORE_SUFFIX   = int(C.AV_DICT_IGNORE_SUFFIX)
	AV_DICT_DONT_STRDUP_KEY = int(C.AV_DICT_DONT_STRDUP_KEY)
	AV_DICT_DONT_STRDUP_VAL = int(C.AV_DICT_DONT_STRDUP_VAL)
	AV_DICT_DONT_OVERWRITE  = int(C.AV_DICT_DONT_OVERWRITE)
	AV_DICT_APPEND          = int(C.AV_DICT_APPEND)
	AV_DICT_MULTIKEY        = int(C.AV_DICT_MULTIKEY)
)

func (d *Dictionary) AvDictGet(key string, prev *DictionaryEntry, flags int) *DictionaryEntry {
	var ckey = C.CString(key)
	defer C.free(unsafe.Pointer(ckey))

	return (*DictionaryEntry)(C.av_dict_get(
		(*C.struct_AVDictionary)(d),
		ckey,
		(*C.struct_AVDictionaryEntry)(prev),
		C.int(flags),
	))
}

func (d *Dictionary) AvDictCount() int {
	return int(C.av_dict_count((*C.struct_AVDictionary)(d)))
}

func (d *Dictionary) AvDictSet(key, value string, flags int) int {
	var ckey = C.CString(key)
	defer C.free(unsafe.Pointer(ckey))

	var cvalue = C.CString(value)
	defer C.free(unsafe.Pointer(cvalue))

	return int(C.av_dict_set(
		(**C.struct_AVDictionary)(unsafe.Pointer(&d)),
		ckey,
		cvalue,
		C.int(flags),
	))
}

func (d *Dictionary) AvDictSetInt(key string, value int64, flags int) int {
	Ckey := C.CString(key)
	defer C.free(unsafe.Pointer(Ckey))

	return int(C.av_dict_set_int(
		(**C.struct_AVDictionary)(unsafe.Pointer(&d)),
		Ckey,
		C.int64_t(value),
		C.int(flags),
	))
}

func (d *Dictionary) AvDictParseString(str, key_val_sep, pairs_sep string, flags int) int {
	var cstr = C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	var ckey_val_sep = C.CString(key_val_sep)
	defer C.free(unsafe.Pointer(ckey_val_sep))

	var cpairs_sep = C.CString(pairs_sep)
	defer C.free(unsafe.Pointer(cpairs_sep))

	return int(C.av_dict_parse_string(
		(**C.struct_AVDictionary)(unsafe.Pointer(&d)),
		cstr,
		ckey_val_sep,
		cpairs_sep,
		C.int(flags),
	))
}

func (d *Dictionary) AvDictCopy(src *Dictionary, flags int) int {
	return int(C.av_dict_copy(
		(**C.struct_AVDictionary)(unsafe.Pointer(&d)),
		(*C.struct_AVDictionary)(unsafe.Pointer(src)),
		C.int(flags),
	))
}

func (d *Dictionary) AvDictFree() {
	C.av_dict_free((**C.struct_AVDictionary)(unsafe.Pointer(&d)))
}

func (d *Dictionary) AvDictGetString(key_val_sep, pairs_sep byte) (int, string) {
	var Cbuf *C.char

	ret := int(C.av_dict_get_string(
		(*C.struct_AVDictionary)(d),
		(**C.char)(&Cbuf),
		C.char(key_val_sep),
		C.char(pairs_sep),
	))

	var buf string
	if ret == 0 {
		buf = C.GoString(Cbuf)
		C.free(unsafe.Pointer(Cbuf))
	}

	return ret, buf
}

func (d *DictionaryEntry) Key() string {
	return C.GoString(d.key)
}

func (d *DictionaryEntry) Value() string {
	return C.GoString(d.value)
}
