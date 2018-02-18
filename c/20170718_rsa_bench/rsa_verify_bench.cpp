#include<openssl/pem.h>
#include<openssl/ssl.h>
#include<openssl/rsa.h>
#include<openssl/evp.h>
#include<openssl/bio.h>
#include<openssl/err.h>
#include <stdio.h> 
#include<iostream>
#include<fstream>

using namespace std;

int padding= RSA_PKCS1_PADDING;

 char publicKey[]="-----BEGIN PUBLIC KEY-----\n"\
     "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAy8Dbv8prpJ/0kKhlGeJY\n"\
     "ozo2t60EG8L0561g13R29LvMR5hyvGZlGJpmn65+A4xHXInJYiPuKzrKUnApeLZ+\n"\
     "vw1HocOAZtWK0z3r26uA8kQYOKX9Qt/DbCdvsF9wF8gRK0ptx9M6R13NvBxvVQAp\n"\
     "fc9jB9nTzphOgM4JiEYvlV8FLhg9yZovMYd6Wwf3aoXK891VQxTr/kQYoq1Yp+68\n"\
     "i6T4nNq7NWC+UNVjQHxNQMQMzU6lWCX8zyg3yH88OAQkUXIXKfQ+NkvYQ1cxaMoV\n"\
     "PpY72+eVthKzpMeyHkBn7ciumk5qgLTEJAfWZpe4f4eFZj/Rc8Y8Jj2IS5kVPjUy\n"\
     "wQIDAQAB\n"\
     "-----END PUBLIC KEY-----\n"; 
     
char privateKey[]="-----BEGIN RSA PRIVATE KEY-----\n"\
     "MIIEowIBAAKCAQEAy8Dbv8prpJ/0kKhlGeJYozo2t60EG8L0561g13R29LvMR5hy\n"\
     "vGZlGJpmn65+A4xHXInJYiPuKzrKUnApeLZ+vw1HocOAZtWK0z3r26uA8kQYOKX9\n"\
     "Qt/DbCdvsF9wF8gRK0ptx9M6R13NvBxvVQApfc9jB9nTzphOgM4JiEYvlV8FLhg9\n"\
     "yZovMYd6Wwf3aoXK891VQxTr/kQYoq1Yp+68i6T4nNq7NWC+UNVjQHxNQMQMzU6l\n"\
     "WCX8zyg3yH88OAQkUXIXKfQ+NkvYQ1cxaMoVPpY72+eVthKzpMeyHkBn7ciumk5q\n"\
     "gLTEJAfWZpe4f4eFZj/Rc8Y8Jj2IS5kVPjUywQIDAQABAoIBADhg1u1Mv1hAAlX8\n"\
     "omz1Gn2f4AAW2aos2cM5UDCNw1SYmj+9SRIkaxjRsE/C4o9sw1oxrg1/z6kajV0e\n"\
     "N/t008FdlVKHXAIYWF93JMoVvIpMmT8jft6AN/y3NMpivgt2inmmEJZYNioFJKZG\n"\
     "X+/vKYvsVISZm2fw8NfnKvAQK55yu+GRWBZGOeS9K+LbYvOwcrjKhHz66m4bedKd\n"\
     "gVAix6NE5iwmjNXktSQlJMCjbtdNXg/xo1/G4kG2p/MO1HLcKfe1N5FgBiXj3Qjl\n"\
     "vgvjJZkh1as2KTgaPOBqZaP03738VnYg23ISyvfT/teArVGtxrmFP7939EvJFKpF\n"\
     "1wTxuDkCgYEA7t0DR37zt+dEJy+5vm7zSmN97VenwQJFWMiulkHGa0yU3lLasxxu\n"\
     "m0oUtndIjenIvSx6t3Y+agK2F3EPbb0AZ5wZ1p1IXs4vktgeQwSSBdqcM8LZFDvZ\n"\
     "uPboQnJoRdIkd62XnP5ekIEIBAfOp8v2wFpSfE7nNH2u4CpAXNSF9HsCgYEA2l8D\n"\
     "JrDE5m9Kkn+J4l+AdGfeBL1igPF3DnuPoV67BpgiaAgI4h25UJzXiDKKoa706S0D\n"\
     "4XB74zOLX11MaGPMIdhlG+SgeQfNoC5lE4ZWXNyESJH1SVgRGT9nBC2vtL6bxCVV\n"\
     "WBkTeC5D6c/QXcai6yw6OYyNNdp0uznKURe1xvMCgYBVYYcEjWqMuAvyferFGV+5\n"\
     "nWqr5gM+yJMFM2bEqupD/HHSLoeiMm2O8KIKvwSeRYzNohKTdZ7FwgZYxr8fGMoG\n"\
     "PxQ1VK9DxCvZL4tRpVaU5Rmknud9hg9DQG6xIbgIDR+f79sb8QjYWmcFGc1SyWOA\n"\
     "SkjlykZ2yt4xnqi3BfiD9QKBgGqLgRYXmXp1QoVIBRaWUi55nzHg1XbkWZqPXvz1\n"\
     "I3uMLv1jLjJlHk3euKqTPmC05HoApKwSHeA0/gOBmg404xyAYJTDcCidTg6hlF96\n"\
     "ZBja3xApZuxqM62F6dV4FQqzFX0WWhWp5n301N33r0qR6FumMKJzmVJ1TA8tmzEF\n"\
     "yINRAoGBAJqioYs8rK6eXzA8ywYLjqTLu/yQSLBn/4ta36K8DyCoLNlNxSuox+A5\n"\
     "w6z2vEfRVQDq4Hm4vBzjdi3QfYLNkTiTqLcvgWZ+eX44ogXtdTDO7c+GeMKWz4XX\n"\
     "uJSUVL5+CVjKLjZEJ6Qc2WZLl94xSwL71E41H4YciVnSCQxVc4Jw\n"\
     "-----END RSA PRIVATE KEY-----\n";   
     

//把字符串写成public.pem文件
int createPublicFile(char *file,const string &pubstr)
{
    if(pubstr.empty())
    {
        printf("public key read error\n");
        return (-1);
    }
    int len = pubstr.length();
    string tmp = pubstr;
    for(int i = 64;i<len;i+=64)
    {
        if(tmp[i] != '\n')
        {
                tmp.insert(i,"\n");
        }
        i++;
    }
    tmp.insert(0, "-----BEGIN PUBLIC KEY-----\n");
    tmp.append("\n-----END PUBLIC KEY-----\n");
    
    //写文件
    ofstream fout(file);
    fout<<tmp;  
    
    return (0);
}

//把字符串写成private.pem文件
int createPrivateFile(char *file,const string &pristr)
{
    if(pristr.empty())
    {
        printf("public key read error\n");
        return (-1);
    }
    int len = pristr.length();
    string tmp = pristr;
    for(int i = 64;i<len;i+=64)
    {
        if(tmp[i] != '\n')
        {
                tmp.insert(i,"\n");
        }
        i++;
    }
    tmp.insert(0, "-----BEGIN RSA PRIVATE KEY-----\n");
    tmp.append("-----END RSA PRIVATE KEY-----\n");

    //写文件
    ofstream fout(file);
    fout<<tmp;  
    
    return (0);
}

int createPUBKEYFromRSA(RSA *rsa)
{
    FILE *file = fopen("pubkey222.crt", "w");
    // PEM_write_RSAPublicKey(file, rsa);
    PEM_write_RSAPublicKey(file, rsa);
    return 0;
}

//读取密钥
RSA* createRSA(unsigned char*key,int publi)
 {    
     RSA *rsa= NULL;    
     BIO*keybio ;    
     keybio= BIO_new_mem_buf(key, -1);   
     if(keybio==NULL) 
    {       
        printf("Failed to create key BIO\n");    
        return 0;   
    }
    
    if(publi)
    {    
    rsa = PEM_read_bio_RSA_PUBKEY(keybio, &rsa,NULL, NULL);   
    }
    else   
    {    
    rsa= PEM_read_bio_RSAPrivateKey(keybio, &rsa,NULL, NULL); 
    }           
    if(rsa== NULL)    
    {        
    printf("Failed to create RSA\n");   
    }      
    return rsa;
 } 
 
    
//公钥加密，私钥解密
int public_encrypt(unsigned char*data,int data_len,unsigned char*key, unsigned char*encrypted)
{    
    RSA* rsa = createRSA(key,1); 
    int result= RSA_public_encrypt(data_len,data,encrypted,rsa,padding);  
    return result;
}

int private_decrypt(unsigned char*enc_data,int data_len,unsigned char*key, unsigned char*decrypted)
{   
    RSA* rsa = createRSA(key,0);   
    int result= RSA_private_decrypt(data_len,enc_data,decrypted,rsa,padding);  
    return result;
}  

//私钥加密，公钥解密
int private_encrypt(unsigned char*data,int data_len,unsigned char*key, unsigned char*encrypted)
{   
    RSA* rsa = createRSA(key,0);  
    int result= RSA_private_encrypt(data_len,data,encrypted,rsa,padding);   
    return result;
}
int public_decrypt(unsigned char*enc_data,int data_len,unsigned char*key, unsigned char*decrypted)
{    
    RSA* rsa = createRSA(key,1);    
    int result= RSA_public_decrypt(data_len,enc_data,decrypted,rsa,padding);   
    return result;
} 

//私钥签名，公钥验签
int private_sign(const unsigned char *in_str,unsigned int in_str_len,unsigned char *outret,unsigned int *outlen,/*unsigned char*key*/RSA *privateRSA)
{
    // RSA* rsa = createRSA(key,0);  
    int result = RSA_sign(NID_sha1,in_str,in_str_len,outret,outlen,privateRSA);
    if(result != 1)
    {
        printf("sign error\n");
        return -1;
    }
    return result;
}

int public_verify(const unsigned char *in_str, unsigned int in_len,unsigned char *outret, unsigned int outlen,/*unsigned char*key*/RSA *publicRSA)
{
    // RSA* rsa = createRSA(key,1); 
    int result = RSA_verify(NID_sha1,in_str,in_len,outret,outlen,publicRSA);
    if(result != 1)
    {
        printf("verify error\n");
        return -1;
    }
    return result;
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
 
int main2()
{   

    char plainText[2048/8]= "Hello 12151 +++ == !@##$$%%^&&*&**()this is Ravi";//key length : 2048  
    printf("create pem file\n");
    string strPublicKey="MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQChNr0TmflORv9C62+tSAYhyj4DwB6fyOHqttddq8Y+R+8cIGT7EKuqSRuUUuLVBN6IIjd14UkxxtjHqrDxPWZz9WfX0LB2lTmnSdkg9Q10IfP9ZrVCW8Pe5vJ7gt5iQ4lOebdqR47+ef9E7oE+eJFQhxSYGGy/FnKjBkadJQtwPQIDAQAB";
    int file_ret = createPublicFile("public_test.pem",strPublicKey);

    unsigned char encrypted[4098]={};
    unsigned char decrypted[4098]={}; 
    unsigned char signret[4098]={};
    unsigned int siglen;

    printf("source data=[%s]\n",plainText);

    // printf("public encrytpt ----private decrypt \n\n");
    // int encrypted_length=public_encrypt((unsigned char*)plainText,strlen(plainText),( unsigned char*)publicKey,encrypted);
    // if(encrypted_length== -1)
    // {   
    //     printf("encrypted error \n");
    //     exit(0);
    // }
    // printf("Encrypted length =%d\n",encrypted_length); 
    // int decrypted_length= private_decrypt((unsigned char*)encrypted,encrypted_length,(unsigned char*)privateKey, decrypted);
    // if(decrypted_length== -1)
    // {  
    //     printf("decrypted error \n");
    //     exit(0);
    // }
    // printf("DecryptedText =%s\n",decrypted);
    // printf("DecryptedLength =%d\n",decrypted_length); 

    // printf("private encrytpt ----public decrypt \n\n");
    // encrypted_length=private_encrypt((unsigned char*)plainText,strlen(plainText),(unsigned char*)privateKey,encrypted);
    // if(encrypted_length== -1)
    // {    
    //     printf("encrypted error \n");
    //     exit(0);
    // }
    // printf("Encrypted length =%d\n",encrypted_length); 
    // decrypted_length= public_decrypt(encrypted,encrypted_length,(unsigned char*)publicKey, decrypted);
    // if(decrypted_length== -1)
    // {    
    //     printf("decrypted error \n");
    //     exit(0);
    // }
    // printf("DecryptedText =%s\n",decrypted);
    // printf("DecryptedLength =%d\n",decrypted_length); 
    
    printf("\nprivate sign ----public verify \n\n");
    RSA *privateRSA = createRSA((unsigned char*)privateKey, 0);
    printf("sizeof(*privateRSA): %d\n", sizeof(*privateRSA));
    printf("RSA_size(): %d\n", RSA_size(privateRSA));
    RSA *publicRSA = createRSA((unsigned char*)publicKey, 1);
    printf("sizeof(*publicRSA): %d\n", sizeof(*publicRSA));

    // 验证通过publicRSA生成publicKey，是否和原来一样
    int ret = createPUBKEYFromRSA(publicRSA);

    char *pubkey = BN_bn2hex(publicRSA->n);
    printf("pubkey: %s\n", pubkey);

    ret = private_sign((const unsigned char*)plainText,strlen(plainText),signret,&siglen,/*(unsigned char*)privateKey*/privateRSA);
    printf("sign ret =[%d]\n",ret);// siglen为256。signret如果修改，verify就报错
    ret =  public_verify((const unsigned char*)plainText, strlen(plainText),signret, siglen,/*( unsigned char*)publicKey*/publicRSA);
    printf("verify ret =[%d]\n",ret);

    const int count = 100000;
    long long start = ustime();
    // RSA_verify benchmark
    for (int i = 0; i < count; i++) {
        public_verify((const unsigned char*)plainText, strlen(plainText),signret, siglen,/*( unsigned char*)publicKey*/publicRSA);
    }
    long long stop = ustime();
    printf("elapsed: %lld us, count: %d\n", stop-start, count);
    
    return (0);
}

int main() {
    const char *MODULUS = "CBC0DBBFCA6BA49FF490A86519E258A33A36B7AD041BC2F4E7AD60D77476F4BBCC479872BC6665189A669FAE7E038C475C89C96223EE2B3ACA52702978B67EBF0D47A1C38066D58AD33DEBDBAB80F2441838A5FD42DFC36C276FB05F7017C8112B4A6DC7D33A475DCDBC1C6F5500297DCF6307D9D3CE984E80CE0988462F955F052E183DC99A2F31877A5B07F76A85CAF3DD554314EBFE4418A2AD58A7EEBC8BA4F89CDABB3560BE50D563407C4D40C40CCD4EA55825FCCF2837C87F3C38042451721729F43E364BD843573168CA153E963BDBE795B612B3A4C7B21E4067EDC8AE9A4E6A80B4C42407D66697B87F8785663FD173C63C263D884B99153E3532C1";
    // PUBLIC_EXPONENT: RSA_F4

    RSA *r = RSA_new();
    r->n = BN_new();
    BN_hex2bn(&r->n, MODULUS);
    r->e = BN_new();
    BN_set_word(r->e, RSA_F4);

    FILE *file = fopen("pubkey_2.crt", "w");
    // PEM_write_RSAPublicKey(file, rsa);
    PEM_write_RSAPublicKey(file, r);

    return 0;
}
