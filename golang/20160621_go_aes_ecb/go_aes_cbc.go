package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

func main() {
	b := []byte("test")
	key := []byte("1234567890123456")

	e, _ := AesEncryptCBC(b, key)
	fmt.Printf("%v\n", e)
	fmt.Printf("%v\n", hex.EncodeToString(e))

	d, _ := AesDecryptCBC(e, key)
	fmt.Printf("%v\n", d)
	fmt.Printf("%v\n", string(d))
}

func AesEncryptCBC(plantText []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key) //选择加密算法
	if err != nil {
		return nil, err
	}
	fmt.Printf("blocksize: %v\n", block.BlockSize())
	plantText = pKCS7Padding(plantText, block.BlockSize())
	blockModel := cipher.NewCBCEncrypter(block, key)
	ciphertext := make([]byte, len(plantText))
	blockModel.CryptBlocks(ciphertext, plantText)
	return ciphertext, nil
}

func AesDecryptCBC(ciphertext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key) //选择加密算法
	if err != nil {
		return nil, err
	}
	blockModel := cipher.NewCBCDecrypter(block, key)
	plantText := make([]byte, len(ciphertext))
	blockModel.CryptBlocks(plantText, ciphertext)
	plantText = pKCS7Unpadding(plantText)
	return plantText, nil
}

func pKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pKCS7Unpadding(plantText []byte) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}

// 工具参考：http://www.seacha.com/tools/aes.html?src=test&mode=CBC&keylen=128&key=1234567890123456&iv=&bpkcs=pkcs5padding&session=S8Y7QED1IGcNWU9kj3tt&aes=4e7da1d0d18bcbe8e0abac3209a042d1&encoding=hex&type=0
// 输出：
// blocksize: 16
// [90 70 142 134 164 138 211 118 117 165 43 143 57 215 52 223]
// 5a468e86a48ad37675a52b8f39d734df
// [116 101 115 116]
// test
