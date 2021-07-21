/*
値を入れずにアドレスだけ取得しておきたい時
*/


#include <stdio.h>

int main(){
    int *a;
    int *b = NULL;
    int c;
    //int d = NULL; error

    printf("%p\n", a);
    printf("%p\n", b);
    printf("%p\n", &c);
    //printf("%p\n", &d);

    return 0;
}

// 0x7ffee1f51780
// 0x0
// 0x7ffee1f51754
