// g++ system_clock.cpp -std=c++11
#include <iostream>
#include <chrono>   
// using namespace std;
// using namespace chrono;


int main() {
  /*
  auto start = std::chrono::system_clock::now();
  std::time_t start_c = std::chrono::system_clock::to_time_t(start);
  std::cout << "start: " << std::put_time(std::localtime(&start_c), "%F %T") << std::endl;
  */
  auto start = std::chrono::steady_clock::now();
  // std::cout << "Hello World\n";
  auto end = std::chrono::steady_clock::now();
  std::cout << "Printing took "
            << std::chrono::duration_cast<std::chrono::microseconds>(end - start).count()
            << "us.\n";
  std::cout << "Printing took "
            << std::chrono::duration_cast<std::chrono::nanoseconds>(end - start).count()
            << "ns.\n";
  return 0;
}

/*
chuqq@chuqq-hp:~/temp/codeeveryday/c++/20171020_system_clock$ ./a.out 
Printing took 0us.
Printing took 250ns.
chuqq@chuqq-hp:~/temp/codeeveryday/c++/20171020_system_clock$ ./a.out 
Printing took 0us.
Printing took 329ns.
chuqq@chuqq-hp:~/temp/codeeveryday/c++/20171020_system_clock$ ./a.out 
Printing took 0us.
Printing took 288ns.
chuqq@chuqq-hp:~/temp/codeeveryday/c++/20171020_system_clock$ ./a.out 
Printing took 0us.
Printing took 253ns.
chuqq@chuqq-hp:~/temp/codeeveryday/c++/20171020_system_clock$ ./a.out 
Printing took 0us.
Printing took 220ns.
chuqq@chuqq-hp:~/temp/codeeveryday/c++/20171020_system_clock$ ./a.out 
Printing took 0us.
Printing took 266ns.
*/
