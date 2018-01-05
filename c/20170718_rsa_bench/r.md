# openssl generate private and public key

```
openssl genrsa -out private.pem 2048
openssl rsa -in private.pem -pubout -out public.pem
```

# encrypt data whith openssl

```
openssl rsautl -encrypt -in data.txt -inkey public.pem -pubin -out encrypted -oaep
openssl rsautl -decrypt -inkey private.pem -in encrypted
```

# code compile

```
gcc -o rsa_bench0 rsa_bench0.c -I. -O2 -Lopenssl/lib -lcrypto -lssl -pthread -ldl
```

# benchmark result

```
$ time ./rsa_bench0 
encrypted len: 256

real	0m1.025s
user	0m1.028s
sys	0m0.000s

$ time ./rsa_bench1
privatekey len: 1676
encrypted len: 256

real	0m1.503s
user	0m1.500s
sys	0m0.000s

$ time ./rsa_bench2

real	0m1.496s
user	0m1.488s
sys	0m0.004s
```

# verify benchmark

```
g++ -o rsa_verify rsa_verify_bench.cpp -O2 -lcrypto -lssl -pthread -ldl
$ time ./rsa_verify_bench
real   	0m0.066s
user   	0m0.060s
sys    	0m0.000s
```
这个数据还是包含了：
1、两套加解密
2、每次create_rsa

channel的算法可以改成：
1、c->s: clientPublicKey,clientRandom
2、c<-s: ServerSignature(clientRandom),RSAPublicKeyEncrypt(clientPublicKey,AesKey) 客户端验证服务器签名，拿到AesKey
3、c->s: clientSignature(aesKey) 服务端校验客户端

