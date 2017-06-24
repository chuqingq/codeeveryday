// http://blog.studygolang.com/2013/01/go%E5%8A%A0%E5%AF%86%E8%A7%A3%E5%AF%86%E4%B9%8Brsa/
// 生成私钥：openssl genrsa -out private.pem 1024
// 生成公钥：openssl rsa -in private.pem -pubout -out public.pem
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"log"
)

func main() {
	// data := []byte("1234567891123456789212345678931234567894123456789512345678911234567892123456789312345678941234567895")
	// log.Printf("data: %v\n", data)

	// encrypted, err := RsaEncrypt(data)
	// if err != nil {
	// 	log.Fatalf("RsaEncrypt error: %v\n", err)
	// }

	// log.Printf("encrypted: %v\n", encrypted)

	// decrypted, err := RsaDecrypt(encrypted)
	// if err != nil {
	// 	log.Fatalf("RsaDecrypt error: %v\n", err)
	// }

	// log.Printf("decrypted: %v\n", decrypted)
	// xxd output.data > 生成16进制的密文
	encrypted, _ := hex.DecodeString("280428a98916a2482447197de31ea8998c090a23f8f8ae016200659ceee8aa22c1a8167f9d571cf3187d6631e312cc86c77eb21764baf2c6a8b9fe4ea675c51e2bb3b1797c5acaba6a1d99fe1aa0063ffd017aaf75540b02d3d9f092e0478cd15d273ebf3f0d5df8565498cb5c4b3222dc707335ffa79630c2ceaa252c7994cb")
	log.Printf("len(encrypted): %v\n", len(encrypted))
	decrypted, err := RsaDecrypt(encrypted)
	if err != nil {
		log.Fatalf("RsaDecrypt error: %v\n", err)
	}

	log.Printf("decrypted: %v\n", string(decrypted))
}

var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDZsfv1qscqYdy4vY+P4e3cAtmvppXQcRvrF1cB4drkv0haU24Y
7m5qYtT52Kr539RdbKKdLAM6s20lWy7+5C0DgacdwYWd/7PeCELyEipZJL07Vro7
Ate8Bfjya+wltGK9+XNUIHiumUKULW4KDx21+1NLAUeJ6PeW+DAkmJWF6QIDAQAB
AoGBAJlNxenTQj6OfCl9FMR2jlMJjtMrtQT9InQEE7m3m7bLHeC+MCJOhmNVBjaM
ZpthDORdxIZ6oCuOf6Z2+Dl35lntGFh5J7S34UP2BWzF1IyyQfySCNexGNHKT1G1
XKQtHmtc2gWWthEg+S6ciIyw2IGrrP2Rke81vYHExPrexf0hAkEA9Izb0MiYsMCB
/jemLJB0Lb3Y/B8xjGjQFFBQT7bmwBVjvZWZVpnMnXi9sWGdgUpxsCuAIROXjZ40
IRZ2C9EouwJBAOPjPvV8Sgw4vaseOqlJvSq/C/pIFx6RVznDGlc8bRg7SgTPpjHG
4G+M3mVgpCX1a/EU1mB+fhiJ2LAZ/pTtY6sCQGaW9NwIWu3DRIVGCSMm0mYh/3X9
DAcwLSJoctiODQ1Fq9rreDE5QfpJnaJdJfsIJNtX1F+L3YceeBXtW0Ynz2MCQBI8
9KP274Is5FkWkUFNKnuKUK4WKOuEXEO+LpR+vIhs7k6WQ8nGDd4/mujoJBr5mkrw
DPwqA3N5TMNDQVGv8gMCQQCaKGJgWYgvo3/milFfImbp+m7/Y3vCptarldXrYQWO
AQjxwc71ZGBFDITYvdgJM1MTqc8xQek1FXn1vfpy2c6O
-----END RSA PRIVATE KEY-----
`)

var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDZsfv1qscqYdy4vY+P4e3cAtmv
ppXQcRvrF1cB4drkv0haU24Y7m5qYtT52Kr539RdbKKdLAM6s20lWy7+5C0Dgacd
wYWd/7PeCELyEipZJL07Vro7Ate8Bfjya+wltGK9+XNUIHiumUKULW4KDx21+1NL
AUeJ6PeW+DAkmJWF6QIDAQAB
-----END PUBLIC KEY-----
`)

func RsaEncrypt(origData []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	// return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
	return rsa.EncryptOAEP(sha1.New(), rand.Reader, pub, origData, []byte(""))
}

func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
	return rsa.DecryptOAEP(sha1.New(), rand.Reader, priv, ciphertext, []byte(""))
}
