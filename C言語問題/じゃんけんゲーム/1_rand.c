/*
じゃんけんで相手側がランダムな手を出すため、ランダム関数を使う
グー、チョキ、パーの３手あるので、０から２までを出すようにする
*/


// #include <time.h>
// time_t time(time_t *timer);


// #include <stdlib.h>
// void srand(unsigned seed);

#include <stdio.h>
#include <stdlib.h>
#include <time.h>

int main(){
    int a, b, c;
    srand((unsigned)time(NULL));
    a = rand() % 3;
    b = rand() % 3;
    c = rand() % 3;
    
    printf("%d\n",a);
    printf("%d\n",b);
    printf("%d\n",c);
    
    return 0;    
}
