package avutil

// #include <libavutil/rational.h>
import "C"
import "unsafe"

// Rational Rational number (pair of numerator and denominator).
type Rational C.AVRational

// RationalNew Create an AVRational.
func RationalNew(num, den int) Rational {
	return Rational(C.av_make_q(C.int(num), C.int(den)))
}

// RationalCmp Compare two rationals.
func RationalCmp(a, b Rational) int {
	return int(C.av_cmp_q(C.AVRational(a), C.AVRational(b)))
}

// Rational2Double Convert an AVRational to a `double`.
func Rational2Double(a Rational) float64 {
	return float64(C.av_q2d(C.AVRational(a)))
}

// Double2Rational Convert a double precision floating point number to a rational.
// In case of infinity, the returned value is expressed as `{1, 0}` or `{-1, 0}` depending on the sign.
func Double2Rational(d float64, max int) Rational {
	return Rational(C.av_d2q(C.double(d), C.int(max)))
}

// Reduce Reduce a fraction.
func Reduce(n, d, max int64) (num, den int) {
	C.av_reduce((*C.int)(unsafe.Pointer(&num)), (*C.int)(unsafe.Pointer(&den)), C.int64_t(n), C.int64_t(d), C.int64_t(max))
	return
}

// RationalMul Multiply two rationals.
func RationalMul(a, b Rational) Rational {
	return Rational(C.av_mul_q(C.AVRational(a), C.AVRational(b)))
}

// RationalDiv Divide one rational by another.
func RationalDiv(a, b Rational) Rational {
	return Rational(C.av_div_q(C.AVRational(a), C.AVRational(b)))
}

// RationalAdd Add two rationals.
func RationalAdd(a, b Rational) Rational {
	return Rational(C.av_add_q(C.AVRational(a), C.AVRational(b)))
}

// RationalSub Subtract one rational from another.
func RationalSub(a, b Rational) Rational {
	return Rational(C.av_sub_q(C.AVRational(a), C.AVRational(b)))
}

// RationalInv Invert a rational.
func RationalInv(q Rational) Rational {
	return Rational(C.av_inv_q(C.AVRational(q)))
}

// RationalNearer Find which of the two rationals is closer to another rational.
// 1 if `q1` is nearer to `q` than `q2`
// -1 if `q2` is nearer to `q` than `q1`
// 0 if they have the same distance
func RationalNearer(q, q1, q2 Rational) int {
	return int(C.av_nearer_q(C.AVRational(q), C.AVRational(q1), C.AVRational(q2)))
}

// RationalNearestID Find the value in a list of rationals nearest a given reference rational.
// Index of the nearest value found in the array
func RationalNearestID(q Rational, list []Rational) int {
	return int(C.av_find_nearest_q_idx(C.AVRational(q), (*C.AVRational)(unsafe.Pointer(&list[0]))))
}

// Rational2IntFloat Convert an AVRational to a IEEE 32-bit `float` expressed in fixed-point format.
func Rational2IntFloat(q Rational) uint32 {
	return uint32(C.av_q2intfloat(C.AVRational(q)))
}
