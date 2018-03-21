#include <map>

static std::map<int, int> mymap;

extern "C" {

void my_set(int a, int b) {
    mymap[a] = b;
}

int my_get(int a) {
    return mymap[a];
}

}
