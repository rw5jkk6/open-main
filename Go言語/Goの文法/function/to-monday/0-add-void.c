#include <stdio.h>
 
void add(int n, int m){
    int r;
    r = n + m;
    printf("%d\n", r);
}

int main(void){
    add(1, 2);
    return 0;
}
