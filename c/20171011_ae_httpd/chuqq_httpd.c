// gcc -o chuqq_httpd ae.c zmalloc.c chuqq_httpd.c -I. -pthread
#include <stdio.h>
#include <stdlib.h>
#include <strings.h>
#include <assert.h>

#include <unistd.h>
#include <fcntl.h>
#include <pthread.h>

#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>

#include "ae.h"

const char RESPONSE[] = "HTTP/1.1 200 OK\r\nServer: fasthttp\r\nContent-Type: text/plain; charset=utf-8\r\nContent-Length: 11\r\nConnection: keep-alive\r\n\r\nhello world";

typedef struct {
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

	connection_t *conn = calloc(sizeof(*conn), 1);
	conn->thread = thread;
	conn->fd = fd;
	aeCreateFileEvent(thread->loop, fd, AE_READABLE, socket_readable, conn);
}

void *thread_main(void *arg) {
	thread_t *thread = arg;

	thread->loop = aeCreateEventLoop(1024);

	thread->listenfd = socket(AF_INET, SOCK_STREAM, 0);
	assert(thread->listenfd >= 0);

	struct sockaddr_in addr;
	bzero(&addr, sizeof(addr));
	addr.sin_family = AF_INET;
	addr.sin_addr.s_addr = inet_addr("0.0.0.0");
	addr.sin_port = htons(8081);

	int reuse = 1;
	int ret = setsockopt(thread->listenfd, SOL_SOCKET, SO_REUSEADDR, &reuse, sizeof(reuse));
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

thread_t threads[8];

int main() {
	int thread_num = sizeof(threads)/sizeof(threads[0]);

	for (int i = 0; i < thread_num; ++i) {
		pthread_create(&threads[i].pthread, NULL, &thread_main, &threads[i]);
	}

	for (int i = 0; i < thread_num; ++i) {
		pthread_join(threads[i].pthread, NULL);
	}

	return 0;
}

