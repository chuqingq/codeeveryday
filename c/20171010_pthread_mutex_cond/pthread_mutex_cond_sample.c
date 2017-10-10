// gcc pthread_mutex_cond_sample.c -O0
// strace ./a.out

#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <pthread.h>

pthread_mutex_t mutex = PTHREAD_MUTEX_INITIALIZER;
//pthread_cond_t cond = PTHREAD_COND_INITIALIZER;

int a = 0;

void* testThreadPool(int *t) {
	printf("thread start: %d\n", *t);
	for (;;) {
		pthread_mutex_lock(&mutex);
		// pthread_cond_wait(&cond, &mutex);
		// printf("thread: %d, a: %d\n", *t, a);
		pthread_mutex_unlock(&mutex);
		// sleep(3);
	}
	return (void*) 0;
}

int main() {
	int thread_num = 5;
	pthread_t *mythread = (pthread_t*) malloc(thread_num* sizeof(*mythread));

	int t;
	for (t = 0; t < 5; t++) {
		int *i=(int*)malloc(sizeof(int));
		*i=t;
		if (pthread_create(&mythread[t], NULL, (void*)testThreadPool, (void*)i) != 0) {
			printf("pthread_create");
		}
	}

	for (t = 0; t < 100; t++) {
		printf("main: a: %d\n", t);
		pthread_mutex_lock(&mutex);
		a = t;
		// pthread_cond_signal(&cond);
		pthread_mutex_unlock(&mutex);
		sleep(2);
	}
}

/*
总结futex的逻辑：
1、wait：先在用户态对原子变量做操作：如果没竞争发生，则成功；如果有竞争发生，则调用futex_wait休眠当前线程。
2、wake：先在用户态对原子变量做操作：如果没有其他线程在wait，则直接成功；如果有其他线程在wait，则调用futex_wake唤醒一个或多个线程。


如果不用cond，只用mutex，没有子线程，貌似没有系统调用：
rt_sigprocmask(SIG_UNBLOCK, [RTMIN RT_1], NULL, 8) = 0
getrlimit(RLIMIT_STACK, {rlim_cur=8192*1024, rlim_max=RLIM64_INFINITY}) = 0
brk(NULL)                               = 0x2438000
brk(0x2459000)                          = 0x2459000
fstat(1, {st_mode=S_IFCHR|0620, st_rdev=makedev(136, 1), ...}) = 0
write(1, "main: a: 0\n", 11main: a: 0
)            = 11
nanosleep({2, 0}, 0x7fffceb3cce0)       = 0
write(1, "main: a: 1\n", 11main: a: 1
)            = 11
nanosleep({2, 0}, 0x7fffceb3cce0)       = 0
write(1, "main: a: 2\n", 11main: a: 2
)            = 11
nanosleep({2, 0}, 0x7fffceb3cce0)       = 0
write(1, "main: a: 3\n", 11main: a: 3
)            = 11
nanosleep({2, 0}, 0x7fffceb3cce0)       = 0
write(1, "main: a: 4\n", 11main: a: 4
)            = 11
nanosleep({2, 0}, ^Cstrace: Process 5137 detached
 <detached ...>

如果不用cond，只用mutex，且子线程中不sleep，有了futex的WAIT和WAKE系统调用，且WAIT会失败EAGAIN：
child_stack=0x7f3795b56ff0, flags=CLONE_VM|CLONE_FS|CLONE_FILES|CLONE_SIGHAND|CLONE_THREAD|CLONE_SYSVSEM|CLONE_SETTLS|CLONE_PARENT_SETTID|CLONE_CHILD_CLEARTID, parent_tidptr=0x7f3795b579d0, tls=0x7f3795b57700, child_tidptr=0x7f3795b579d0) = 5163
write(1, "main: a: 0\n", 11main: a: 0
)            = 11
nanosleep({2, 0}, 0x7ffdc6e1a630)       = 0
write(1, "main: a: 1\n", 11main: a: 1
)            = 11
futex(0x601080, FUTEX_WAKE_PRIVATE, 1)  = 0
nanosleep({2, 0}, 0x7ffdc6e1a630)       = 0
write(1, "main: a: 2\n", 11main: a: 2
)            = 11
nanosleep({2, 0}, 0x7ffdc6e1a630)       = 0
write(1, "main: a: 3\n", 11main: a: 3
)            = 11
futex(0x601080, FUTEX_WAIT_PRIVATE, 2, NULL) = -1 EAGAIN (Resource temporarily unavailable)
futex(0x601080, FUTEX_WAIT_PRIVATE, 2, NULL) = -1 EAGAIN (Resource temporarily unavailable)
futex(0x601080, FUTEX_WAKE_PRIVATE, 1)  = 0
nanosleep({2, 0}, 0x7ffdc6e1a630)       = 0
write(1, "main: a: 4\n", 11main: a: 4
)            = 11
nanosleep({2, 0}, 0x7ffdc6e1a630)       = 0
write(1, "main: a: 5\n", 11main: a: 5
)            = 11
futex(0x601080, FUTEX_WAKE_PRIVATE, 1)  = 0
nanosleep({2, 0}, 0x7ffdc6e1a630)       = 0
write(1, "main: a: 6\n", 11main: a: 6
)            = 11
futex(0x601080, FUTEX_WAIT_PRIVATE, 2, NULL) = 0
futex(0x601080, FUTEX_WAKE_PRIVATE, 1)  = 0
nanosleep({2, 0}, 0x7ffdc6e1a630)       = 0
write(1, "main: a: 7\n", 11main: a: 7
)            = 11
futex(0x601080, FUTEX_WAIT_PRIVATE, 2, NULL) = -1 EAGAIN (Resource temporarily unavailable)
futex(0x601080, FUTEX_WAKE_PRIVATE, 1)  = 0
nanosleep({2, 0}, ^Cstrace: Process 5158 detached
 <detached ...>


如果mutex+cond，且子线程没有sleep，不会有WAIT失败：
mmap(NULL, 8392704, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS|MAP_STACK, -1, 0) = 0x7f9509ffc000
mprotect(0x7f9509ffc000, 4096, PROT_NONE) = 0
clone(thread start: 4
child_stack=0x7f950a7fbff0, flags=CLONE_VM|CLONE_FS|CLONE_FILES|CLONE_SIGHAND|CLONE_THREAD|CLONE_SYSVSEM|CLONE_SETTLS|CLONE_PARENT_SETTID|CLONE_CHILD_CLEARTID, parent_tidptr=0x7f950a7fc9d0, tls=0x7f950a7fc700, child_tidptr=0x7f950a7fc9d0) = 5196
write(1, "main: a: 0\n", 11main: a: 0
)            = 11
futex(0x6010e4, FUTEX_WAKE_OP_PRIVATE, 1, 1, 0x6010e0, {FUTEX_OP_SET, 0, FUTEX_OP_CMP_GT, 1}) = 1
futex(0x6010a0, FUTEX_WAKE_PRIVATE, 1)  = 1
nanosleep({2, 0}, 0x7ffd8a3e9960)       = 0
write(1, "main: a: 1\n", 11main: a: 1
)            = 11
futex(0x6010e4, FUTEX_WAKE_OP_PRIVATE, 1, 1, 0x6010e0, {FUTEX_OP_SET, 0, FUTEX_OP_CMP_GT, 1}) = 1
futex(0x6010a0, FUTEX_WAKE_PRIVATE, 1)  = 1
nanosleep({2, 0}, 0x7ffd8a3e9960)       = 0
write(1, "main: a: 2\n", 11main: a: 2
)            = 11
futex(0x6010e4, FUTEX_WAKE_OP_PRIVATE, 1, 1, 0x6010e0, {FUTEX_OP_SET, 0, FUTEX_OP_CMP_GT, 1}) = 1
futex(0x6010a0, FUTEX_WAKE_PRIVATE, 1)  = 1
nanosleep({2, 0}, 0x7ffd8a3e9960)       = 0
write(1, "main: a: 3\n", 11main: a: 3
)            = 11
futex(0x6010e4, FUTEX_WAKE_OP_PRIVATE, 1, 1, 0x6010e0, {FUTEX_OP_SET, 0, FUTEX_OP_CMP_GT, 1}) = 1
futex(0x6010a0, FUTEX_WAKE_PRIVATE, 1)  = 1
nanosleep({2, 0}, 0x7ffd8a3e9960)       = 0
write(1, "main: a: 4\n", 11main: a: 4
)            = 11
futex(0x6010e4, FUTEX_WAKE_OP_PRIVATE, 1, 1, 0x6010e0, {FUTEX_OP_SET, 0, FUTEX_OP_CMP_GT, 1}) = 1
futex(0x6010a0, FUTEX_WAKE_PRIVATE, 1)  = 1
nanosleep({2, 0}, ^Cstrace: Process 5190 detached
 <detached ...>
*/

