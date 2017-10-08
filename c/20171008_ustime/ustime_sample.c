#include <stdio.h>
#include <sys/time.h>
#include <unistd.h>

static long long ustime(void) {
    struct timeval tv;
    // long long ust;

    gettimeofday(&tv, NULL);
    // ust = ((long)tv.tv_sec)*1000000;
    // ust += tv.tv_usec;
    // return ust;
    return ((long)tv.tv_sec)*1000000 + tv.tv_usec;
}

int main() {
    long long t1, t2, t3;
    t1 = ustime();
    t3 = ustime();
    sleep(1);
    t2 = ustime();

    printf("time diff: %lld\n", t2-t1);
    printf("t1: %lld, t2: %lld, t3: %lld\n", t1, t2, t3);
    return 0;
}

/*
~/work/codeeveryday/c/20171008_ustime $ gcc ustime_sample.c
~/work/codeeveryday/c/20171008_ustime $ ./a.out
time diff: 1002938
t1: 1507427133844181, t2: 1507427134847119, t3: 1507427133844181
*/

