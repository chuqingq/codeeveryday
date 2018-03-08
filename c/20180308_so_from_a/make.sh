gcc -c add.c
ar -rc libadd.a add.o

gcc -o libadd2.so add2.c -fPIC -shared -L. -ladd

gcc -o main main.c -L. -ladd2

LD_LIBRARY_PATH=`pwd` ./main
# 3

