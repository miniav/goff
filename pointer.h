#include <stdio.h>
#include <stdlib.h>

void print2(int **a) { printf("%d %p %p\n", **a, *a, a); }
void print1(int *a) {
  printf("%d %p %p\n", *a, a, &a);
  print2(&a);
}
