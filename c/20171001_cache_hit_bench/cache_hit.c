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

int main()
{
    int array[1024][1024];
    int num = 0;

    long long start, stop;

    // cache hit + prefetch
    // start = ustime();
    // for (int i = 0; i < 1024; ++i)
    // {
    //     __builtin_prefetch(&array[i+1][0]);
    //     for (int j = 0; j < 1024; ++j)
    //     {
    //         array[i][j] = num++;
    //     }
    // }
    // stop = ustime();
    // printf("cache hit + prefetch: %lld us\n", stop-start);

    // cache hit
    start = ustime();
    for (int i = 0; i < 1024; ++i)
    {
        for (int j = 0; j < 1024; ++j)
        {
            array[i][j] = num++;
        }
    }
    stop = ustime();
    printf("cache hit : %lld us\n", stop-start);

    // cache miss
    num = 0;
    start = ustime();
    for (int i = 0; i < 1024; ++i)
    {
        for (int j = 0; j < 1024; ++j)
        {
            array[j][i] = num++;
        }
    }
    stop = ustime();
    printf("cache miss: %lld us\n", stop-start);


    // cache miss + prefetch
    num = 0;
    start = ustime();
    for (int i = 0; i < 1024; ++i)
    {
        for (int j = 0; j < 1024; ++j)
        {
            array[j][i] = num++;
            __builtin_prefetch(&array[j+1][i]);
        }
    }
    stop = ustime();
    printf("cache miss+prefetch: %lld us\n", stop-start);

    return 0;
}

/*
$ ./a.out 
cache hit : 12782 us
cache miss: 25964 us
cache miss+prefetch: 23989 us
*/
