openssl genrsa -out private.pem 2048
openssl rsa -in private.pem -pubout -out public.pem

openssl rsautl -encrypt -in data.txt -inkey public.pem -pubin -out encrypted -oaep
openssl rsautl -decrypt -inkey private.pem -in encrypted

gcc -o rsa_bench0 rsa_bench0.c -I. -O2 -Lopenssl/lib -lcrypto -lssl -pthread -ldl

