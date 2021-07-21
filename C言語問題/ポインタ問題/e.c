#include <stdio.h>

int main(){
   int a = 10;
   int* b;
   int* c;

   b = &a;
   c = &*b;
   printf("%d\n", *c);
}

