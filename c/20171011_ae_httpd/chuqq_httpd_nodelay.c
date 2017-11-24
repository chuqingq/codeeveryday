// gcc -o chuqq_httpd ae.c zmalloc.c chuqq_httpd.c -I. -pthread
#define _GNU_SOURCE
#include <stdio.h>
#include <stdlib.h>
#include <strings.h>
#include <assert.h>

#include <unistd.h>
#include <fcntl.h>
#include <sched.h>
#include <pthread.h>

#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <netinet/tcp.h>

#include "ae.h"

const char RESPONSE[] = "HTTP/1.1 200 OK\r\nServer: fasthttp\r\nContent-Type: text/plain; charset=utf-8\r\nContent-Length: 11\r\nConnection: keep-alive\r\n\r\nhello world";

typedef struct {
	int index;
	pthread_t pthread;
	aeEventLoop *loop;
	int listenfd;
} thread_t;

typedef struct {
	char buf[1024];
	thread_t* thread;
	int fd;
	int cur;
	int write_cur;
} connection_t;

static void socket_writeable(aeEventLoop *loop, int fd, void *data, int mask);

static void socket_readable(aeEventLoop *loop, int fd, void *data, int mask) {
	connection_t *conn = data;
	int n = read(fd, conn->buf+conn->cur, sizeof(conn->buf)-conn->cur);

	if (n > 0) {
		conn->cur += n;

		// 判断请求接收完整
		if (conn->cur >= 4
			&& conn->buf[conn->cur-1] == '\n'
			&& conn->buf[conn->cur-2] == '\r'
			&& conn->buf[conn->cur-3] == '\n'
			&& conn->buf[conn->cur-4] == '\r') {
			aeDeleteFileEvent(conn->thread->loop, fd, AE_READABLE);
			// printf("socket_readable: aeDeleteFileEvent\n");
			conn->cur = 0;
			aeCreateFileEvent(conn->thread->loop, fd, AE_WRITABLE, socket_writeable, conn);
			// printf("socket_readable: aeCreateFileEvent\n");
			conn->write_cur = 0;
			// 默认1024的buf不会满
		}
	} else {
		aeDeleteFileEvent(loop, fd, AE_READABLE);
		close(fd);
		free(conn);
		if (n != 0 && conn->cur != 0) {
			printf("read error: %d. socket will close\n", n);
		}
	}
}

static void socket_writeable(aeEventLoop *loop, int fd, void *data, int mask) {
	connection_t *conn = data;

	int n = write(conn->fd, RESPONSE+conn->write_cur, sizeof(RESPONSE)-1-conn->write_cur);

	if (n != -1) {
		conn->write_cur += n;
		if (conn->write_cur == sizeof(RESPONSE)-1) {
			aeDeleteFileEvent(conn->thread->loop, fd, AE_WRITABLE);
			// printf("socket_writeable: aeDeleteFileEvent\n");
			aeCreateFileEvent(conn->thread->loop, fd, AE_READABLE, socket_readable, conn);
			// printf("socket_writeable: aeCreateFileEvent\n");
		} else {
			printf("write error: %d\n", n);
			aeDeleteFileEvent(conn->thread->loop, fd, AE_WRITABLE);
			close(fd);
			free(conn);
		}
	}
}

static void socket_accept(aeEventLoop *loop, int listenfd, void *data, int mask) {
	thread_t *thread = data;

	struct sockaddr_in addr;
	socklen_t addrlen = sizeof(struct sockaddr);
	int fd = accept(thread->listenfd, (struct sockaddr*)&addr, &addrlen);
	if (fd < 0) {
		perror("accept erorr:");
		return;
	}

	int flags = fcntl(fd, F_GETFL, 0);
	int ret = fcntl(fd, F_SETFL, flags | O_NONBLOCK);
	assert(ret == 0 && "setnonblock(fd) success");

	// setnodelay
	int on = 1;
	ret = setsockopt(fd, IPPROTO_TCP, TCP_NODELAY, (void *)&on, sizeof(on));
	assert(ret == 0 && "setsockopt(fd, TCP_NODELAY) success");

	connection_t *conn = calloc(sizeof(*conn), 1);
	conn->thread = thread;
	conn->fd = fd;
	aeCreateFileEvent(thread->loop, fd, AE_READABLE, socket_readable, conn);
}

void *thread_main(void *arg) {
	thread_t *thread = arg;
	printf("thread(%d) started\n", thread->index);

	cpu_set_t cpuset;
	CPU_ZERO(&cpuset);
	CPU_SET(thread->index, &cpuset);
	int ret = pthread_setaffinity_np(thread->pthread, sizeof(cpuset), &cpuset);
	assert(ret == 0 && "pthread_setaffinity_np success");

	thread->loop = aeCreateEventLoop(1024);

	thread->listenfd = socket(AF_INET, SOCK_STREAM, 0);
	assert(thread->listenfd >= 0);

	struct sockaddr_in addr;
	bzero(&addr, sizeof(addr));
	addr.sin_family = AF_INET;
	addr.sin_addr.s_addr = inet_addr("0.0.0.0");
	addr.sin_port = htons(8081);

	int reuse = 1;
	ret = setsockopt(thread->listenfd, SOL_SOCKET, SO_REUSEADDR, &reuse, sizeof(reuse));
	assert(ret == 0);

	int reuseport = 1;
	ret = setsockopt(thread->listenfd, SOL_SOCKET, SO_REUSEPORT, &reuseport, sizeof(reuseport));
	assert(ret == 0);

	ret = bind(thread->listenfd, (struct sockaddr*)&addr, sizeof(struct sockaddr));
	assert(ret == 0 && "bind success");

	ret = listen(thread->listenfd, 1024);
	assert(ret == 0 && "listen success");

	int flags = fcntl(thread->listenfd, F_GETFL, 0);
	ret = fcntl(thread->listenfd, F_SETFL, flags | O_NONBLOCK);
	assert(ret == 0 && "setnonblock(listenfd) success");

	ret = aeCreateFileEvent(thread->loop, thread->listenfd, AE_READABLE, socket_accept, thread);
	assert(ret == AE_OK && "aeCreateFileEvent(listenfd) success");

	aeMain(thread->loop);

	aeDeleteEventLoop(thread->loop);
}

thread_t threads[4];

int main() {
	int thread_num = sizeof(threads)/sizeof(threads[0]);

	for (int i = 0; i < thread_num; ++i) {
		threads[i].index = i;
		pthread_create(&threads[i].pthread, NULL, &thread_main, &threads[i]);
	}

	for (int i = 0; i < thread_num; ++i) {
		pthread_join(threads[i].pthread, NULL);
	}

	return 0;
}

/*
加nodelay前：
$ time wrk -c 100 -d 20 -t 4 http://127.0.0.1:8081/hello
Running 20s test @ http://127.0.0.1:8081/hello
  4 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.04ms    2.41ms  28.43ms   91.11%
    Req/Sec   139.70k     7.28k  209.02k    87.53%
  11147732 requests in 20.10s, 1.38GB read
Requests/sec: 554602.58
Transfer/sec:     70.35MB

real	0m20.136s
user	0m17.000s
sys	1m1.296s

加nodelay后：
$ time wrk -c 100 -d 20 -t 4 http://127.0.0.1:8081/hello
Running 20s test @ http://127.0.0.1:8081/hello
  4 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   736.96us    1.18ms  16.53ms   82.32%
    Req/Sec   140.74k     5.70k  165.73k    95.64%
  11231796 requests in 20.10s, 1.39GB read
Requests/sec: 558785.90
Transfer/sec:     70.88MB

real	0m20.135s
user	0m17.136s
sys	1m1.020s




ae.c的aeProcessEvents中去掉对timer的处理。
gcc -o chuqq_httpd -O3 ae.c zmalloc.c chuqq_httpd_nodelay.c -g -I. -pthread
sudo perf record -e cache-misses,branch-misses -g ./chuqq_httpd
sudo perf report
发现cache-misses和branch-misses都没有chuqq_httpd内部的了。应该在用户态epoll的方式已经基本到顶了，56w TPS。
*/
