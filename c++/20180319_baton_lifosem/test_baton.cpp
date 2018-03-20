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
环境：chuqq@hp
chuqq@chuqq-hp:~/temp/codeeveryday/c++/20180319_baton$ g++ test_baton.cpp -std=c++14 -O3 -I/home/chuqq/temp/folly/folly_bin/include -L/home/chuqq/temp/folly/folly_bin/lib -lfolly -lglog -ldl -ldouble-conversion -pthread
chuqq@chuqq-hp:~/temp/codeeveryday/c++/20180319_baton$ ./a.out 
elapsed: 37202 ns
chuqq@chuqq-hp:~/temp/codeeveryday/c++/20180319_baton$ ./a.out 
elapsed: 36375 ns
chuqq@chuqq-hp:~/temp/codeeveryday/c++/20180319_baton$ ./a.out 
elapsed: 40041 ns
chuqq@chuqq-hp:~/temp/codeeveryday/c++/20180319_baton$ ./a.out 
elapsed: 11221 ns
chuqq@chuqq-hp:~/temp/codeeveryday/c++/20180319_baton$ ./a.out 
elapsed: 36328 ns
chuqq@chuqq-hp:~/temp/codeeveryday/c++/20180319_baton$ ./a.out 
elapsed: 37874 ns
chuqq@chuqq-hp:~/temp/codeeveryday/c++/20180319_baton$ ./a.out 
elapsed: 37060 ns
chuqq@chuqq-hp:~/temp/codeeveryday/c++/20180319_baton$ ./a.out 
elapsed: 37398 ns
结论：很稳定，大约37us。和lifosem一致。
*/
