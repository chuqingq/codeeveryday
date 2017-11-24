#include <stdio.h>
#include <fcntl.h>
#include <stdlib.h>
#include <unistd.h>
#include <map>
#include <sys/time.h>

// ----time---------------------------------------------------------------------
static long long ustime(void) {
    struct timeval tv;
    long long ust;

    gettimeofday(&tv, NULL);
    ust = ((long)tv.tv_sec)*1000000;
    ust += tv.tv_usec;
    return ust;
}


// ----random utils-------------------------------------------------------------
static int random_fd = -1;

int xrand(char *out, int len) 
{
    // TODO 要避免多线程同时首次调用xrand
    if (random_fd <= 0) {
        random_fd = open("/dev/urandom", O_RDONLY);
        if (random_fd < 0) {
            perror("xrand: open error");
            return random_fd;
        }
    }

    int ret = read(random_fd, out, len);
    if (ret == -1) {
        perror("xrand: read error");
        return ret;
    }
    return 0;
}

// ----msg----------------------------------------------------------------------

typedef struct msg_offset_t {
	long fileid;
	long offset;
	msg_offset_t* next;
} msg_offset_t;

typedef struct msg_t {
	long long channelid;
	int requestid;
	char msg[500];
} msg_t;


int main(int argc, char const *argv[])
{
	const long count = 10000000;
	// ---- 打印各个字段长度
	printf("count: %ld\n", count);
	printf("sizeof(msg_t): %lu\n", sizeof(msg_t));
	printf("sizeof(msg_offset_t): %lu\n", sizeof(msg_offset_t));

	// 创建全局变量
	std::map<long long, msg_offset_t*> devices;
	FILE *file = fopen("messages", "a+");

	// 随机生成channelid，保存到devices中，并追加到磁盘
	long long start = ustime();
	long long channelid;
	msg_t msg;
	for (long i = 0; i < count; ++i)
	{
		xrand((char*)&channelid, sizeof(channelid));

		msg_offset_t *offset = (msg_offset_t *)malloc(sizeof(*offset));
		offset->offset = i;

		msg_offset_t *old = devices[channelid];
		if (old == NULL) {
			// devices.insert(channelid, offset);
			devices[channelid] = offset;
		} else {
			old->next = offset;
		}

		msg.channelid = channelid;
		msg.requestid = 0;
		fwrite(&msg, sizeof(msg), 1, file);
	}

	fclose(file);
	long long speed = ((long long)count*1000*1000)/(ustime()-start);
	printf("finished write records! speed: %lld records/second, %lld MB/second\n", speed, speed*512/1024/1024);

	getchar();
	return 0;
}

/*

1000w数据，内存约900+MB

env: lbuntu ssd
$ ./a.out 
count: 10000000
sizeof(msg_t): 512
sizeof(msg_offset_t): 24
finished write records! speed: 295911 records/second, 144 MB/second

env: 虚机（225）
# ./a.out 
count: 10000000
sizeof(msg_t): 512
sizeof(msg_offset_t): 24
finished write records! speed: 96335 records/second, 47 MB/second
# ./a.out 
count: 10000000
sizeof(msg_t): 512
sizeof(msg_offset_t): 24
finished write records! speed: 104793 records/second, 51 MB/second

env: 虚机（1.38）
# ./a.out 
count: 10000000
sizeof(msg_t): 512
sizeof(msg_offset_t): 24
finished write records! speed: 149473 records/second, 72 MB/second
# ./a.out 
count: 10000000
sizeof(msg_t): 512
sizeof(msg_offset_t): 24
finished write records! speed: 148159 records/second, 72 MB/second

env: 物理机（14.17）
# ./a.out 
count: 10000000
sizeof(msg_t): 512
sizeof(msg_offset_t): 24
finished write records! speed: 164020 records/second, 80 MB/second
# ./a.out 
count: 10000000
sizeof(msg_t): 512
sizeof(msg_offset_t): 24
finished write records! speed: 165956 records/second, 81 MB/second

*/

