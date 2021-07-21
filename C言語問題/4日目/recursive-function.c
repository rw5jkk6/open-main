#include <stdio.h>


int rsum(int n){
  if (n==0){
    return 0;
  }else{
    return n + rsum(n-1);
  }
}

int main(){
  int x;
  x = rsum(4);
  printf("%d\n", x);  // 10
  return 0;
  
}
