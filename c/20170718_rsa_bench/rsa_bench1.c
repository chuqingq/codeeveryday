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

RSA* rsa_load_privatekey(const char *privatekey) {
    BIO* bp = BIO_new_mem_buf(privatekey, -1);
    if (bp == NULL) {
        printf("BIO_new_mem_buf error\n");
        return NULL;
    }

    RSA *rsa = PEM_read_bio_RSAPrivateKey(bp, NULL, NULL, NULL);
    if (rsa == NULL) {
        printf("PEM_read_bio_RSAPrivateKey error\n");
        return NULL;
    }

    return rsa;
}

int main() {
    // 私钥
    char privatekey[2048];
    int privatekey_len = 2048;
    read_file("./private.pem", privatekey, &privatekey_len);
    printf("privatekey len: %d\n", privatekey_len);

    // 密文
    char encrypted[1024];
    int encrypted_len = 1024;
    read_file("./encrypted", encrypted, &encrypted_len);
    printf("encrypted len: %d\n", encrypted_len);

    int count = 1000;
    char decrypted[256];
    for (int i = 0; i < count; i++) {
        // rsa
        RSA* rsa = rsa_load_privatekey(privatekey);

        // decrypt
        int ret = RSA_private_decrypt(encrypted_len, encrypted, decrypted, rsa, /*RSA_PKCS1_PADDING*/RSA_PKCS1_OAEP_PADDING);
        if (ret <= 0) {
            printf("RSA_private_decrypt error\n");
            return -1;
        }
        RSA_free(rsa);
    }
}

