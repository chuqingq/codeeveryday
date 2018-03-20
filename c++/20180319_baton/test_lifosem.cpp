// g++ test_lifosem.cpp -std=c++14 -O3 -Ifolly_bin/include -Lfolly_bin/lib -lfolly -lglog -ldl -ldouble-conversion -pthread

#include <iostream>
#include <thread>
#include <chrono>

// #include <folly/synchronization/Baton.h>
#include <folly/synchronization/LifoSem.h>
using namespace folly;

int main() {
    LifoSem b;
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
chuqq@chuqq-hp:~/temp/codeeveryday/c++/20180319_baton$ g++ test_lifosem.cpp -std=c++14 -O3 -I/home/chuqq/temp/folly/folly_bin/include -L/home/chuqq/temp/folly/folly_bin/lib -lfolly -lglog -ldl -ldouble-conversion -pthread
chuqq@chuqq-hp:~/temp/codeeveryday/c++/20180319_baton$ ./a.out 
elapsed: 37910 ns
chuqq@chuqq-hp:~/temp/codeeveryday/c++/20180319_baton$ ./a.out 
elapsed: 37260 ns
chuqq@chuqq-hp:~/temp/codeeveryday/c++/20180319_baton$ ./a.out 
elapsed: 39367 ns
chuqq@chuqq-hp:~/temp/codeeveryday/c++/20180319_baton$ ./a.out 
elapsed: 37262 ns
chuqq@chuqq-hp:~/temp/codeeveryday/c++/20180319_baton$ ./a.out 
elapsed: 37486 ns
chuqq@chuqq-hp:~/temp/codeeveryday/c++/20180319_baton$ ./a.out 
elapsed: 37139 ns
chuqq@chuqq-hp:~/temp/codeeveryday/c++/20180319_baton$ ./a.out 
elapsed: 11902 ns
chuqq@chuqq-hp:~/temp/codeeveryday/c++/20180319_baton$ ./a.out 
elapsed: 11646 ns
chuqq@chuqq-hp:~/temp/codeeveryday/c++/20180319_baton$ ./a.out 
elapsed: 39247 ns
chuqq@chuqq-hp:~/temp/codeeveryday/c++/20180319_baton$ ./a.out 
elapsed: 36984 ns
结论：这个很稳定，在37us左右。
*/
