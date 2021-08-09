#include <stdio.h>

void copy(char *dst, char *src){
    while (*src != '\0')
        *dst++ = *src++;
    *dst = '\0';
    return ;
}

int main(void){
    char s1[] = "Hello world";
    char s2[1024];

    copy(s2, s1);
    printf("%s\n", s2);
    return 0;
}