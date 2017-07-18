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

