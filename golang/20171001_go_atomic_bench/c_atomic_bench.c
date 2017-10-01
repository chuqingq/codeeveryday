#include <stdio.h>
#include <sys/time.h>

static long long ustime(void) {
    struct timeval tv;
    long long ust;

    gettimeofday(&tv, NULL);
    ust = ((long)tv.tv_sec)*1000000;
    ust += tv.tv_usec;
    return ust;
}

int main() {
	long long start, stop;
	int value;

	start = ustime();
	for (int i = 0; i < 1000000000; i++) {
		value++;
	}
	stop = ustime();

	printf("%d\n", value);
	printf("%lld\n", stop-start);
	printf("use value++: %.6f\n", (stop-start)/1000000000.0);


	value = 0;
	start = ustime();
        for (int i = 0; i < 1000000000; i++) {
		__sync_fetch_and_add(&value, 1);;
        }
        stop = ustime();

        printf("%d\n", value);
        printf("%lld\n", stop-start);
        printf("use atomicadd: %.6f\n", (stop-start)/1000000000.0);
}
