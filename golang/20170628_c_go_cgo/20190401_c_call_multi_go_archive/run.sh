# https://www.darkcoding.net/software/building-shared-libraries-in-go-part-1/
cd libmy; go build -v -x -buildmode=c-archive -o libmy.a; cd ..
cd libmyadd; go build -v -x -buildmode=c-archive -o libmyadd.a; cd ..
#gcc -o test c/test.c -Ilibmy -Ilibmyadd -Llibmy -Llibmyadd libmy.so libmyadd.so
gcc -o test c/test.c -Ilibmy -Ilibmyadd -Llibmy -Llibmyadd -lmy -lmyadd
./test
# 报错：同一个函数
# runtime: address space conflict
#libmyadd/libmyadd.a(go.o): In function `_cgo_topofstack':
#/usr/lib/go-1.10/src/runtime/asm_amd64.s:2351: multiple definition of `_cgo_topofstack'
#libmy/libmy.a(go.o):/usr/lib/go-1.10/src/runtime/asm_amd64.s:2351: first defined here
#libmyadd/libmyadd.a(go.o): In function `_cgo_panic':
#/usr/lib/go-1.10/src/runtime/cgo/callbacks.go:45: multiple definition of `_cgo_panic'
#libmy/libmy.a(go.o):/usr/lib/go-1.10/src/runtime/cgo/callbacks.go:45: first defined here
#libmyadd/libmyadd.a(go.o): In function `crosscall2':
#/usr/lib/go-1.10/src/runtime/cgo/asm_amd64.s:11: multiple definition of `crosscall2'
#libmy/libmy.a(go.o):/usr/lib/go-1.10/src/runtime/cgo/asm_amd64.s:11: first defined here
#collect2: error: ld returned 1 exit status

# 结论：c不好同时调用两个go的a库

