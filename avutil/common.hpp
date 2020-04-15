#include <libavutil/common.h>

int ffmax(int a, int b) {
    return FFMAX(a, b);
}

int ffmax3(int a, int b, int c) {
    return FFMAX3(a, b, c);
}

int ffmin(int a, int b) {
    return FFMIN(a, b);
}

int ffmin3(int a, int b, int c) {
    return FFMIN3(a, b, c);
}

int ffabs(int a) {
    return FFABS(a);
}