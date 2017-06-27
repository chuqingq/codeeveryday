#include<stdio.h>
#include"libmy.h"

int main() {
    GoString gs = Hello();
    printf("str: %s\n", gs.p);
    printf("n: %d\n", (int)gs.n);
    getchar();
    return 0;
}

