#include <stdio.h>
 
int main(void){
   
   int i;
   printf("%d\n", i);      // 364556325 適当な数字が出てくる
   printf("%d\n", i + 1);  // 364556326
   
   int *ii;
   printf("%p\n", ii);     //  0x7ffee0eb2780
   printf("%d\n", *ii);    // 0

   int *iii=NULL;
   printf("%p\n", iii);    // 0x0
   printf("%d\n", *iii);   // Segmentation fault: 11
   
}
