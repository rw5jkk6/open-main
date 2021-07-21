/*
文字は７文字まで入力できる。最後の一つは終末文字(\0)なので、
それ以上入力するとエラーになる。
ここではアルファベットなので7文字でOKだが、
日本語なら2文字ぐらいしか入らない

下のでもできるはずだが、macではできない。
printf("valのアドレス方法　2: %d\n", (int)&val);
*/


#include <stdio.h>

int main(){
    int val;
    int *ptr;
    val = 123;
    printf("valに格納されている値: %d\n", val);
    printf("valのアドレス方法　1: %p\n", &val);
    return 0;    
}



