GOOS=js GOARCH=wasm go build -o static/main.wasm main.go

cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" static

go run server.go

# 
使用log.Printf()会打印的浏览器console;
增加test()之后，也能编译js成功，但是浏览器console中打印日志：dial error: dial tcp: Protocol not available
