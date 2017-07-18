#include <openssl/ssl.h>
#include <openssl/rsa.h>

#include <stdio.h>

char encrypted[] = "";

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
    if (rsa == NULL) {
        printf("PEM_read_bio_RSAPrivateKey error\n");
        return NULL;
    }

    return rsa;
}

int main() {
  int count = 1000;
  for (int i = 0; i < count; i++) {
    // 私钥可以预先转成rsa
    RSA* rsa = rsa_load_privatekey("./private.pem");

    int rsa_size = RSA_size(rsa);
    // printf("rsa_size: %d\n", rsa_size);

    // 从文件中读取密文
    char encrypted[1024];
    int len = 1024;
    read_file("./encrypted", encrypted, &len);
    // printf("encrypted len: %d\n", len);

    // int count = 1000;
    char decrypted[rsa_size];
    // for (int i = 0; i < count; i++) {
        int ret = RSA_private_decrypt(len, encrypted, decrypted, rsa, /*RSA_PKCS1_PADDING*/RSA_PKCS1_OAEP_PADDING);
        if (ret <= 0) {
            printf("RSA_private_decrypt error\n");
            return -1;
        }
    // }

    RSA_free(rsa);
  }
}

