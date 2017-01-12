#include <stdio.h>
#include <assert.h>
#include <sys/time.h>

#include <uv.h>

uv_loop_t* loop;

static void timer_cb(uv_timer_t* timer, int status) {
	// struct timeval tv;
	// gettimeofday(&tv, NULL);
	// printf("timer_cb: %ld, %ld\n", tv.tv_sec, tv.tv_usec);
	printf("%lld\n", uv_now(loop));
}

int main(int argc, char const *argv[])
{
	uv_timer_t timer;
	int r;

	loop = uv_default_loop();
	r = uv_timer_init(loop, &timer);
	assert(r == 0);

	r = uv_timer_start(&timer, timer_cb, 1000, 1000);
	assert(r == 0);

	uv_run(loop, UV_RUN_DEFAULT);
	return 0;
}
