#include <stdio.h>
// #include <sys/time.h>
#include <time.h>
#include <unistd.h>

static unsigned long long nstime(void) {
    struct timespec ts;
    // long long ust;

    clock_gettime(CLOCK_REALTIME, &ts);
    // ust = ((long)tv.tv_sec)*1000000;
    // ust += tv.tv_usec;
    // return ust;
    return ((unsigned long long)ts.tv_sec)*1000000000 + ts.tv_nsec;
}

int main() {
    unsigned long long t1, t2, t3;
    t1 = nstime();
    t3 = nstime();
    sleep(1);
    t2 = nstime();

    printf("time diff: %llu\n", t2-t1);
    printf("t1: %llu, t2: %llu, t3: %llu\n", t1, t2, t3);
    return 0;
}

/*
~/work/codeeveryday/c/20171008_ustime $ gcc nstime_sample.c
~/work/codeeveryday/c/20171008_ustime $ ./a.out
time diff: 1005163000
t1: 1507427455285188000, t2: 1507427456290351000, t3: 1507427455285188000
*/

