# prepare folly_bin
cp /usr/lib/gcc/x86_64-linux-gnu/5/libstdc++.a libstdcxx.a
ar -M < libmylib.mri
g++ -std=c++14 -g -O3 test.cpp -I/home/chuqq/temp/folly/folly_bin/include -L. -lmylib -pthread -ldl
