// 1. git clone ...
// 2. 安装依赖
// sudo apt-get install \
//     g++ \
//     cmake \
//     libboost-all-dev \
//     libevent-dev \
//     libdouble-conversion-dev \
//     libgoogle-glog-dev \
//     libgflags-dev \
//     libiberty-dev \
//     liblz4-dev \
//     liblzma-dev \
//     libsnappy-dev \
//     make \
//     zlib1g-dev \
//     binutils-dev \
//     libjemalloc-dev \
//     libssl-dev \
//     pkg-config
// 3. cmake生成makefile
//     cmake -DCMAKE_INSTALL_PREFIX=/home/chuqq/temp/folly/folly_bin .
// 4. make && make install
// 5. g++ -std=c++14 -Ifolly_bin/include -Lfolly_bin/lib fbstring_sample.cpp

#include <iostream>
#include <folly/FBString.h>

int main() {
    folly::fbstring str("hello world");
    std::cout << str << std::endl;
    return 0;
}