package xxtea_test

import (
	"log"

	"github.com/feiyuw/xxtea"
)

func ExampleEncrypt() {
	data := []byte("abcdefghijklmn")
	key := []byte("1231241adfsdfh3456sadfasdf")
	encryptedData, err := xxtea.Encrypt(data, key)
	if err != nil {
		log.Printf("failed to encrypt data, err: %v", err)
	}

	log.Println(encryptedData)
}

func ExampleDecrypt() {
	encryptedData := []byte{168, 165, 198, 9, 147, 202, 117, 61, 22, 111, 29, 51, 20, 115, 251, 206, 222, 124, 171, 10, 45, 171, 52, 227, 131, 143, 52, 107, 93, 188, 87, 198, 113, 123, 122, 86, 63, 114, 109, 210, 139, 140, 150, 11}
	key := []byte("1231241adfsdfh3456sadfasdf")
	data, err := xxtea.Decrypt(encryptedData, key)
	if err != nil {
		log.Printf("failed to decrypt data, err: %v", err)
	}

	log.Println(data)
}
