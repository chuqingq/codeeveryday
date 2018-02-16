// g++ -static -static-libgcc -static-libstdc++  -std=c++11 test3.cpp -Wl,--whole-archive -lpthread -Wl,--no-whole-archive
#include <iostream>
#include <thread>

void foo() {
	std::cout << "hello world" << std::endl;
}

int main() {
	std::thread t(foo);
	t.join();
	return 0;
}
