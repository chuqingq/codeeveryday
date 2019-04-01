# https://www.darkcoding.net/software/building-shared-libraries-in-go-part-1/
cd libmy; go build -v -x -buildmode=c-shared -o libmy.so; cd ..
cd libmyadd; go build -v -x -buildmode=c-shared -o libmyadd.so; cd ..
#gcc -o test c/test.c -Ilibmy -Ilibmyadd -Llibmy -Llibmyadd libmy.so libmyadd.so
gcc -o test c/test.c -Ilibmy -Ilibmyadd -Llibmy -Llibmyadd libmy/libmy.so libmyadd/libmyadd.so
./test
# 报错：
# runtime: address space conflict
# 结论：c不好同时调用两个go的so库

