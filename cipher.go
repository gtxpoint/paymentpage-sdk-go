package paymentpage

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"encoding/base64"
	"fmt"
)

const keySize = 32

func Ase256(plaintext string, key string) string {
	blockSize := aes.BlockSize
	iv := make([]byte, aes.BlockSize)
	io.ReadFull(rand.Reader, iv)

	bKey := PKCS5PaddingKey([]byte(key))
	bIV := []byte(iv)
	bPlaintext := PKCS5Padding([]byte(plaintext), blockSize)

	block, _ := aes.NewCipher(bKey)
	ciphertext := make([]byte, len(bPlaintext))
	mode := cipher.NewCBCEncrypter(block, bIV)
	mode.CryptBlocks(ciphertext, bPlaintext)
	encodedString := base64.StdEncoding.EncodeToString(ciphertext) 

	toEncode := fmt.Sprintf("%s::%s", encodedString, base64.StdEncoding.EncodeToString(bIV))

	return base64.StdEncoding.EncodeToString([]byte(toEncode))
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := (blockSize - len(ciphertext)%blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5PaddingKey(ciphertext []byte) []byte {
	padding := (keySize - len(ciphertext)%keySize)
	padtext := bytes.Repeat([]byte{byte(0)}, padding)
	return append(ciphertext, padtext...)[:keySize]
}