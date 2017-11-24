#include <stdio.h>
#include <time.h>
#include <unistd.h>

static unsigned long long nstime(void) {
    struct timespec ts;
    clock_gettime(CLOCK_REALTIME, &ts);
    return ((unsigned long long)ts.tv_sec)*1e9 + ts.tv_nsec;
}

int count = 100000000;
int main() {
    unsigned long long t1, t2, t3;
    t1 = nstime();
    for (int i = 0; i < count; i++) {
        t2 = nstime();
    }
    t3 = nstime();
    printf("time diff: %llu\n", t3-t1);
    return 0;
}

/*
chuqq@chuqq-hp:~/temp/codeeveryday/c/20171008_ustime$ ./a.out 
time diff: 2142726912
chuqq@chuqq-hp:~/temp/codeeveryday/c/20171008_ustime$ time ./a.out 
time diff: 2143505408

real	0m2.144s
user	0m2.144s
sys	0m0.000s
*/

