先确保goshared目录在$GOPATH下

```
go install -v -x -buildmode=shared runtime sync/atomic #构建核心基本库
go install -v -x -buildmode=shared -linkshared #构建GO动态库
```

```
go build -v -x -linkshared #构架调用可执行文件
```

编译时依赖源码，执行时只需动态库。如果把$GOPATH//pkg/linux_amd64_dynlink/libgoshared.so删掉，执行失败
```
./go_shared: error while loading shared libraries: libgoshared.so: cannot open shared object file: No such file or directory
```
