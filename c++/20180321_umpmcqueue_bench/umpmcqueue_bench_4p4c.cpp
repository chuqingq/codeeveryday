// g++ umpmcqueue_bench_4p4c.cpp -std=c++14 -Wall -faligned-new -O3 -I/home/chuqq/temp/folly/folly_bin/include -L/home/chuqq/temp/folly/folly_bin/lib -lfolly -lglog -ldl -ldouble-conversion -pthread
#include <folly/executors/task_queue/UnboundedBlockingQueue.h>

#include <thread>
#include <chrono>
#include <iostream>

int main(int argc, char const *argv[])
{
    folly::UnboundedBlockingQueue<int> q;

    const int count = 100000000;

    std::thread c1([&] {
        for (int i = 0; i < count/4; ++i)
        {
            q.take();
        }
    });

    std::thread c2([&] {
        for (int i = 0; i < count/4; ++i)
        {
            q.take();
        }
    });

    std::thread c3([&] {
        for (int i = 0; i < count/4; ++i)
        {
            q.take();
        }
    });

    std::thread c4([&] {
        for (int i = 0; i < count/4; ++i)
        {
            q.take();
        }
    });

    // std::this_thread::sleep_for(std::chrono::seconds(10));
    auto start = std::chrono::steady_clock::now();

    std::thread p1([&] {
        for (int i = 0; i < count/4; ++i)
        {
            q.add(i);
        }
    });

    std::thread p2([&] {
        for (int i = 0; i < count/4; ++i)
        {
            q.add(i);
        }
    });

    std::thread p3([&] {
        for (int i = 0; i < count/4; ++i)
        {
            q.add(i);
        }
    });

    std::thread p4([&] {
        for (int i = 0; i < count/4; ++i)
        {
            q.add(i);
        }
    });
    
    c1.join();
    p1.join();

    c2.join();
    p2.join();

    c3.join();
    p3.join();

    c4.join();
    p4.join();

    if (q.size() != 0) {
        std::cout << "ERROR: q.size(): " << q.size() << std::endl;
    }

    auto stop = std::chrono::steady_clock::now();
    std::cout << "count: " << count << std::endl;
    std::cout << "elapsed: " << std::chrono::duration_cast<std::chrono::nanoseconds>(stop - start).count() << " ns" << std::endl;

    return 0;
}
/*
原dpdk-rte_ring的1e8+4p4c大约12秒：
SIZE=1<<10
chuqq@chuqq-hp:~/temp/dpdk-rte_ring$ ./rte_ring_main_4p4c 
complete: count=100000000, ns diff=12220387584

umpmcqueue的1e8的4p4c大约7.5秒：
chuqq@chuqq-hp ~/t/c/c/20180321_umpmcqueue_bench> g++ umpmcqueue_bench_4p4c.cpp -std=c++14 -Wall -faligned-new -O3 -I/home/chuqq/temp/folly/folly_bin/include -L/home/chuqq/temp/folly/folly_bin/lib -lfolly -lglog -ldl -ldouble-conversion -pthread
chuqq@chuqq-hp ~/t/c/c/20180321_umpmcqueue_bench> time ./a.out
count: 100000000
elapsed: 7534004338 ns
53.32user 0.18system 0:07.54elapsed 709%CPU (0avgtext+0avgdata 286260maxresident)k
0inputs+0outputs (0major+74581minor)pagefaults 0swaps
chuqq@chuqq-hp ~/t/c/c/20180321_umpmcqueue_bench> time ./a.out
count: 100000000
elapsed: 7722776377 ns
55.11user 0.16system 0:07.72elapsed 715%CPU (0avgtext+0avgdata 281220maxresident)k
0inputs+0outputs (0major+72415minor)pagefaults 0swaps
chuqq@chuqq-hp ~/t/c/c/20180321_umpmcqueue_bench> time ./a.out
count: 100000000
elapsed: 7373156676 ns
54.41user 0.16system 0:07.38elapsed 739%CPU (0avgtext+0avgdata 197964maxresident)k
0inputs+0outputs (0major+50220minor)pagefaults 0swaps
chuqq@chuqq-hp ~/t/c/c/20180321_umpmcqueue_bench> time ./a.out
count: 100000000
elapsed: 7721659198 ns
51.22user 0.21system 0:07.72elapsed 665%CPU (0avgtext+0avgdata 428872maxresident)k
0inputs+0outputs (0major+108038minor)pagefaults 0swaps

结论：感觉效果比rte_ring好；而且空闲时不会占用cpu；从空闲转到繁忙也不慢

*/
