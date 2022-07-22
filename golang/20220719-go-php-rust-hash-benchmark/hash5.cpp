#include <iostream>
#include <chrono>
#include <thread>
#include <string>
#include <unordered_map>
#include <map>

using namespace std;

// static long mstime() {
//     return std::chrono::duration_cast<std::chrono::milliseconds>(std::chrono::steady_clock::now().time_since_epoch()).count();
// }

#include <sys/time.h>
#include <unistd.h>

inline unsigned long long mstime(void) {
    struct timeval tv;
    gettimeofday(&tv, NULL);
    return ((unsigned long long)tv.tv_sec)*1000 + tv.tv_usec/1000;
}

int main() {
    auto t1 = mstime();
    std::unordered_map<string, long> arr;
	
	arr.reserve(1000000);
    
	for (auto i = 0; i < 1000000; i++) {	
    	auto value = mstime()/1000;
		arr[to_string(i) + string("_") + to_string( value)] = value;
	}

	auto t2 = mstime();
    printf("%llu ms, count: %lu\n", (t2 - t1), arr.size());

	return 0;
}

// g++ hash5.cpp -o hash5cpp -O3 -std=c++11 -Wall -static
// 315 ms
