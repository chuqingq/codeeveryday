// g++ -O3 -std=c++11 condition_var.cpp -pthread

#include <iostream>
#include <string>
#include <thread>
#include <mutex>
#include <condition_variable>
 
#include <time.h>
#include <unistd.h>

static unsigned long long nstime(void) {
    struct timespec ts;
    clock_gettime(CLOCK_REALTIME, &ts);
    return ((unsigned long long)ts.tv_sec)*1e9 + ts.tv_nsec;
}

unsigned long long starttime, stoptime;

std::mutex m;
std::condition_variable cv;
std::string data;
bool ready = false;
bool processed = false;
 
void worker_thread()
{
    // Wait until main() sends data
    std::unique_lock<std::mutex> lk(m);
    // std::mutex * lk = &m;
    cv.wait(lk/*, []{return ready;}*/);
    // cv.wait(lk);
    std::cout << "worker: " << nstime() - starttime << std::endl;
 
    // after the wait, we own the lock.
    // std::cout << "Worker thread is processing data\n";
    // data += " after processing";
 
    // Send data back to main()
    processed = true;
    // std::cout << "Worker thread signals data processing completed\n";
 
    // Manual unlocking is done before notifying, to avoid waking up
    // the waiting thread only to block again (see notify_one for details)
    lk.unlock();
    cv.notify_one();
    starttime = nstime();
    // std::cout << "worker: " << starttime << std::endl;
}
 
int main()
{
    std::thread worker(worker_thread);
 
    // data = "Example data";
    // send data to the worker thread
    // {
        // std::lock_guard<std::mutex> lk(m);
        // ready = true;
        // std::cout << "main() signals data ready for processing\n";
        // starttime = nstime();
        // m.unlock();
    // }
    // std::cout << "main: " << starttime << std::endl;    
    sleep(1);
    starttime = nstime();
    cv.notify_one();
 
    // wait for the worker
    {
        std::unique_lock<std::mutex> lk(m);
        cv.wait(lk, []{return processed;});
    }
    std::cout << "main: " << nstime() - starttime << std::endl;
    // std::cout << "Back in main(), data = " << data << '\n';
 
    worker.join();
}

/*
chuqq@chuqq:~/work/codeeveryday/c++/20180310_condition_var$ g++ -O3 -std=c++11 condition_var.cpp -pthread
chuqq@chuqq:~/work/codeeveryday/c++/20180310_condition_var$
chuqq@chuqq:~/work/codeeveryday/c++/20180310_condition_var$
chuqq@chuqq:~/work/codeeveryday/c++/20180310_condition_var$
chuqq@chuqq:~/work/codeeveryday/c++/20180310_condition_var$ ./a.out
worker: 47104
main: 40960
chuqq@chuqq:~/work/codeeveryday/c++/20180310_condition_var$ ./a.out
worker: 51456
main: 42752
chuqq@chuqq:~/work/codeeveryday/c++/20180310_condition_var$ ./a.out
worker: 62208
main: 389888
chuqq@chuqq:~/work/codeeveryday/c++/20180310_condition_var$ ./a.out
worker: 57856
main: 84480
chuqq@chuqq:~/work/codeeveryday/c++/20180310_condition_var$ ./a.out
worker: 64768
main: 44800
chuqq@chuqq:~/work/codeeveryday/c++/20180310_condition_var$ ./a.out
worker: 64768
main: 40704
结论：大约40us，futex和pthread_cond_x大约20us，好像还是慢一些
*/
