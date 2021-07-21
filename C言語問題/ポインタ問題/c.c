#include <stdio.h>

int main(){
   int a = 10;
   int* c;
   int** d;
   c = &a;
   d = &c;
   printf("%d\n", **d);
}

