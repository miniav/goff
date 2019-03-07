#include <stdint.h>
#include <stdio.h>

typedef struct buffer_data {
  uint8_t *ptr;
  size_t size;
} buffer_data;

pthread_mutex_t buffer_in_mutex = PTHREAD_MUTEX_INITIALIZER;
pthread_mutex_t buffer_out_mutex = PTHREAD_MUTEX_INITIALIZER;

buffer_data buffer_in, buffer_out;

int main(int argc, char const *argv[]) {
  printf("%ld\n", buffer_in.size);
  return 0;
}
