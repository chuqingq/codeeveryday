```bat
set GOOS=linux
set GOARCH=arm

cd go/src
make.bat

cd ..
cp bin/%GOOS%_%GOARCH% %GOROOT%/bin -r
cp pkg/%GOOS%_%GOARCH% %GOROOT%/pkg -r
cp pkg/tool/%GOOS%_%GOARCH% %GOROOT%/pkg/tool -r
```