package xxtea_test

import (
	"log"

	"github.com/feiyuw/xxtea"
)

func ExampleEncrypt() ([]byte, error) {
	data := []byte("abcdefghijklmn")
	key := []byte("1231241adfsdfh3456sadfasdf")
	encryptedData, err := xxtea.Encrypt(data, key)
	if err != nil {
		log.Printf("failed to encrypt data, err: %v", err)
	}

	return encryptedData, err
}

func ExampleDecrypt() ([]byte, error) {
	encryptedData := []byte{256 - 88, 256 - 91, 256 - 58, 9, 256 - 109, 256 - 54, 117, 61, 22, 111, 29, 51, 20, 115, 256 - 5, 256 - 50, 256 - 34, 124, 256 - 85, 10, 45, 256 - 85, 52, 256 - 29, 256 - 125, 256 - 113, 52, 107, 93, 256 - 68, 87, 256 - 58, 113, 123, 122, 86, 63, 114, 109, 256 - 46, 256 - 117, 256 - 116, 256 - 106, 11}
	key := []byte("1231241adfsdfh3456sadfasdf")
	data, err := xxtea.Decrypt(encryptedData, key)
	if err != nil {
		log.Printf("failed to decrypt data, err: %v", err)
	}

	return data, err
}
