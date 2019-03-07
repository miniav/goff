#include <stdio.h>
#include <stdlib.h>

void functionA() {
  printf("This is functionA\n");
}
void functionB() {
  printf("This is functionB\n");
}

int main() {
  /* register the termination function */
  atexit(functionA);
  atexit(functionB);
  printf("Starting  main program...\n");
  printf("Exiting main program...\n");
  return (0);
}
