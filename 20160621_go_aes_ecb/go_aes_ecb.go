package main

import (
	"bytes"
	"crypto/aes"
	"encoding/hex"
	// "github.com/cabins/Go_AES_ECB"
)

// 说明：http://www.07net01.com/2015/12/1020894.html
// 结果验证：http://www.seacha.com/tools/aes.html

// 加密
func Encrypt(plaintext []byte, key string) []byte {
	cipher, err := aes.NewCipher([]byte(key[:aes.BlockSize]))
	if err != nil {
		panic(err.Error())
	}

	if len(plaintext)%aes.BlockSize != 0 {
		panic("Need a multiple of the blocksize 16")
	}

	ciphertext := make([]byte, 0)
	text := make([]byte, 16)
	for len(plaintext) > 0 {
		// 每次运算一个block
		cipher.Encrypt(text, plaintext)
		plaintext = plaintext[aes.BlockSize:]
		ciphertext = append(ciphertext, text...)
	}
	return ciphertext
}

// 解密
func Decrypt(ciphertext []byte, key string) []byte {
	cipher, err := aes.NewCipher([]byte(key[:aes.BlockSize]))
	if err != nil {
		panic(err.Error())
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		panic("Need a multiple of the blocksize 16")
	}

	plaintext := make([]byte, 0)
	text := make([]byte, 16)
	for len(ciphertext) > 0 {
		cipher.Decrypt(text, ciphertext)
		ciphertext = ciphertext[aes.BlockSize:]
		plaintext = append(plaintext, text...)
	}
	return plaintext
}

// Padding补全
func PKCS7Pad(data []byte) []byte {
	padding := aes.BlockSize - len(data)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(0)}, padding)
	return append(data, padtext...)
}

func PKCS7UPad(data []byte) []byte {
	padLength := int(data[len(data)-1])
	return data[:len(data)-padLength]
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5Unpadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func main() {
	// key := "thisiskeyandlen>16"
	key := "1234567812345678"
	//加密
	// ciphertext := security.Encrypt( /*security.PKCS7Pad*/ ([]byte("testtesttesttest")), key)
	ciphertext := Encrypt(PKCS5Padding([]byte("test"), 16), key)
	print(hex.EncodeToString(ciphertext))
	//解密
	plaintext := string(PKCS5Unpadding([]byte(Decrypt(ciphertext, key))))
	print(plaintext)
}
