#include <iostream>
#include <thread>
#include <chrono>

/*
int add(int a, int b) {
    std::cout << a << "+" << b << std::endl;
    return a+b;
}
*/

#ifdef __cplusplus 
extern "C" { 
#endif

int add(int a, int b) {
    std::this_thread::sleep_for(std::chrono::seconds(2));
    std::cout << a << "+" << b << std::endl;
    return a+b;
}

#ifdef __cplusplus 
} 
#endif
