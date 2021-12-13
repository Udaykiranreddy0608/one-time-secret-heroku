package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"log"
)

func main() {
	key := []byte("dddf0fa0-59eb-11ec-bf63-0242ac13") // 32 bytes
	// dddf0fa0-59eb-11ec-bf63-0242ac13
	secret := "adsasd"
	if secret == "" {
		panic("Secret can't be null")
	}
	plaintext := []byte(secret)
	fmt.Printf("%s\n", plaintext)
	ciphertext, err := encrypt(key, plaintext)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%0x\n", ciphertext)
	result, err := decrypt(key, ciphertext)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", result)
}

// See alternate IV creation from ciphertext below
//var iv = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func encrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	ciphertext := make([]byte, aes.BlockSize+len(text))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], text)
	return ciphertext, nil
}

func decrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(text) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	return text, nil
}
