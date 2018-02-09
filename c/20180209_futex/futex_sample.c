// gcc futex_sample.c -pthread
#include <stdio.h>
#include <linux/futex.h>
#include <sys/syscall.h>
#include <time.h>
#include <sys/time.h>
#include <unistd.h>
#include <pthread.h>

int futex(int *uaddr, int op, int val,
          const struct timespec *timeout,
          int *uaddr2, int val3) {
              return syscall(__NR_futex, uaddr, op, val, timeout, uaddr2, val3);
}

int fu = 0;

static unsigned long long nstime(void) {
    struct timespec ts;
    clock_gettime(CLOCK_REALTIME_COARSE, &ts);
    return ((unsigned long long)ts.tv_sec)*1e9 + ts.tv_nsec;
}

static void *worker_loop(void *arg) {
    struct timespec ts;
    ts.tv_sec = 3;
    ts.tv_nsec = 0;
    futex(&fu, FUTEX_WAIT, 0, &ts, NULL, 0);
    printf("worker time: %llu\n", nstime());
    return NULL;
}

int main() {
    fu = 0;
    pthread_t worker_tid;
    if (pthread_create(&worker_tid, NULL, worker_loop, NULL) != 0) {
        perror("pthread_create error");
        return -1;
    }
    sleep(2);
    printf("main time: %llu\n", nstime());
    futex(&fu, FUTEX_WAKE, 1, NULL, NULL, 0);
    sleep(2);
    return 0;
}
// chuqq@chuqq-hp:~/temp/codeeveryday/c/20180209_futex$ ./a.out 
// main time: 1518142828234522880
// worker time: 1518142828234522880
// 结论：这个速度比pthread_cond_signal快很多

