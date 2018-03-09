#include <iostream>

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
    std::cout << a << "+" << b << std::endl;
    return a+b;
}

#ifdef __cplusplus 
} 
#endif
