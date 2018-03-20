// g++ test_microspinlock.cpp -std=c++14 -O3 -Ifolly_bin/include -Lfolly_bin/lib -lfolly -lglog -ldl -ldouble-conversion -pthread

#include <iostream>
#include <thread>
#include <chrono>

#include <folly/synchronization/MicroSpinLock.h>
using namespace folly;

int main() {
    MicroSpinLock lock;
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
*/
