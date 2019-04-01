#include<stdio.h>

#include"libmy.h"
#include"libmyadd.h"

int main() {
    GoString gs = Hello();
    printf("str: %s\n", gs.p);
    printf("n: %d\n", (int)gs.n);

    GoInt a = 1;
    GoInt b = 2;
    GoInt r = Add(a, b);
    printf("add: %d\n", r);

    getchar();
    return 0;
}

