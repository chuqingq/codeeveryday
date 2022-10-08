#include <atomic>
#include <chrono>
#include <condition_variable>
#include <mutex>
#include <thread>

#include <iostream>

#include <folly/synchronization/SaturatingSemaphore.h>
using folly::SaturatingSemaphore;

void PingPong2(uint32_t numRounds) {
  std::cout << "pingpong" << std::endl;

  auto startTime = std::chrono::steady_clock::now();

  uint32_t pingTimes = numRounds;
  uint32_t pongTimes = numRounds;

  using WF = SaturatingSemaphore<true, std::atomic>;
  std::array<WF, 17> flags;
  WF& a = flags[0];
  WF& b = flags[16]; // different cache line

  std::thread ping([&] {
    for (int i = 0; i < numRounds; ++i) {
      pongTimes = i;

      a.post();
      // b.try_wait();
      // if (b.try_wait_until(std::chrono::steady_clock::now() + std::chrono::microseconds(1))) {
      //   printf("error\n");
      // }
      b.wait();
      b.reset();

      if (pingTimes != i) {
        printf("error2\n");
      }
    }
  });

  std::thread pong([&] {
    for (int i = 0; i < numRounds; ++i) {
      // a.try_wait();
      // if (a.try_wait_until(std::chrono::steady_clock::now() + std::chrono::microseconds(1))) {
      //   printf("error\n");
      // }
      a.wait();
      a.reset();
      if (pongTimes != i) {
        printf("error3\n");
      }

      pingTimes = i;
      b.post();
    }
  });

  // std::cout << "pingpong is running!" << std::endl;

  ping.join();
  pong.join();

  auto endTime = std::chrono::steady_clock::now();

  auto duration = std::chrono::duration_cast<std::chrono::milliseconds>(
      endTime - startTime);

  auto perDuration =
      std::chrono::duration_cast<std::chrono::nanoseconds>(endTime - startTime)
          .count() /
      numRounds;

  std::cout << "Using condtion_variable pingpong test times:" << numRounds
            << " duration: " << duration.count() << "ms " << perDuration
            << "ns/op" << std::endl;
}

int main() {
    const uint32_t attempts = 10000000;
    PingPong2(attempts);
    return 0;
}

/*
$ g++ chuqq_pingpong_folly.cpp -o chuqq_pingpong_folly -I. -I.. -L. -lfolly -ldouble-conversion -lglog -lgflags -lfmt

版本1：只wait，不try：
chuqq@chuqq-r7000p/m/d/t/f/build $ ./chuqq_pingpong_folly
pingpong
Using condtion_variable pingpong test times:10000000 duration: 1859ms 185ns/op


*/

