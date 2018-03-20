// gcc -O3 futex_sample.c -pthread
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

long long int starttime;

static unsigned long long nstime(void) {
    struct timespec ts;
    clock_gettime(CLOCK_REALTIME, &ts);
    return ((unsigned long long)ts.tv_sec)*1e9 + ts.tv_nsec;
}

static void *worker_loop(void *arg) {
    struct timespec ts;
    ts.tv_sec = 3;
    ts.tv_nsec = 0;
    futex(&fu, FUTEX_WAIT, 0, &ts, NULL, 0);
    // printf("worker time: %llu\n", nstime());
    printf("diff ns: %lld\n", nstime()-starttime);
    return NULL;
}

int main() {
    fu = 0;
    pthread_t worker_tid;
    if (pthread_create(&worker_tid, NULL, worker_loop, NULL) != 0) {
        perror("pthread_create error");
        return -1;
    }
    sleep(1);
    // printf("main time: %llu\n", nstime());
    starttime = nstime();
    futex(&fu, FUTEX_WAKE, 1, NULL, NULL, 0);
    // sleep(2);
    pthread_join(worker_tid, NULL);
    return 0;
}
// chuqq@chuqq-hp:~/temp/codeeveryday/c/20180209_futex$ ./a.out 
// main time: 1518142828234522880
// worker time: 1518142828234522880
// lubuntu@virtualbox
// > ./a.out
// diff ns: 0
// 结论：这个速度比pthread_cond_signal快很多

/*
chuqq@chuqq:~/work/codeeveryday/c/20180209_futex$ ./a.out
diff ns: 22016
chuqq@chuqq:~/work/codeeveryday/c/20180209_futex$ ./a.out
diff ns: 83456
chuqq@chuqq:~/work/codeeveryday/c/20180209_futex$ ./a.out
diff ns: 25088
chuqq@chuqq:~/work/codeeveryday/c/20180209_futex$ ./a.out
diff ns: 22272
chuqq@chuqq:~/work/codeeveryday/c/20180209_futex$ ./a.out
diff ns: 20224
chuqq@chuqq:~/work/codeeveryday/c/20180209_futex$ ./a.out
diff ns: 147200
chuqq@chuqq:~/work/codeeveryday/c/20180209_futex$ ./a.out
diff ns: 14080
结论：大约20us，并不比pthread_cond_signal快很多。
前面是因为用了_COARSE，统计的时间不准

验证环境：chuqq@hp
chuqq@chuqq-hp:~/temp/codeeveryday/c/20180209_futex$ ./a.out 
diff ns: 36864
chuqq@chuqq-hp:~/temp/codeeveryday/c/20180209_futex$ ./a.out 
diff ns: 37120
chuqq@chuqq-hp:~/temp/codeeveryday/c/20180209_futex$ ./a.out 
diff ns: 37120
chuqq@chuqq-hp:~/temp/codeeveryday/c/20180209_futex$ ./a.out 
diff ns: 36352
chuqq@chuqq-hp:~/temp/codeeveryday/c/20180209_futex$ ./a.out 
diff ns: 36608
chuqq@chuqq-hp:~/temp/codeeveryday/c/20180209_futex$ ./a.out 
diff ns: 37376
*/
