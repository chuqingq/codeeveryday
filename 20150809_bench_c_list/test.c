#include <stdio.h>
#include <stdlib.h>
#include <sys/time.h>


static long long timeInMilliseconds(void) {
    struct timeval tv;

    gettimeofday(&tv,NULL);
    return (((long long)tv.tv_sec)*1000)+(tv.tv_usec/1000);
}


typedef struct node_s {
    struct node_s* next;
    long long score;
    void* value;
} node_t;


int main(int argc, char const *argv[])
{
    int count = 50000000;
    node_t head;
    node_t* prev = &head;
    node_t* mynode;
    long long time1, time2, time3;

    int i;
    printf("sizeof(node_t): %ld\n", sizeof(node_t));
    printf("sizeof(node_t): %ld\n", count*sizeof(node_t)/1024/1024);
    time1 = timeInMilliseconds();
    for (i = 0; i < count; i++) {
    	mynode = malloc(sizeof(*mynode));
    	if (!mynode) return -1;

        mynode->score=(long long) i;
    	prev->next = mynode;
    	prev = mynode;
    }

    time2 = timeInMilliseconds();
    printf("malloc millisec: %lld\n", time2-time1);

    mynode = head.next;
    while (!mynode && mynode->score >= 0) {
        mynode = mynode->next;
    }

    time3 = timeInMilliseconds();
    printf("while millisec: %lld\n", time3-time2);

    getchar();
	return 0;
}
/*
count=1000w

sizeof(node_t): 24
sizeof(node_t): 228
malloc millisec: 2310
while millisec: 0


====
count = 5000w

sizeof(node_t): 24
sizeof(node_t): 1144
malloc millisec: 14625
while millisec: 0
*/