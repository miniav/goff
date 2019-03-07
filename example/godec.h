#include <pthread.h>
#include <stdbool.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_BUFFER_SIZE 10240 // buffer 每次缓存的最大长度

typedef struct buffer_data { // c 缓存数据的结构体
  uint8_t *ptr;
  size_t size;
} buffer_data;

buffer_data buffer_in; // 传入的数据对象

pthread_mutex_t buffer_in_mutex = PTHREAD_MUTEX_INITIALIZER;

// 初始化传入的 buffer 的内存对象
buffer_data init_buffer() {
  buffer_data buffer_in; // 传入的数据对象
  buffer_in.ptr  = malloc(MAX_BUFFER_SIZE * sizeof(uint8_t));
  buffer_in.size = 0;
  return buffer_in;
}

// 从 Golang 中传入数据到 c 的内存中，返回每次读取的数据的数量
// 鉴于内存中不可以缓存过多的数据，也是为了节省内存，那么就需要每次仅将 buffer 填充的一定长度即可
int buffer_append(buffer_data *buffer, uint8_t *buf, int buf_size) {
  if (buffer->size == MAX_BUFFER_SIZE) {
    return 0;
  }
  pthread_mutex_lock(&buffer_in_mutex);
  if (MAX_BUFFER_SIZE - buffer->size > buf_size) {
    memcpy(buffer->ptr + buffer->size, buf, buf_size);
    buffer->size += buf_size;
    pthread_mutex_unlock(&buffer_in_mutex); // 解锁线程
    return buf_size;
  }
  memcpy(buffer->ptr + buffer->size, buf, MAX_BUFFER_SIZE - buffer->size);
  int read     = MAX_BUFFER_SIZE - buffer->size;
  buffer->size = MAX_BUFFER_SIZE;
  pthread_mutex_unlock(&buffer_in_mutex); // 解锁线程
  return read;
}

int buffer_read(buffer_data *buffer, uint8_t *buf, int buf_size) {
  if (buf_size == 0) {
    return 0;
  }
  pthread_mutex_lock(&buffer_in_mutex);
  if (buf_size >= buffer->size) {
    int read = buffer->size;
    memcpy(buf, buffer->ptr, buffer->size);
    buffer->size = 0;
    pthread_mutex_unlock(&buffer_in_mutex); // 解锁线程
    return read;
  }
  memcpy(buf, buffer->ptr, buf_size); 
  memmove(buffer->ptr,buffer->ptr+buf_size,buffer->size-buf_size);
  buffer->size -= buf_size;
  pthread_mutex_unlock(&buffer_in_mutex); // 解锁线程
  return buf_size;
}
