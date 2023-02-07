#include <atomic>
#include <chrono>
#include <condition_variable>
#include <mutex>
#include <thread>

#include <iostream>
//#include "my_common.h"

void PingPong2(uint32_t attempts) {
  std::cout << "pingpong" << std::endl;

  auto startTime = std::chrono::steady_clock::now();

  uint32_t pingTimes = attempts;
  uint32_t pongTimes = attempts;

  std::mutex pingMutex;
  std::condition_variable pingVariable;

  std::mutex pongMutex;
  std::condition_variable pongVariable;

  std::thread ping([&] {
    for (auto i = 0; i < attempts; ++i) {
      {
        std::lock_guard<std::mutex> lk(pongMutex);
        pongTimes = i;
      }
      pongVariable.notify_one();

    retry:
      while (!pingMutex.try_lock());
      if (pingTimes != i) {
        pingMutex.unlock();
        goto retry;
      } else {
        pingMutex.unlock();
        continue;
      }

      // std::unique_lock<std::mutex> pingLock(pingMutex);
      // pingVariable.wait(pingLock, [&] { return (pingTimes == i); });
      // pingLock.unlock();
    }
  });

  std::thread pong([&] {
    for (auto i = 0; i < attempts; ++i) {

    retry2:
      while (!pongMutex.try_lock());
      if (pongTimes != i) {
        pongMutex.unlock();
        goto retry2;
      } else {
        pongMutex.unlock();
        // continue;
      }

      // std::unique_lock<std::mutex> pongLock(pongMutex);
      // pongVariable.wait(pongLock, [&] { return (pongTimes == i); });
      // pongLock.unlock();

      {
        std::lock_guard<std::mutex> lk(pingMutex);
        pingTimes = i;
      }
      pingVariable.notify_one();
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
      attempts;

  std::cout << "Using condtion_variable pingpong test times:" << attempts
            << " duration: " << duration.count() << "ms " << perDuration
            << "ns/op" << std::endl;
}

int main() {
    const uint32_t attempts = 1000000;
    PingPong2(attempts);
    return 0;
}

/*
Busy loop on pingpong.cc.

$ ./pingpong2
pingpong
Using condtion_variable pingpong test times:1000000 duration: 412ms 412ns/op
*/

