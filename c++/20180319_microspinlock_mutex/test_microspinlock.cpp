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
    // auto start = std::chrono::high_resolution_clock::now();

    std::thread t([&] {
        lock.lock();
        auto stop = std::chrono::steady_clock::now();
	// auto stop = std::chrono::high_resolution_clock::now();
        std::cout << "elapsed: " << std::chrono::duration_cast<std::chrono::nanoseconds>(stop - start).count() << " ns" << std::endl;
    });

    std::this_thread::sleep_for(std::chrono::milliseconds(0));

    start = std::chrono::steady_clock::now();
    // start = std::chrono::high_resolution_clock::now();
    lock.unlock();

    t.join();
    return 0;
}

/*
chuqq@chuqq-hp ~/t/c/c/20180319_microspinlock_mutex> ./a.out 
elapsed: 38761 ns
chuqq@chuqq-hp ~/t/c/c/20180319_microspinlock_mutex> ./a.out 
elapsed: 37325 ns
chuqq@chuqq-hp ~/t/c/c/20180319_microspinlock_mutex> ./a.out 
elapsed: 36449 ns
chuqq@chuqq-hp ~/t/c/c/20180319_microspinlock_mutex> ./a.out 
elapsed: 38162 ns
chuqq@chuqq-hp ~/t/c/c/20180319_microspinlock_mutex> ./a.out 
elapsed: 35964 ns
chuqq@chuqq-hp ~/t/c/c/20180319_microspinlock_mutex> ./a.out 
elapsed: 16696 ns
chuqq@chuqq-hp ~/t/c/c/20180319_microspinlock_mutex> ./a.out 
elapsed: 37801 ns
chuqq@chuqq-hp ~/t/c/c/20180319_microspinlock_mutex> ./a.out 
elapsed: 12028 ns
*/
