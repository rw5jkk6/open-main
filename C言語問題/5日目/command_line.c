#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[]){
    int sum = 0;
    int count = 0;

    if(argc < 3){
        printf("Error\n");
        return 1;
    }
    for (count=1; count<argc; count++){
        sum += atoi(argv[count]);
    }
    printf("%d\n", sum);
    return 0;
}
