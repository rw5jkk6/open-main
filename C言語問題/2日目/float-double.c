
#include <stdio.h>

int main(){
  float d=0.1;

  puts("float");
  printf("%f\n", d);
  printf("%.20f\n", d);


  puts("-------------------");
  puts("double");
  double e=0.1;

  double f =0.1+0.1+0.1+0.1+0.1+
    0.1+0.1+0.1+0.1+0.1;

  printf("%f\n", e);
  printf("%.20f\n", e);

  // 浮動小数点数はデフォルトでdouble
  if (f == 1.0){
    printf("Equal\n");
  }else{
    printf("no\n");
  }
}


// float
// 0.100000
// 0.10000000149011611938
// 0x1.99999ap-4

// -------------------
// double
// 0.100000
// 0.10000000000000000555
// 0x1.999999999999ap-4

// no
