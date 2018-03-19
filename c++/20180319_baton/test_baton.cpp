// g++ test_baton.cpp -std=c++14 -O3 -Ifolly_bin/include -Lfolly_bin/lib -lfolly -lglog -ldl -ldouble-conversion -pthread

#include <iostream>
#include <thread>
#include <chrono>

#include <folly/synchronization/Baton.h>
using namespace folly;

int main() {
    Baton<> b;
    auto start = std::chrono::steady_clock::now();

    std::thread t([&] {
        b.wait();
        auto stop = std::chrono::steady_clock::now();
        std::cout << "elapsed: " << std::chrono::duration_cast<std::chrono::nanoseconds>(stop - start).count() << " ns" << std::endl;
    });

    std::this_thread::sleep_for(std::chrono::milliseconds(1000));

    start = std::chrono::steady_clock::now();
    b.post();

    t.join();
    return 0;
}

/*
chuqq@chuqq:~/temp/folly$ g++ test_baton.cpp -std=c++14 -O3 -Ifolly_bin/include -Lfolly_bin/lib -lfolly -lglog -ldl -ldouble-conversion -pthread
chuqq@chuqq:~/temp/folly$ ./a.out
elapsed: 107167 ns
chuqq@chuqq:~/temp/folly$ ./a.out
elapsed: 31814 ns
chuqq@chuqq:~/temp/folly$ ./a.out
elapsed: 30605 ns
chuqq@chuqq:~/temp/folly$ ./a.out
elapsed: 30303 ns
chuqq@chuqq:~/temp/folly$ ./a.out
elapsed: 35084 ns
chuqq@chuqq:~/temp/folly$ ./a.out
elapsed: 21251 ns
chuqq@chuqq:~/temp/folly$ ./a.out
elapsed: 194283 ns
chuqq@chuqq:~/temp/folly$ ./a.out
elapsed: 133229 ns
*/
