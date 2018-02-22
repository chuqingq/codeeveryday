// gcc eventfd_sample.c
#include <sys/eventfd.h>
#include <stdio.h>
#include <pthread.h>
#include <unistd.h>
#include <sys/time.h>


long long int starttime;

static unsigned long long nstime(void) {
    struct timespec ts;
    clock_gettime(CLOCK_REALTIME_COARSE, &ts);
    return ((unsigned long long)ts.tv_sec)*1e9 + ts.tv_nsec;
}

int fd;
uint64_t buffer;

void *threadFunc(void *arg) //线程函数
{
    int t;
    while (1)
    {
        t = read(fd, &buffer, sizeof(buffer)); //阻塞等待fd可读，及通知事件发生
        printf("diff ns: %lld\n", nstime()-starttime);
        if (sizeof(buffer) < 8)
        {
            printf("buffer错误\n");
        }
        printf("111 t = %llu   buffer = %llu\n", t, buffer);
        if (t == 8)
        {
            printf("唤醒成功\n");
        }
        sleep(3);
    }
}

void *threadFunc2(void *arg) //线程函数
{
    int t;
    while (1)
    {
        t = read(fd, &buffer, sizeof(buffer)); //阻塞等待fd可读，及通知事件发生
        printf("diff ns: %lld\n", nstime()-starttime);
        if (sizeof(buffer) < 8)
        {
            printf("buffer错误\n");
        }
        printf("222 t = %llu   buffer = %llu\n", t, buffer);
        if (t == 8)
        {
            printf("唤醒成功\n");
        }

        sleep(3);
    }
}

int main(void)
{
    uint64_t buf = 1;
    int ret;
    pthread_t tid, tid2;

    if ((fd = eventfd(0, 0)) == -1) //创建事件驱动的文件描述符
    {
        printf("创建失败\n");
    }

    //创建线程
    if (pthread_create(&tid, NULL, threadFunc, NULL) < 0)
    {
        printf("线程创建失败\n");
    }

    if (pthread_create(&tid2, NULL, threadFunc2, NULL) < 0)
    {
        printf("线程创建失败\n");
    }

    while (1)
    {
        starttime = nstime();
        ret = write(fd, &buf, sizeof(buf)); // 向eventfd写，其实是在count上累加
        if (ret != 8)
        {
            printf("写错误： %d\n", ret);
            perror("write eventfd error");
        }
        sleep(1); // 每秒通知一次
    }

    return 0;
}

// 不一定哪个先启动
// 222 t = 8   buffer = 1
// 唤醒成功
// 111 t = 8   buffer = 1
// 唤醒成功
// 222 t = 8   buffer = 1
// 唤醒成功
// 111 t = 8   buffer = 1
// 唤醒成功
// 222 t = 8   buffer = 1
// 唤醒成功
// 111 t = 8   buffer = 1
// 唤醒成功
// 222 t = 8   buffer = 1
// 唤醒成功
// 111 t = 8   buffer = 1
// 唤醒成功
// 222 t = 8   buffer = 1
// 唤醒成功
// 111 t = 8   buffer = 1
// 唤醒成功
// 222 t = 8   buffer = 1
// 唤醒成功
// 111 t = 8   buffer = 1
// 唤醒成功

// lubuntu@virtualbox 时延也可以忽略了：0
// > ./a.out
// diff ns: 0
// 222 t = 8   buffer = 1
// 唤醒成功
// diff ns: 0
// 111 t = 8   buffer = 1
// 唤醒成功
// diff ns: 0
// 222 t = 8   buffer = 2
// 唤醒成功
// diff ns: 0
// 111 t = 8   buffer = 1
// 唤醒成功
// diff ns: 1000029952
// 222 t = 8   buffer = 1
// 唤醒成功
// diff ns: 1004030208
// 111 t = 8   buffer = 1
// 唤醒成功
// diff ns: 1000030208
// 222 t = 8   buffer = 2
// 唤醒成功
