// g++ uspscqueue_delay.cpp -std=c++14 -Wall -faligned-new -O3 -I/home/chuqq/temp/folly/folly_bin/include -L/home/chuqq/temp/folly/folly_bin/lib -lfolly -lglog -ldl -ldouble-conversion -pthread
#include <folly/concurrency/UnboundedQueue.h>

#include <thread>
#include <chrono>
#include <iostream>

int main(int argc, char const *argv[])
{
    folly::USPSCQueue<int, false, 6> q;

    const int count = 100000000;

    std::thread c1([&] {
        int r;
        for (int i = 0; i < count; ++i)
        {
            q.dequeue(r);
        }
    });

    // std::this_thread::sleep_for(std::chrono::seconds(10));
    auto start = std::chrono::steady_clock::now();

    std::thread p1([&] {
        for (int i = 0; i < count; ++i)
        {
            q.enqueue(i);
        }
    });

    c1.join();
    p1.join();

    if (q.size() != 0) {
        std::cout << "ERROR: q.size(): " << q.size() << std::endl;
    }

    auto stop = std::chrono::steady_clock::now();
    std::cout << "count: " << count << std::endl;
    std::cout << "elapsed: " << std::chrono::duration_cast<std::chrono::nanoseconds>(stop - start).count() << " ns" << std::endl;

    return 0;
}
/*
chuqq@chuqq-hp ~/t/c/c/20180321_umpmcqueue_bench> ./a.out 
count: 100000000
elapsed: 1296336256 ns
*/
