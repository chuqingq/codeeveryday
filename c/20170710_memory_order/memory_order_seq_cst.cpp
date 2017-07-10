// g++ -std=c++11 -o memory_order_seq_cst memory_order_seq_cst.cpp -pthread 
// 如果是relax内存序，可能出现z=0的情况；如果是seq_cst，要么是x先于y，要么y先于x，所以至少有一个会z++，而且可能x和y都变成true时两个z++才执行。
#include <atomic>
#include <thread>
#include <assert.h>

std::atomic<bool> x,y;
std::atomic<int> z;

void write_x()
{
  x.store(true,std::memory_order_seq_cst);
} // 1

void write_y()
{
  y.store(true,std::memory_order_seq_cst);
}  // 2

void read_x_then_y()
{
  while(!x.load(std::memory_order_seq_cst));
  if(y.load(std::memory_order_seq_cst))  // 3
    ++z;
}

void read_y_then_x()
{
  while(!y.load(std::memory_order_seq_cst));
  if(x.load(std::memory_order_seq_cst))  // 4
    ++z;
}

int main() {
  x=false;
  y=false;
  z=0;
  std::thread a(write_x);
  std::thread b(write_y);
  std::thread c(read_x_then_y);
  std::thread d(read_y_then_x);
  a.join();
  b.join();
  c.join();
  d.join();
  assert(z.load()==1||z.load()==2);  // 5
}
