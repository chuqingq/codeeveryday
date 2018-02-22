#include <openssl/ssl.h>
#include <openssl/rsa.h>

#include <stdio.h>

int read_file(const char *path, char *buf, int *len) {
    FILE *file = fopen(path, "r");
    if (file == NULL) {
        printf("read_file error\n");
        return -1;
    }

    int ret = fread(buf, 1, *len, file);
    fclose(file);
    if (ret <= 0) {
        printf("fread error\n");
        return -1;
    }

    *len = ret;
    return ret;
}

RSA* rsa_load_privatekey(const char *path) {
    BIO* bp = BIO_new( BIO_s_file() );
    BIO_read_filename( bp, path );

    RSA *rsa = PEM_read_bio_RSAPrivateKey(bp, NULL, NULL, NULL);
    BIO_free(bp);
    if (rsa == NULL) {
        printf("PEM_read_bio_RSAPrivateKey error\n");
        return NULL;
    }

    return rsa;
}

static long long ustime(void) {
    struct timeval tv;
    // long long ust;

    gettimeofday(&tv, NULL);
    // ust = ((long)tv.tv_sec)*1000000;
    // ust += tv.tv_usec;
    // return ust;
    return ((long)tv.tv_sec)*1000000 + tv.tv_usec;
}

int main() {
    // 私钥可以预先转成rsa
    RSA* rsa = rsa_load_privatekey("./private.pem");
    // printf("sizeof(RSA): %d\n", sizeof(RSA)); // sizeof(RSA): 168。但是没什么用的，里面有BIGNUM*

    int rsa_size = RSA_size(rsa);
    printf("rsa_size: %d\n", rsa_size);

    // 从文件中读取密文
    char encrypted[1024];
    int len = 1024;
    read_file("./encrypted", encrypted, &len);
    printf("encrypted len: %d\n", len);

    int count = 1000;
    char decrypted[rsa_size];
    long long start = ustime();
    for (int i = 0; i < count; i++) {
        int ret = RSA_private_decrypt(len, encrypted, decrypted, rsa, /*RSA_PKCS1_PADDING*/RSA_PKCS1_OAEP_PADDING);
        if (ret <= 0) {
            printf("RSA_private_decrypt error\n");
            return -1;
        }
        if (ret != 32) { // 明文长度是32
            printf("decrypt len: %d\n", ret);
        }
    }
    long long stop = ustime();
    printf("cout: %d, elapsed: %lld, avg: %lld us/op\n", count, stop-start, (stop-start)/count);
    RSA_free(rsa);
    
}

