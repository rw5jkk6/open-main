#include <stdio.h>

int main(){
   int a = 100;
   int b = 200;

   int *c;
   int *d;

   c = &a;
   d = &b;

   a = *c + *d;
   d = c;
   a = *d + a;
   printf("%d\n", a);

}

