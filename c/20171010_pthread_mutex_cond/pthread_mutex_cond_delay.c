// gcc pthread_mutex_cond_delay.c -O3 -pthread
// ./a.out

#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <pthread.h>

#include <time.h>
#include <unistd.h>

static unsigned long long nstime(void) {
    struct timespec ts;
    clock_gettime(CLOCK_REALTIME, &ts);
    return ((unsigned long long)ts.tv_sec)*1000000000 + ts.tv_nsec;
}

pthread_mutex_t mutex = PTHREAD_MUTEX_INITIALIZER;
pthread_cond_t cond = PTHREAD_COND_INITIALIZER;

long long int starttime;

void* testThreadPool(int *t) {
	// printf("thread start: %d\n", *t);
	// for (;;) {
		pthread_mutex_lock(&mutex);
		// printf("worker mutex\n");
		pthread_cond_wait(&cond, &mutex);
		// printf("pthread_cond_wait time: %llu\n", nstime());
		printf("diff ns: %lld\n", nstime()-starttime);
		pthread_mutex_unlock(&mutex);
		// sleep(3);
	// }
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
	// printf("create thread\n");
	sleep(1);

	for (t = 0; t < thread_num; t++) {
		// printf("pthread_cond_signal time: %llu\n", nstime());
		starttime = nstime();
                pthread_mutex_lock(&mutex);
		// a = t;
		pthread_cond_signal(&cond);
		pthread_mutex_unlock(&mutex);
		// printf("after unlock\n");
		// sleep(2);
	}

	// sleep(2);
        for (t = 0; t < thread_num; t++) {
            pthread_join(mythread[t], NULL);
        }
}
// 验证pthread_cond_signal唤醒线程需要多久：大约40~50us
// chuqq-mbpr:
// > ./a.out
// diff ns: 45000
// virtualbox:
// > ./a.out
// diff ns: 141776


/*
chuqq@chuqq:~/work/codeeveryday/c/20171010_pthread_mutex_cond$ gcc pthread_mutex_cond_delay.c -O3 -pthread
chuqq@chuqq:~/work/codeeveryday/c/20171010_pthread_mutex_cond$ ./a.out
diff ns: 105583
chuqq@chuqq:~/work/codeeveryday/c/20171010_pthread_mutex_cond$ ./a.out
diff ns: 21451
chuqq@chuqq:~/work/codeeveryday/c/20171010_pthread_mutex_cond$ ./a.out
diff ns: 116579
chuqq@chuqq:~/work/codeeveryday/c/20171010_pthread_mutex_cond$ ./a.out
diff ns: 23621
*/
