# https://www.darkcoding.net/software/building-shared-libraries-in-go-part-1/
go build -v -x -buildmode=c-shared -o libmy.so
gcc -o test c/test.c -I. -L. libmy.so
LD_LIBRARY_PATH=`pwd` ./test
# 在8080端口启动了/hello的http服务
#str: Hello
#n: 5
