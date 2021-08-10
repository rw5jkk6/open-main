/*
コマンドライン引数に与えないか、与えるかを条件分岐で分けている
- 標準入力
- ファイルで受けるか
C言語では ctrl + d で EOF(int型) を標準入力する 
*/

#include <stdio.h>
#include <stdlib.h>

void type_one_file(FILE *fp){
    int ch;
    
    // cf287のファイルを参照
    // int putchar(int)
    while ((ch = getc(fp)) != EOF){
        putchar(ch);
    }
}

int main(int argc, char **argv){
    if (argc == 1){
        // FILE *stdin
        type_one_file(stdin);
    }else{
        int i;
        FILE *fp;

        for (i =1; i <argc; i++){
            fp = fopen(argv[i], "rb");
            
            // エラー処理
            if (fp == NULL){
                fprintf(stderr, "%s:%s can not open.\n", argv[0], argv[i]);
                exit(1);
            }
            type_one_file(fp);
        }
    }
    return 0;
}