# https://www.darkcoding.net/software/building-shared-libraries-in-go-part-1/
go build -v -x -buildmode=c-shared -o libmy.so
gcc -o test c/test.c -I. -L. libmy.so
LD_LIBRARY_PATH=`pwd` ./test
#str: Hello
#n: 5
