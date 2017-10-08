#include <stdio.h>
#include <stdlib.h>
// #include <stdint.h>
// #include <inttypes.h>
#include <unistd.h>

typedef unsigned long long ticks;

static __inline__ ticks getticks(void)
{
    u_int32_t a, d;

    asm volatile("rdtsc" : "=a" (a), "=d" (d));
    return (((ticks)a) | (((ticks)d) << 32));
}

int main() {
    ticks t1, t2;
    t1 = getticks();
    sleep(1);
    t2 = getticks();

    printf("time diff: %llu\n", t2 - t1);

    return 0;
}

/*
~/work/codeeveryday/c/20171008_rdtsc $ gcc rdtsc_sample.c
~/work/codeeveryday/c/20171008_rdtsc $ ./a.out
time diff: 2700288580
~/work/codeeveryday/c/20171008_rdtsc $ time ./a.out
time diff: 2713833132

real   	0m1.012s
user   	0m0.001s
sys    	0m0.001s
*/

