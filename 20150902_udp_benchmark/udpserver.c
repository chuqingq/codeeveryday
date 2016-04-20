// gcc -o udpserver udpserver.c -I /home/chuqq/bin/libuv/include/ -L /home/chuqq/bin/libuv/lib/ -luv
// LD_LIBRARY_PATH=/home/chuqq/bin/libuv/lib ./udpserver

#define _GNU_SOURCE
#include <sched.h>
#include <pthread.h>

#include "uv.h"

// #include <assert.h>
#include <stdio.h>
#include <stdlib.h>


long recv_count = 0;// TODO atomic
long send_count = 0;

#define CPU_NUM 24

uv_loop_t* loop[CPU_NUM];
uv_udp_t udp_handle[CPU_NUM];
#define MSG_SIZE 32
char data[MSG_SIZE];
// uv_buf_t buf;
uv_udp_send_t send_req[CPU_NUM];

uv_timer_t timer_handle;

void timer_cb(uv_timer_t* timer) {
	printf("recv: %ld, send: %ld\n", recv_count, send_count);
	recv_count = send_count = 0;
}

void alloc_cb(uv_handle_t* handle, size_t suggested_size, uv_buf_t* buf) {
	buf->base = data;
	buf->len = MSG_SIZE;
	return;
}

void send_cb(uv_udp_send_t* req, int status);

void recv_cb(uv_udp_t* handle, ssize_t nread, const uv_buf_t* buf, const struct sockaddr* addr, unsigned flags) {
	if (nread < 0) { printf("recv_cb error\n"); return; }
	if (nread == 0) return;
	// count++;
	__sync_fetch_and_add(&recv_count, 1);
	// uv_udp_recv_stop(handle);
	uv_udp_send_t* send_req = malloc(sizeof(*send_req));
	uv_udp_send(send_req, handle, buf, 1, addr, send_cb);
	// uv_udp_send(&send_req[((char*)handle-(char*)udp_handle)/sizeof(uv_udp_t)], handle, buf, 1, addr, send_cb);
}

void send_cb(uv_udp_send_t* req, int status) {
	if (status != 0) { printf("send_cb error\n"); return; }
	// uv_udp_recv_start(&udp_handle[((char*)req-(char*)send_req)/sizeof(uv_udp_send_t)], alloc_cb, recv_cb);
	// count++;
	__sync_fetch_and_add(&send_count, 1);
	free(req);
}

void thread_cb(void* arg) {
	long i = (long)arg;

	// cpu_set_t mask;
	// CPU_ZERO(&mask);
	// CPU_SET(i, &mask);
	// if (pthread_setaffinity_np(pthread_self(), sizeof(mask), &mask) < 0) {
	// 	perror("pthread_setaffinity_np error");
	// }

	loop[i] = malloc(sizeof(*loop[i]));
	uv_loop_init(loop[i]);

	uv_udp_init(loop[i], &udp_handle[i]);
	struct sockaddr_in myaddr;
	uv_ip4_addr("0.0.0.0", (int)20000 + i, &myaddr);
	uv_udp_bind(&udp_handle[i], (const struct sockaddr *)&myaddr, UV_UDP_REUSEADDR);
	uv_udp_recv_start(&udp_handle[i], alloc_cb, recv_cb);

	uv_run(loop[i], UV_RUN_DEFAULT);
}

int main() {
	uv_thread_t thread[CPU_NUM];
	int i = 0;
	for (i = 0; i < CPU_NUM; ++i)
	{
		uv_thread_create(&thread[i], thread_cb, (void*)i);
	}

	uv_timer_init(uv_default_loop(), &timer_handle);
	uv_timer_start(&timer_handle, timer_cb, 0, 1000);
	uv_run(uv_default_loop(), UV_RUN_DEFAULT);
}
