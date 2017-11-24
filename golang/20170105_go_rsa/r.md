https://stackoverflow.com/questions/22344149/to-use-rsa-pkcs1-oaep-padding-for-rsa-signature

# 用openssl验证

## 生成rsa私钥。我是从terminal.go中拷贝出来的
openssl genrsa -out private.key 2048

## 提取公钥
openssl rsa -in private.key -pubout -out public.key

## 生成明文数据
echo "Hello world" > input.data

## 加密
openssl rsautl -encrypt -oaep -inkey private.key -in input.data -out output.data
openssl rsautl -encrypt -oaep -inkey public.key -pubin  -in input.data -out output.data
这两个命令是一样的效果。

## 解密验证
openssl rsautl -decrypt -oaep -inkey private.key -in output.data -out back.data

## 可以验证多次。每次密文不同，但还原明文都和明文是一致的。

# 验证私钥加密、公钥解密
openssl rsautl -encrypt -oaep -inkey private.key -in input.data -out output.data
openssl rsautl -decrypt -oaep -inkey public.key -pubin -in output.data -out back.data
报错：
A private key is needed for this operation

结论：私钥加密、公钥解密是不行的。

## go的方案

## 把密文直接以16进制的形式输出，并拷贝到go源码中
xxd output.data
## 在go源码中按hex方式解码，拿到密文内容
## 然后把hash改成sha1.New()，label为空，就能解码成功。
