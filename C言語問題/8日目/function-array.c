#include <stdio.h>
#define NELEMS(a) (sizeof(a)/ sizeof((a)[0]))

void add1(int a[], int n){
    int i;

    for (i = 0; i < n; i++){
        a[i]++;
    }
    return ;
}

int main(){
    int a[] = {1, 2, 3, 4, 5};
    int i;

    for (i =0; i < NELEMS(a); i++){
        printf("%d ", a[i]);
    }
    printf("\n");
}