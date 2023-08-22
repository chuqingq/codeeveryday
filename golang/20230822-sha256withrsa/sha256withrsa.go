package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"log"
)

const message = `jhproWG10011234512345123456`
const privatekey = `-----BEGIN PRIVATE KEY-----
MIIEvAIBADALBgkqhkiG9w0BAQEEggSoMIIEpAIBAAKCAQEAxughVLKYpW8tRHxp
xMYy0FevbrIoOd4AVmgGV6pkIi5czWJIf71LKYRVhKUcq6mNyvYK/hTn1sxoa3as
jBOD1ehs5neqGEnWToXxIjFLvAlhQ+SQuQOVGfvAwAuILL6DYl4n+QfSdg8un9qP
cGiytQkYH2XKATPUN9O4UUgO84DNvbl2Dzx5yNqT+IFUgIzHtTs+NzKT3N+quy1J
F/9OxBDmc1XkVQalBpAmlAUzwb4gd7P2gfsLjrMEFl9fOQEclUKdkcy2bFyCzbzY
ApyQzJ/qaLTnLpwDNduMgyoF3+Yr9OWy2vcPYML4CCKpSVFEkteUKBegOYF9FE7y
YNuCewIDAQABAoIBAQCbKOsPOf5PVsmWGgMb54wt76i/DiTI9z+GJ8GC0z0nWMk1
wcxSMSSXr05SmcYitrIQOBxdFYvAiFWQNtPktThrPdLteT1rkvWk7WErzg6JETwZ
jQvD92JxEWzLonNIjBjLPC2sWoi3ZaJ2OjUYd+Onyv9RRsLsazTJk9O9PBvFoDsH
Pjj5spypnt5JmsfWwHWxvq8auMI3GoI0NJL0tVur2n3GdXT37TvCQHMydklDrPrq
u7MP7pLN7x/7HnTv/J6GL0svDYGvTcLBVl3NTvimleVaRJ6MgTY+cB/+rImR0ktY
T830OLh4p2Bpt7xUMTlltGCdXrANOTi6kBmx2H4BAoGBAPzdBko9whmERgsbLILw
c8mDGQH3lDpSBWNeMo45dqeqSighNHYj48krgGHq/DSNHHAWYq5i63/EdMy0zf7B
EmA9t8x8WRdQntdJGyDhCKKddUwjOLjI34p9YBwJTXcPMhMTmzhFiQAFNzzwkTVd
nYfecXxvgUY9mDPizLH88E97AoGBAMlfw8UhFbZ9zHG1Fyidu0zAwZ7sxQc7+vyF
ptU+Ff+XvCY8wyWBdgFjxWoE+JLVMmTF2wo4Vy9jdXRvzVYzaWcTIBHBGzmSVP9N
uThSc9ZcUlfifxEFycH/YciaxLSuLqkiSv5IGaojdQ0nroGbTaz59FwqTUO0ulSd
3NOSaakBAoGBAM1JS4/+b5RztMHTf+GWAQq6ahUUwLxQVpuDoBujP1eDgsztmD/J
h2aM8J+OPM8VON8u7VKScIq8He8LYqnOaXLE6HEVCudIxowVh/a7e1055D654ZTz
T7iJbPuV+dQM/CRMqJmYqk7f7SaGT/05UWk7CHtzs0opO2X0XSarKRX9AoGAGIqe
PkEI92OfbeAnAWEvuWvobOjoHjiWHv5e1bAqWCry2CohkkmTyxmQrpoKfUKUUKm4
RyeUoIbbgqQ5fx7m4pP3HZLOMZb+2tprD00lJuO7eVB2Menlq8nm7d7GyEpOD3jJ
cPHyhsSpeD/0yYDW15TizfSt0+mLp9JRXkuCqwECgYBM7ZGXMfm8wIhSnEYPu1Xq
RZY0/AjX2hzbtFVV5Ub6BPQvVHVEtaFlcvPGCQr0zFwDA/+JeMyXwXgUl+sFWS8K
rlmt5WP2Imepza86BDlO2PMkXsxq7HR6HrtcgWiEaUNWnj9XYH6pulfWEXxeEPrL
6jDPsSv9RFJdgmzJmXZNiA==
-----END PRIVATE KEY-----`
const publicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxughVLKYpW8tRHxpxMYy
0FevbrIoOd4AVmgGV6pkIi5czWJIf71LKYRVhKUcq6mNyvYK/hTn1sxoa3asjBOD
1ehs5neqGEnWToXxIjFLvAlhQ+SQuQOVGfvAwAuILL6DYl4n+QfSdg8un9qPcGiy
tQkYH2XKATPUN9O4UUgO84DNvbl2Dzx5yNqT+IFUgIzHtTs+NzKT3N+quy1JF/9O
xBDmc1XkVQalBpAmlAUzwb4gd7P2gfsLjrMEFl9fOQEclUKdkcy2bFyCzbzYApyQ
zJ/qaLTnLpwDNduMgyoF3+Yr9OWy2vcPYML4CCKpSVFEkteUKBegOYF9FE7yYNuC
ewIDAQAB
-----END PUBLIC KEY-----`

const expect_sign = `Gz4oaXCNdSVJHrkMo6j-H-wF4D2AV6kpXh2e4HpM61yCGgrYLxwiIF8nvVJBETgGSSWqgwqkSJEy15Z4iR4LWCQaKQQAaIdNJYSLH1ewo1InIAZWTVqOJ83sxBsimnimDzA3VjzXQB4eIUa0LVyBYGLhJP62a4pYTxnET6zZKrDCAoDDb6QKNyn1iZx2W61-yui8QrnD4wiPLQDyQ37JRMLAH_-39mcMkngOP2rbl9AK5eTKqhn1ke9XtQ8huNUB8Hk6casw3CIG2ED4Tv_RXehPObRZystJ7r2fE4xvIku-X24YrfPPcmKB7x4EWMlbdDEijuUZBDQdmrddsMluvw`

func main() {
	sign := SHA256WithRSASign(privatekey, message)
	if sign != expect_sign {
		log.Printf("sign[%v] != expect_sign[%v]", sign, expect_sign)
	} else {
		log.Printf("sign success")
	}

	err := SHA256WithRSAVerify(publicKey, message, expect_sign)
	if err != nil {
		log.Printf("RsaVerify error: %v", err)
	} else {
		log.Printf("verify success")
	}
}

func SHA256WithRSASign(privateKey, signContent string) string {
	hashed := sha256.Sum256([]byte(message))

	priKey, err := parsePrivateKey(privateKey)
	if err != nil {
		panic(err)
	}

	signature, err := rsa.SignPKCS1v15(rand.Reader, priKey, crypto.SHA256, hashed[:])
	if err != nil {
		panic(err)
	}
	return base64.RawURLEncoding.EncodeToString(signature)
}

func SHA256WithRSAVerify(publicKey string, message string, sig string) error {
	hashed := sha256.Sum256([]byte(message))

	pubKey, err := parsePublicKey(publicKey)
	if err != nil {
		log.Printf("parse publickey error: %v", err)
		return err
	}

	sig1, err := base64.RawURLEncoding.DecodeString(sig)
	if err != nil {
		log.Printf("base64 decode signature error: %v", err)
		return err
	}

	return rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hashed[:], sig1)
}

func parsePublicKey(publicKey string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return nil, errors.New("pem.Decode publickey error")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pubKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("invalid rsa publickey")
	}
	return pubKey, nil
}

// func formatPublicKey(publicKey string) string {
// 	const (
// 		PUB_PEM_BEGIN = "-----BEGIN CERTIFICATE-----\n"
// 		PUB_PEM_END   = "\n-----END CERTIFICATE-----"
// 	)
// 	if !strings.HasPrefix(publicKey, PUB_PEM_BEGIN) {
// 		publicKey = PUB_PEM_BEGIN + publicKey
// 	}
// 	if !strings.HasSuffix(publicKey, PUB_PEM_END) {
// 		publicKey = publicKey + PUB_PEM_END
// 	}
// 	return publicKey
// }

func parsePrivateKey(privateKey string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return nil, errors.New("pem.Decode error")
	}

	pri, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	priKey, ok := pri.(*rsa.PrivateKey)
	if !ok {
		return priKey, errors.New("invalid rsa privatekey")
	}
	return priKey, nil
}

// func formatPrivateKey(privateKey string) string {
// 	const (
// 		PEM_BEGIN = "-----BEGIN RSA PRIVATE KEY-----\n"
// 		PEM_END   = "\n-----END RSA PRIVATE KEY-----"
// 	)
// 	if !strings.HasPrefix(privateKey, PEM_BEGIN) {
// 		privateKey = PEM_BEGIN + privateKey
// 	}
// 	if !strings.HasSuffix(privateKey, PEM_END) {
// 		privateKey = privateKey + PEM_END
// 	}
// 	return privateKey
// }
