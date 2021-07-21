#include <stdio.h>

int plus(double a, double b){
    return a + b;
}

int main(){
    double c = 1.0;
    int d = 2;

    printf("%d\n", plus(c, d));
}

// int plus(double a, int b)  <- ok
// int plus(double a, double b)  <- ok
// c言語は型が異なっても計算できるが、
// 他の言語(Go, swift)ではダメ
// rustに至ってはi32とi64でもダメ


