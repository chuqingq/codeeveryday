# https://www.darkcoding.net/software/building-shared-libraries-in-go-part-1/
go build -v -x -buildmode=c-archive -o libmy.a
gcc -o test c/test.c -static -I. -L. -lmy -pthread
./test
#str: Hello
#n: 5
