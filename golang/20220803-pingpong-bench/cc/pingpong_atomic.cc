#include <iostream>
#include <atomic>
#include <chrono>
#include <thread>

// #include "my_common.h"
const uint32_t attempts = 10000000;

// void PingPong(uint32_t attempts) {
int main() {
  std::cout << "pingpong" << std::endl;

  auto startTime = std::chrono::steady_clock::now();

  std::atomic<uint32_t> vPong{attempts};
  std::atomic<uint32_t> vPing{attempts};

  std::thread ping([&] {
    for (auto i = 0; i < attempts; ++i) {
      vPong.store(i);
      while (vPing.load() != i) {
      }
    }
  });

  std::thread pong([&] {
    for (auto i = 0; i < attempts; ++i) {
      while (vPong.load() != i) {
      }
      vPing.store(i);
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

  std::cout << "pingpong test times:" << attempts
            << " duration: " << duration.count() << "ms " << perDuration
            << "ns/op" << std::endl;
}

/*
chuqq@arch-vb~/t/p/c/g/2/cc $ g++ pingpong_atomic.cc -std=c++11 -O3
chuqq@arch-vb~/t/p/c/g/2/cc $ ./a.out 
pingpong
pingpong test times:10000000 duration: 921ms 92ns/op

*/
