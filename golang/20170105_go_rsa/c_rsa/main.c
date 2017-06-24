#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <stdint.h>

#include <openssl/ssl.h>
#include <openssl/rsa.h>
#include <openssl/aes.h>

#define RSA_KEY_TYPE_PUBLIC   1
#define RSA_KEY_TYPE_PRIVATE  0

static RSA * create_rsa(uint8_t *key, int key_type)
{
    RSA *rsa = NULL;
    BIO *keybio ;
    keybio = BIO_new_mem_buf(key, -1);
    if (keybio == NULL)
    {
        printf("Failed to create key BIO.\n");
        return NULL;
    }

    if (key_type == RSA_KEY_TYPE_PUBLIC)
    {
        rsa = PEM_read_bio_RSA_PUBKEY(keybio, &rsa, NULL, NULL);
    }
    else
    {
        rsa = PEM_read_bio_RSAPrivateKey(keybio, &rsa, NULL, NULL);
    }

    return rsa;
}

static int rsa_private_decrypt_raw(const uint8_t *rsa_input, const int input_len, uint8_t *dec_out, uint8_t *private_key)
{
    RSA * rsa = create_rsa(private_key, RSA_KEY_TYPE_PRIVATE);
    if (rsa == NULL)
    {
        printf("Failed to create RSA.\n");
        return 0;
    }

    if (input_len != RSA_size(rsa))
    {
        printf("length of data to decrypt is not equal to %d.\n", RSA_size(rsa));
        return 0;
    }

    int ret = RSA_private_decrypt(input_len, rsa_input, dec_out, rsa, RSA_PKCS1_OAEP_PADDING);
    if (ret <= 0)
    {
        printf("Failed to decrypt with RSA private key.\n");
        return 0;
    }
    return ret;
}

int main() {
    // 先获取private.key
    int privatekeyfd = open("../private.key", O_RDONLY);
    if (privatekeyfd < 0) {
        printf("open priavte.key error\n");
        return -1;
    }

    uint8_t buf[1024];
    int n = read(privatekeyfd, buf, sizeof(buf));
    if (n < 0) {
        printf("read error\n");
        return -1;
    }
    close(privatekeyfd);
    printf("read private.key len: %d\n", n);

    // 再创建RSA
    RSA* rsa = create_rsa(buf, RSA_KEY_TYPE_PRIVATE);
    if (rsa == NULL) {
        printf("create_rsa error\n");
        return -1;
    }

    // 读取密文数据
    int encryptfd = open("../output.data", O_RDONLY);
    if (encryptfd < 0) {
        printf("open output.data error\n");
        return -1;
    }

    char encrypt[128+1];
    int encrypt_len = read(encryptfd, encrypt, sizeof(encrypt));
    if (n < 0) {
        printf("read encrypt error\n");
        return -1;
    }
    
    // 解码
    int res = RSA_private_decrypt(encrypt_len, encrypt, buf, rsa, RSA_PKCS1_OAEP_PADDING);
    if (res < 0) {
        printf("RSA_private_decrypt error\n");
        return -1;
    }
    buf[res] = 0;
    printf("decrypt: %s\n", buf);

    return 0;
}
