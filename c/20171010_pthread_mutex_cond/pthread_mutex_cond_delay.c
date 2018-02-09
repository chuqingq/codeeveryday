// gcc pthread_mutex_cond_sample.c -O0 -pthread
// strace ./a.out

#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <pthread.h>

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

pthread_mutex_t mutex = PTHREAD_MUTEX_INITIALIZER;
pthread_cond_t cond = PTHREAD_COND_INITIALIZER;


// int a = 0;

void* testThreadPool(int *t) {
	printf("thread start: %d\n", *t);
	for (;;) {
		// pthread_mutex_lock(&mutex);
		pthread_cond_wait(&cond, &mutex);
		printf("pthread_cond_wait time: %llu\n", nstime());
		pthread_mutex_unlock(&mutex);
		// sleep(3);
	}
	return (void*) 0;
}

int main() {
	int thread_num = 1;
	pthread_t *mythread = (pthread_t*) malloc(thread_num* sizeof(*mythread));

	// pthread_mutex_lock(&mutex);

	int t;
	for (t = 0; t < thread_num; t++) {
		int *i=(int*)malloc(sizeof(int));
		*i=t;
		if (pthread_create(&mythread[t], NULL, (void*)testThreadPool, (void*)i) != 0) {
			printf("pthread_create");
		}
	}

	sleep(2);

	for (t = 0; t < thread_num; t++) {
		printf("pthread_cond_signal time: %llu\n", nstime());
		// a = t;
		pthread_cond_signal(&cond);
		// pthread_mutex_unlock(&mutex);
		// sleep(2);
	}

	sleep(2);
}
// 验证pthread_cond_signal唤醒线程需要多久
// chuqq@chuqq-hp:~/temp/codeeveryday/c/20171010_pthread_mutex_cond$ ./a.out 
// thread start: 0
// pthread_cond_signal time: 1518140277785214707
// pthread_cond_wait time: 1518140277785270798
// chuqq@chuqq-hp:~/temp/codeeveryday/c/20171010_pthread_mutex_cond$ ./a.out 
// thread start: 0
// pthread_cond_signal time: 1518140306485181135
// pthread_cond_wait time: 1518140306485249478
// 结论：大约50～60us
