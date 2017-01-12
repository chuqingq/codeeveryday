go build -v -x -buildmode=c-shared -o lib.so
gcc -o test test.c -L. lib.so
LD_LIBRARY_PATH=/home/chuqq ./test
#str: Hello
#n: 5
