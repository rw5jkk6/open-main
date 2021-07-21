#include <stdio.h>

int main(){
    int flag = 0;
    char buf[4];

    printf("4桁の数字を入れてください\n");
    scanf("%s", buf);
    if (flag != 0){
        printf("Success\n");
    }else{
        printf("failed\n");
    }
    return 0;
}

//   stack
// __________
// | buf[0] |
// |________|
// | buf[1] |
// |________|
// | buf[2] |
// |________|
// | buf[3] |
// |________|
// | flag 0 |
// |________|
// |        |
// |________|
// |        |
//_|________|_____segmentaion__
