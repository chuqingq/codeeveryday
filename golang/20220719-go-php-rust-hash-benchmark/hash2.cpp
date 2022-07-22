#include <iostream>
#include <chrono>
#include <thread>
#include <string>
#include <map>

using namespace std;

// static long mstime() {
//     return std::chrono::duration_cast<std::chrono::milliseconds>(std::chrono::steady_clock::now().time_since_epoch()).count();
// }

#include <sys/time.h>
#include <unistd.h>

static unsigned long long mstime(void) {
    struct timeval tv;
    gettimeofday(&tv, NULL);
    return ((unsigned long long)tv.tv_sec)*1000 + tv.tv_usec/1000;
}

int main() {
	auto t1 = mstime();
	map<string, long> arr;
	for (int i = 0; i < 1000000; i++) {
		long value = mstime()/1000;
		string key = to_string(i) + "_" + to_string(value);
		arr[key] = value;
	}
	auto t2 = mstime();
	printf("%d ms, count: %d\n", (t2 - t1), arr.size());
}