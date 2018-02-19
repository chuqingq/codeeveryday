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
elapsed: 3309904 us, count: 100000

real   	0m3.314s
user   	0m3.308s
sys    	0m0.000s
```
这个性能数据是在virtualbox上验证的结果。

# channel方案

channel的算法可以改成：

1、c->s: clientPublicKey+clientRandom 294+16
2、c<-s: 重定向。RSASign(serverPrivateKey, clientRandom)+RSAPublicKeyEncrypt(clientPublicKey, channels) 256+256
2、c<-s: RSASign(serverPrivateKey, clientRandom)+RSAPublicKeyEncrypt(clientPublicKey, AESKey) 客户端验证服务器签名，拿到AesKey
3、c->s: RSASign(clientPrivateKey, AESKey) 服务端校验客户端