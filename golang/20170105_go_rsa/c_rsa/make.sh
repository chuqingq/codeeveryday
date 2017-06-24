gcc main.o -L./deps/openssl/lib -lssl -lcrypto -lz -lm -ldl -o main
gcc -I./deps/openssl/include -L./deps/openssl/lib -lssl -lcrypto -lz -lm -c main.c