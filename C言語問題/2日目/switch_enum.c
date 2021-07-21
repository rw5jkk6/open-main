#include <stdio.h>


typedef enum week{
    Mon,
    Tue,
    Wed,
    Thu,
    Fri,
    Sat,
    Sun
} week;

int main(){
    week wk;
    wk = Wed;

    switch (wk)
    {
    case Mon:
    case Tue:
    case Wed:
    case Thu:
    case Fri:
        printf("平日です\n");
        break;
    case Sat:
    case Sun:
        printf("休日です\n");
        break;
    default:
        break;
    }
    return 0;
}

// 平日です
