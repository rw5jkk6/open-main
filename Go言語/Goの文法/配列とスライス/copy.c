#include <stdio.h>

void copy()

void copy(char *dst, char *src){
    while (*src != '\0'){
        *dst++ = *src++;
    }
    *dst = '\0';
}

int main(void){
    char s1[] = "Hello world";
    char s2[1024];

    // copy(&s2, &s1);  これだとエラーになる
    copy(s2, s1);
    printf("%s\n", s2);
}