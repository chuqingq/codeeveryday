https://stackoverflow.com/questions/22344149/to-use-rsa-pkcs1-oaep-padding-for-rsa-signature

# 生成rsa私钥。我是从terminal.go中拷贝出来的
openssl genrsa -out private.key 2048

# 生成明文数据
echo "Hello world" > input.data

# 加密
openssl rsautl -encrypt -oaep -inkey private.key -in input.data -out output.data

# 解密验证
openssl rsautl -decrypt -oaep -inkey private.key -in output.data -out back.data

# 可以验证多次。每次密文不同，但还原明文都和明文是一致的。

# go的方案

# 把密文直接以16进制的形式输出，并拷贝到go源码中
xxd output.data
# 在go源码中按hex方式解码，拿到密文内容
# 然后把hash改成sha1.New()，label为空，就能解码成功。

