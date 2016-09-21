#include<stdio.h>
#include"lib.h"

int main() {
    GoString gs = Hello();
    printf("str: %s\n", gs.p);
    printf("n: %d\n", (int)gs.n);
    return 0;
}

