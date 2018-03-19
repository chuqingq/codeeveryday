// g++ test_mutex.cpp -std=c++14 -O3 -Ifolly_bin/include -Lfolly_bin/lib -lfolly -lglog -ldl -ldouble-conversion -pthread

#include <iostream>
#include <thread>
#include <chrono>
#include <mutex>

#include <folly/synchronization/MicroSpinLock.h>
using namespace folly;

int main() {
    std::mutex lock;
    lock.lock();
    auto start = std::chrono::steady_clock::now();

    std::thread t([&] {
        lock.lock();
        auto stop = std::chrono::steady_clock::now();
        std::cout << "elapsed: " << std::chrono::duration_cast<std::chrono::nanoseconds>(stop - start).count() << " ns" << std::endl;
    });

    std::this_thread::sleep_for(std::chrono::milliseconds(1000));

    start = std::chrono::steady_clock::now();
    lock.unlock();

    t.join();
    return 0;
}

/*
chuqq@chuqq:~/temp/folly$ g++ test_mutex.cpp -std=c++14 -O3 -Ifolly_bin/include -Lfolly_bin/lib -lfolly -lglog -ldl -ldouble-conversion -pthread
chuqq@chuqq:~/temp/folly$ ./a.out
elapsed: 152870 ns
chuqq@chuqq:~/temp/folly$ ./a.out
elapsed: 20625 ns
chuqq@chuqq:~/temp/folly$ ./a.out
elapsed: 29521 ns
chuqq@chuqq:~/temp/folly$ ./a.out
elapsed: 28865 ns
chuqq@chuqq:~/temp/folly$ ./a.out
elapsed: 80340 ns
chuqq@chuqq:~/temp/folly$ ./a.out
elapsed: 33831 ns
chuqq@chuqq:~/temp/folly$ ./a.out
elapsed: 151835 ns
*/
