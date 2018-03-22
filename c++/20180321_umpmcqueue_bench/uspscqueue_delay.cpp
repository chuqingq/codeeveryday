// g++ uspscqueue_delay.cpp -std=c++14 -Wall -faligned-new -O3 -I/home/chuqq/temp/folly/folly_bin/include -L/home/chuqq/temp/folly/folly_bin/lib -lfolly -lglog -ldl -ldouble-conversion -pthread
#include <folly/concurrency/UnboundedQueue.h>

#include <thread>
#include <chrono>
#include <iostream>

int main(int argc, char const *argv[])
{
    folly::USPSCQueue<int, false, 6> q;

    auto start = std::chrono::steady_clock::now();

    std::thread c1([&] {
        int r;
        q.dequeue(r);
        auto stop = std::chrono::steady_clock::now();
        std::cout << "elapsed: " << std::chrono::duration_cast<std::chrono::nanoseconds>(stop - start).count() << " ns" << std::endl;
    });

    // std::this_thread::sleep_for(std::chrono::seconds(10));

    start = std::chrono::steady_clock::now();
    q.enqueue(1);

    c1.join();

    if (q.size() != 0) {
        std::cout << "ERROR: q.size(): " << q.size() << std::endl;
    }

    return 0;
}
/*
chuqq@chuqq-hp ~/t/c/c/20180321_umpmcqueue_bench> ./a.out 
elapsed: 36759 ns
chuqq@chuqq-hp ~/t/c/c/20180321_umpmcqueue_bench> ./a.out 
elapsed: 15967 ns
chuqq@chuqq-hp ~/t/c/c/20180321_umpmcqueue_bench> ./a.out 
elapsed: 39964 ns
chuqq@chuqq-hp ~/t/c/c/20180321_umpmcqueue_bench> ./a.out 
elapsed: 39365 ns
chuqq@chuqq-hp ~/t/c/c/20180321_umpmcqueue_bench> ./a.out 
elapsed: 37837 ns
chuqq@chuqq-hp ~/t/c/c/20180321_umpmcqueue_bench> ./a.out 
elapsed: 39480 ns
chuqq@chuqq-hp ~/t/c/c/20180321_umpmcqueue_bench> ./a.out 
elapsed: 38056 ns
chuqq@chuqq-hp ~/t/c/c/20180321_umpmcqueue_bench> ./a.out 
elapsed: 37062 ns
*/
