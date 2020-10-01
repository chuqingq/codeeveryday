gcc -c myfunc.c
ar cr libmyfunc.a myfunc.o
go build myfunc.go
./myfunc
