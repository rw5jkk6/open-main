// for文とif文で以下の文を表示するシェルスクリプトを作る

// 1は奇数です。
// 2は偶数です。  
// 3は奇数です。 
// 4は偶数です。
// 5は奇数です。 
// 6は偶数です。 
// 7は奇数です。 
// 8は偶数です。 
// 9は奇数です。 
// 10は偶数です

#include <stdio.h>

int main(){
    for (int i =0; i < 10; i++){
        if (i % 2 == 0){
            printf("%dは偶数です\n", i);
        }else{
            printf("%dは奇数です\n", i);
        }
    }
}

