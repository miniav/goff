#include <stdio.h>

int main(int argc, char const *argv[]) {
  int i = 5;
  int *j = &i;
  int **k = &j;
  printf("%p %p %p\n", &i, j, k);
  return 0;
}
