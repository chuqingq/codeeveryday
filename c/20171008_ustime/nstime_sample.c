#include <stdio.h>
#include <time.h>
#include <unistd.h>

static unsigned long long nstime(void) {
    struct timespec ts;
    clock_gettime(CLOCK_REALTIME, &ts);
    return ((unsigned long long)ts.tv_sec)*1e9 + ts.tv_nsec;
}

int main() {
    unsigned long long t1, t2, t3;
    t1 = nstime();
    t3 = nstime();
    sleep(1);
    t2 = nstime();

    printf("time diff1: %llu\n", t3-t1);
    printf("time diff: %llu\n", t2-t1);
    printf("t1: %llu\nt3: %llu\nt2: %llu\n", t1, t3, t2);
    return 0;
}

/*
连续两次获取nanoseconds，大约花费250ns左右，无论是否开启-O3。

chuqq@chuqq-hp:~/temp/codeeveryday/c/20171008_ustime$ ./a.out 
time diff1: 256
time diff: 1000087552
t1: 1508460772922802688
t3: 1508460772922802944
t2: 1508460773922890240
chuqq@chuqq-hp:~/temp/codeeveryday/c/20171008_ustime$ ./a.out 
time diff1: 256
time diff: 1000046848
t1: 1508460779526198272
t3: 1508460779526198528
t2: 1508460780526245120


# test on MPBr:

~/work/codeeveryday/c/20171008_ustime $ gcc nstime_sample.c
~/work/codeeveryday/c/20171008_ustime $ ./a.out
time diff: 1005163000
t1: 1507427455285188000, t2: 1507427456290351000, t3: 1507427455285188000

# test on ubuntu:

chuqq@chuqq-VPCS:~/work/codeeveryday/c/20171008_ustime$ gcc -O3 nstime_sample.c 
chuqq@chuqq-VPCS:~/work/codeeveryday/c/20171008_ustime$ ./a.out 
time diff: 1000167424
t1: 1507428223080044288
t3: 1507428223080044544
t2: 1507428224080211712
*/

