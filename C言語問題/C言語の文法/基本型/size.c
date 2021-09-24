#include <stdio.h>

int main(){
  
  printf("_Bool..%d\n", (int)sizeof(_Bool)); 
  printf("char..%d\n", (int)sizeof(char)); 
  printf("short..%d\n", (int)sizeof(short));   // short int
  printf("int..%d\n", (int)sizeof(int));
  printf("unsigned..%d\n", (int)sizeof(unsigned));  // unsigned int
  printf("long..%d\n", (int)sizeof(long));  // long int
  printf("long long..%d\n", (int)sizeof(long long));  // long long int
  printf("long double..%d\n", (int)sizeof(long double));  // long double
}

// _Bool..1
// char..1
// short..2
// int..4
// unsigned..4
// long..8
// long long..8
// long double..16
