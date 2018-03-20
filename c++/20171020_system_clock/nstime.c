
#include <time.h>
#include <unistd.h>

static inline unsigned long long nstime(void) {
    struct timespec ts;
    clock_gettime(CLOCK_REALTIME, &ts);
    return ((unsigned long long)ts.tv_sec)*1e9 + ts.tv_nsec;
}

#include <stdio.h>

int main() {
    unsigned long long start = nstime();
    unsigned long long stop = nstime();
    printf("elapsed: %llu ns\n", stop-start);
    return 0;
}
