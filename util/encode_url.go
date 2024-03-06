package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

var keyText = "5618172a-dbb1-11ee-aa6f-37b636981ae9"

func EncodURLParam(param string) {
	originalText := param
	fmt.Println(originalText)

	// encrypt value to base64
	cryptoText := URLEncode(originalText)
	fmt.Println(cryptoText)

	// encrypt base64 crypto to original value
	text := URLDecode(cryptoText)
	fmt.Print(text)
}

// encrypt string to base64 crypto using AES
func URLEncode(param string) string {
	key := []byte(keyText[:16])
	plaintext := []byte(param)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// convert to base64
	return base64.URLEncoding.EncodeToString(ciphertext)
}

// decrypt from base64 to decrypted string
func URLDecode(cryptoParam string) string {
	key := []byte(keyText[:16])
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoParam)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext)
}
