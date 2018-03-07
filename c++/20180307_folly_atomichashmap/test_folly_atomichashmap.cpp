// g++ -std=c++14 -Ifolly_bin/include -Lfolly_bin/lib test.cpp -lfolly -lglog -ldl -ldouble-conversion -pthread
// 静态链接 g++ -g -O3 -std=c++14 -Ifolly_bin/include test.cpp -Lfolly_bin/lib -lfolly -static-libgcc -static-libstdc++ -Wl,-Bstatic -lglog -lgflags -ldouble-conversion -Wl,-Bdynamic -lunwind -ldl -ldouble-conversion -pthread
// g++ -g -O3 -std=c++14 -Ifolly_bin/include test.cpp -Lfolly_bin/lib -lfolly  -static-libstdc++ -Wl,-Bstatic -ldouble-conversion -lglog -lgflags -lunwind -llzma -Wl,-Bdynamic -ldl  -pthread

#include <folly/AtomicHashMap.h>

#include <utility>

#include <stdio.h>
#include <stddef.h>
#include <stdlib.h>
#include <string.h>
#include <pthread.h>

// #include "atomic_hash.h"

#define COMMON_SUCCESS 1
#define COMMON_FAIL 0

#define CHANNEL_TOKEN_LEN 8
#define MAX_NODE 100000
#define TTL 660000

using folly::AtomicHashMap;

#include <time.h>
#include <unistd.h>

static unsigned long long nstime(void) {
    struct timespec ts;
    clock_gettime(CLOCK_REALTIME_COARSE, &ts);
    return ((unsigned long long)ts.tv_sec)*1e9 + ts.tv_nsec;
}

/**
 * 终端ip及port构造的hash key数据结构
 */
#pragma pack(2)
typedef struct
{
    uint32_t ip;   // IP地址
    uint16_t port; // 端口号
    uint16_t unused;
} device_key_t;
#pragma pack()

/**
 * 终端deviceid构造的hash value数据结构
 */
// #pragma pack(2)
// typedef struct
// {
//     uint32_t ip;                       // IP地址
//     uint16_t port;                     // 端口号
//     uint8_t id[8]; // 设备编号，8字节
// } deviceid_value_t;
// #pragma pack()

/**
 * 设备信息
 * 终端数据结构
 */
#pragma pack(2)
typedef struct
{
    uint32_t ip;   // IP地址.前两个字段跟device_key_t结构统一，不要乱动。
    uint16_t port; // 端口号
    uint16_t unused;

    uint8_t id[8]; // 设备编号，8字节
    // struct ether_addr mac;             // MAC地址，6字节

    uint32_t dev_seq_be;  // 网络序，设备最近一次tcp ack消息的序号
    uint32_t dev_ack_be;  // 网络序，设备最近一次tcp ack消息的确认号
    uint32_t sev_seq_cpu; // 主机序，服务端的发送序号，每次发送自增，如果设备连续发送两次相同的且小于该序号的ack，则将该序号重置为该ack
                          //    uint8_t status;       // 设备当前的状态，如：STATUS_DEVICE_GET_AESKEY， STATUS_DEVICE_GO_ONLINE，STATUS_DEVICE_ONLINE
    uint8_t type;         // 设备类型，TYPE_DEVICE_TCP_V1/TYPE_DEVICE_TCP_V3/TYPE_DEVICE_UDP_V3

    uint8_t aes_key[16]; // 终端AESKey，用作终端和Channel之间的安全传输加密. 16字节，兼容老版本TYPE_DEVICE_TCP_V1（8字节）
    time_t aes_key_time;                    // aeskey 生成时间
    time_t heart_beat_time;                 // 心跳时间戳，4字节
    uint8_t pub_key[294];
} device_t;
#pragma pack()

// static hash_t *dict_device;
// static hash_t *dict_device2;

void convert_int2string(const uint8_t *array, int len, char *str)
{
    int i;
    for (i = 0; i < len; i++)
    {
        snprintf(&(str[i * 2]), 3, "%02X" , *(array + i));
    }
}

// static int cb_add_token(void *data, void *arg)
// {
//     // char str_channel_token[2 * CHANNEL_TOKEN_LEN + 1] = {0};
//     // convert_int2string((uint8_t *)data, CHANNEL_TOKEN_LEN, str_channel_token);
//     // TLOGD("cb_add_token", "", "Device token [%s] added in dict_device.", str_channel_token);
//     return PLEASE_DO_NOT_CHANGE_TTL;
// }

// static int cb_dup_token(void *data, void *arg)
// {
//     // char str_channel_token[2 * CHANNEL_TOKEN_LEN + 1] = {0};
//     // convert_int2string((uint8_t *)data, CHANNEL_TOKEN_LEN, str_channel_token);
//     // TLOGW("cb_dup_token", "", "Device token [%s] duplicated in dict_device.", str_channel_token);
//     return PLEASE_SET_TTL_TO_DEFAULT;
// }

// static int cb_get_token(void *data, void *arg)
// {
//     *((void **)arg) = data;
//     // char str_channel_token[2 * CHANNEL_TOKEN_LEN + 1] = {0};
//     // convert_int2string((uint8_t *)data, CHANNEL_TOKEN_LEN, str_channel_token);
//     // TLOGD("cb_get_token", "", "Device token [%s] get in dict_device.", str_channel_token);
//     return PLEASE_SET_TTL_TO_DEFAULT;
// }

// static int cb_del_token(void *data, void *arg)
// {
//     // char str_channel_token[2 * CHANNEL_TOKEN_LEN + 1] = {0};
//     // convert_int2string((uint8_t *)data, CHANNEL_TOKEN_LEN, str_channel_token);
//     // TLOGD("cb_del_token", "", "Device token [%s] deleted in dict_device and free.", str_channel_token);
//     free (data); data = NULL;
//     return PLEASE_REMOVE_HASH_NODE;
// }

// static int cb_ttl_token(void *data, void *arg)
// {
//     // char str_channel_token[2 * CHANNEL_TOKEN_LEN + 1] = {0};
//     // convert_int2string((uint8_t *)data, CHANNEL_TOKEN_LEN, str_channel_token);
//     // TLOGD("cb_ttl_token", "", "Device token [%s] ttled, and free.", str_channel_token);
//     free (data); data = NULL;
//     return PLEASE_REMOVE_HASH_NODE;
// }

AtomicHashMap<uint64_t, device_t*> gMap(100000000);

int init_device_token_dict()
{
    // dict_device = atomic_hash_create(100000000, TTL);
    // if (dict_device == NULL)
    // {
    //     // TLOGE("init_device_token_dict", "", "token dict init failed!");
    //     return COMMON_FAIL;
    // }
    // dict_device->on_add = cb_add_token;
    // dict_device->on_dup = cb_dup_token;
    // dict_device->on_get = cb_get_token;
    // dict_device->on_del = cb_del_token;
    // dict_device->on_ttl = cb_ttl_token;



    // TLOGI("init_device_token_dict", "", "Device token dict init succeed.");
    return COMMON_SUCCESS;
}

int add_device_token()
{
    int i = 0;
    for (; i < 10000; i++)
    {
        device_t *device = (device_t *)malloc(sizeof(device_t));
        device->ip = (uint32_t)i;
        device->port = (uint16_t)i;
        device->unused = 0;
        // int ret = atomic_hash_add(dict_device, (device_key_t *)device, sizeof(device_key_t), device, TTL, NULL, NULL);
        gMap.insert(std::make_pair(*(uint64_t *)device, device));
        // if (ret != 0)
        // {
        //     return ret;
        // }
    }
    // for (i = 0; i < 10000; i++)
    // {
    //     device_t *device = (device_t *)malloc(sizeof(device_t));
    //     device->ip = (uint32_t)i;
    //     device->port = (uint16_t)i;
    //     int ret = atomic_hash_add(dict_device2, (device_key_t *)device, sizeof(device_key_t), device, TTL, NULL, NULL);
    //     if (ret != 0)
    //     {
    //         return ret;
    //     }
    // }
    return 0;
}

void *loop_get_device_token(void *arg)
{
    int total = 1000000;
    int failed = 0;
    int i = 0;
    for (; i < total; i++)
    {
        int key = i % 10000;
        device_key_t device_key;
        device_key.ip = (uint32_t)key;
        device_key.port = (uint16_t)key;
        device_key.unused = 0;
        // device_t *retv = NULL;
        // int ret = atomic_hash_get(dict_device, &device_key, sizeof(device_key_t), NULL, &retv);
        // if (ret != 0 || retv == NULL)
        // {
        //     failed++;
        // }
        device_t *ret = gMap.find(*(uint64_t *) &device_key)->second;
        if (ret->ip != device_key.ip || ret->port != device_key.port)
        {
            printf("not match\n");
            failed++;
        }
    }
    printf("atomic hash get total %d, failed %d.\n", total, failed);
}

int main()
{
    init_device_token_dict();
    int ret = add_device_token();
    if (ret != 0)
    {
        return ret;
    }

    printf("init success\n");

    unsigned long long start = nstime();

    #define THREAD_NUM 10
    pthread_t ids[THREAD_NUM];
    int i = 0;
    for (; i < THREAD_NUM; i++)
    {
        pthread_create(&ids[i], NULL, loop_get_device_token, NULL);
    }
    for (i = 0; i < THREAD_NUM; i++)
    {
        pthread_join(ids[i], NULL);
    }
    printf("threads finish: elapsed: %lld ns; avg: %llu ns/op\n", nstime()-start, (nstime()-start)/1000000);

    getchar();
}

// $ ./a.out 
// init success
// atomic hash get total 1000000, failed 0.
// atomic hash get total 1000000, failed 0.
// atomic hash get total 1000000, failed 0.
// atomic hash get total 1000000, failed 0.
// atomic hash get total 1000000, failed 0.
// atomic hash get total 1000000, failed 0.
// atomic hash get total 1000000, failed 0.
// atomic hash get total 1000000, failed 0.
// atomic hash get total 1000000, failed 0.
// atomic hash get total 1000000, failed 0.
// threads finish: elapsed: 284007168 ns; avg: 284 ns/op

// $ g++ -O3 -std=c++14 -Ifolly_bin/include -Lfolly_bin/lib test.cpp -lfolly -lglog -ldl -ldouble-conversion -pthread
// chuqq@chuqq-hp:~/temp/folly$ ./a.out 
// init success
// atomic hash get total 1000000, failed 0.
// atomic hash get total 1000000, failed 0.
// atomic hash get total 1000000, failed 0.
// atomic hash get total 1000000, failed 0.
// atomic hash get total 1000000, failed 0.
// atomic hash get total 1000000, failed 0.
// atomic hash get total 1000000, failed 0.
// atomic hash get total 1000000, failed 0.
// atomic hash get total 1000000, failed 0.
// atomic hash get total 1000000, failed 0.
// threads finish: elapsed: 16000256 ns; avg: 16 ns/op

