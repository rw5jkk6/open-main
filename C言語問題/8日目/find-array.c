/*
新しいc言語の教科書　10-2

配列の中から、指定した数字を見つける関数
ここでは４を見つけると、foundを返してくる
*/


#include <stdio.h>

int *find(int, int []);

int *find(int n, int a[]){
    int i;

    for (i = 0; a[i] != -1; i++){
        if (a[i] == n)
            break;
    }
    return &a[i];
}

int main(){
    int a[] = {2, 5, 9, 3, 4, 5, 1, 6, -1 };
    int *p;

    p = find(4, a);
    if (*p == -1)
        printf("Not found\n");
    else    
        printf("Found\n");
    return 0;
}
