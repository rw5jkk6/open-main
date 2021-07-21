#include <stdio.h>

int main(){
    int i = 100;
    int j = 200;
    
    int *p1;
    int *p2;
    
    p1 = &i;
    p2 = &j;
    
    i = *p1 + *p2;
    p2 = p1;
    printf("%d\n", *p1);
    j = *p2 + i;
    printf("%d\n", j);
    return 0;
}

// 300
// 600
