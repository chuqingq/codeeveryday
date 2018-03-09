# prepare folly_bin
cp /usr/lib/gcc/x86_64-linux-gnu/7/libstdc++.a libstdcxx.a
ar -M < libmylib.mri
g++ -std=c++14 -g -O3 test.cpp -Ifolly_bin/include -L. -lmylib -pthread -ldl
