package util

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

// RsaEncrypt 代付到银行卡加密方式
// Wechat的公钥为：PKCS#1格式
func RsaEncrypt(publicKey string, plainText string) string {
	puk, err := parsePKCS1Key([]byte(publicKey))
	if err != nil {
		return ""
	}
	groups := oaepPadding(puk, []byte(plainText))
	buffer := bytes.Buffer{}
	for _, plainTextBlock := range groups {
		cipherText, err := rsa.EncryptOAEP(sha1.New(), rand.Reader, puk, plainTextBlock, make([]byte, 0))
		if err != nil {
			return ""
		}
		buffer.Write(cipherText)
	}
	return base64.StdEncoding.EncodeToString(buffer.Bytes())
}

// parsePKCS1Key PKCS#1公钥解析
func parsePKCS1Key(publicKey []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("publicKey is not pem formate")
	}
	pub, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return pub, nil
}

func oaepPadding(puk *rsa.PublicKey, src []byte) [][]byte {
	size := len(puk.N.Bytes()) - 41
	var groups [][]byte
	srcSize := len(src)
	if srcSize <= size {
		groups = append(groups, src)
	} else {
		for len(src) != 0 {
			if len(src) <= size {
				groups = append(groups, src)
				break
			} else {
				v := src[:size]
				groups = append(groups, v)
				src = src[size:]
			}
		}
	}
	return groups
}
