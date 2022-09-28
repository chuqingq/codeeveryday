#include <unistd.h>
#include <stdio.h>
#include <stdlib.h>
#include <chrono>
#include <iostream>
#include "co_routine.h"
using namespace std;
/*
struct stTask_t
{
	int id;
};
*/
struct stEnv_t
{
	stCoEpoll_t* ctx;

	stCoCond_t* cond_p;
	stCoCond_t* cond_c;

	int data_p;
	int data_c;
	// queue<stTask_t*> task_queue;
};

const uint32_t attempts = 1000000;

void* Producer(void* args)
{
	co_enable_hook_sys();
	stEnv_t* env=  (stEnv_t*)args;
	int id = 0;
	for (auto i = 0; i < attempts; i++)
	// while (true)
	{
		// printf("produce %d\n", i);
		env->data_p = i;
		co_cond_signal(env->cond_c);

		co_cond_timedwait(env->cond_p, -1);
		if (env->data_c == i) {
		} else {
			printf("unexpected: %d, %d\n", env->data_c, i);
		}
		// poll(NULL, 0, 1000);
	}
	printf("producer exit\n");
	return NULL;
}
void* Consumer(void* args)
{
	co_enable_hook_sys();
	stEnv_t* env = (stEnv_t*)args;
	for (auto i = 0; i < attempts; i++)
	// while (true)
	{
		//printf("consume %d\n", i);
		co_cond_timedwait(env->cond_c, -1);
		if (env->data_p == i) {
		} else {
			printf("2unexpected: %d, %d\n", env->data_p, i);
		}

		env->data_c = i;
		co_cond_signal(env->cond_p);
	}
	printf("consumer exit\n");
	// exit(1);
	return NULL;
}

int wait_end(void* data) {
	stEnv_t* env = (stEnv_t *)data;
	if (env->data_p == (attempts-1) && env->data_c == (attempts-1)) {
		return -1;
	}
	// stCoEpoll_t* ctx = ((stEnv_t*)env)->ctx;
	// if (!ctx->pstActiveList->head && !ctx->pstTimeoutList->head) {
	// 	return -2;
	// } else {
	// 	return 0;
	// }
	return 0;
}

int main()
{
	stEnv_t* env = new stEnv_t;
	env->cond_p = co_cond_alloc();
	env->cond_c = co_cond_alloc();
auto startTime = std::chrono::steady_clock::now();

	stCoRoutine_t* consumer_routine;
	co_create(&consumer_routine, NULL, Consumer, env);
	co_resume(consumer_routine);

	stCoRoutine_t* producer_routine;
	co_create(&producer_routine, NULL, Producer, env);
	co_resume(producer_routine);
	
	// co_eventloop(co_get_epoll_ct(), NULL, NULL);
	env->ctx = co_get_epoll_ct();
	co_eventloop(env->ctx, wait_end, env);

  auto endTime = std::chrono::steady_clock::now();
	  auto duration = std::chrono::duration_cast<std::chrono::milliseconds>(
      endTime - startTime);

  auto perDuration =
      std::chrono::duration_cast<std::chrono::nanoseconds>(endTime - startTime)
          .count() /
      attempts;

  std::cout << "Using condtion_variable pingpong test times:" << attempts
            << " duration: " << duration.count() << "ms " << perDuration
            << "ns/op" << std::endl;

	return 0;
}
