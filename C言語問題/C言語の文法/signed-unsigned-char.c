#include <stdio.h>
#include <stdlib.h>

int main(void){
   unsigned char x = 0;

   while(1){
       printf("unsigned: %4d, signed: %4d\n", x, (char)x);
       if (x++ == 255) break;
   }
   exit(0);
}

/*
unsigned:    0, signed:    0
unsigned:    1, signed:    1
unsigned:    2, signed:    2

(省略)

unsigned:  253, signed:   -3
unsigned:  254, signed:   -2
unsigned:  255, signed:   -1
*/
