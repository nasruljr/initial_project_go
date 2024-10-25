package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

var iv = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Encrypt(plain string, secret string) (string, error) {
	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return "", err
	}

	plainText := []byte(plain)
	cfb := cipher.NewCFBEncrypter(block, iv)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return Encode(cipherText), nil
}

func Decode(s string) ([]byte, any, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return data, "Decode Data Failed", err
	}
	return data, "Success Decoded", nil
}

func Decrypt(chiper string, secret string) (string, error) {
	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return "", err
	}
	cipherText, _, err := Decode(chiper)
	if err != nil {
		return "", err
	}
	cfb := cipher.NewCFBDecrypter(block, iv)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}

func HmacSha256(plain string, key string) (string, error) {
	h := hmac.New(sha256.New, []byte(key))
	_, err := h.Write([]byte(plain))
	if err != nil {
		return "", err
	}
	encrypted := hex.EncodeToString(h.Sum(nil))
	return encrypted, nil
}
