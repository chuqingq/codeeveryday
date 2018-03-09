# main.c依赖libadd.so，libadd.so中包含了libstdc++

g++ -o libadd.so add.c -fPIC -shared -static-libstdc++

# nm libadd.so 检查发现c++的符号都不是U，也就是已经resolved，也就是已经链接进去了 

gcc -o main main.c -L. -ladd

LD_LIBRARY_PATH=`pwd` ./main
# 1+2
# 3

