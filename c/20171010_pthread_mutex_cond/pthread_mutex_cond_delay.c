// gcc pthread_cond_delay.c -O0 -pthread
// strace ./a.out

#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <pthread.h>
#include <sys/time.h>

#include <time.h>
#include <unistd.h>

static unsigned long long nstime(void) {
    struct timespec ts;
    clock_gettime(CLOCK_REALTIME, &ts);
    return ((unsigned long long)ts.tv_sec)*1000000000 + ts.tv_nsec;
}


static long long ustime(void) {
    struct timeval tv;
    gettimeofday(&tv, NULL);
    return ((long)tv.tv_sec)*1000000 + tv.tv_usec;
}

pthread_mutex_t mutex = PTHREAD_MUTEX_INITIALIZER;
pthread_cond_t cond = PTHREAD_COND_INITIALIZER;

long long int starttime;

void* testThreadPool(int *t) {
	// printf("thread start: %d\n", *t);
	for (;;) {
		// pthread_mutex_lock(&mutex);
		pthread_cond_wait(&cond, &mutex);
		// printf("pthread_cond_wait time: %llu\n", nstime());
		printf("diff ns: %lld\n", nstime()-starttime);
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

	sleep(20);

	for (t = 0; t < thread_num; t++) {
		// printf("pthread_cond_signal time: %llu\n", nstime());
		starttime = nstime();
		// a = t;
		pthread_cond_signal(&cond);
		// pthread_mutex_unlock(&mutex);
		// sleep(2);
	}

	sleep(2);
}
// 验证pthread_cond_signal唤醒线程需要多久：大约40~50us
// chuqq-mbpr:
// > ./a.out
// diff ns: 45000
// virtualbox:
// > ./a.out
// diff ns: 141776
