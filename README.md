Another xxtea encrypt/decrypt library for golang, compared to [xxtea/xxtea-go](https://github.com/xxtea/xxtea-go), it has better performance and lower memory usage.

[xxtea](https://en.wikipedia.org/wiki/XXTEA) is a fast and secure encryption algorithm. This package is its golang implementation.

## Installation

```sh
go get github.com/feiyuw/xxtea
```

## Usage

```go
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
```

## Test and Benchmark

* Test

```sh
➜  xxtea git:(master) ✗ go test -v -cover .
=== RUN   TestXXTeaEncryptAndDecrypt

  test asUint32Array ✔✔


2 total assertions


  test asKey ✔✔


4 total assertions


  asByteArray without length ✔✔


6 total assertions


  encrypt data ✔✔


8 total assertions


  decrypt data ✔✔


10 total assertions

--- PASS: TestXXTeaEncryptAndDecrypt (0.00s)
PASS
coverage: 93.8% of statements
ok  	github.com/feiyuw/xxtea	0.006s	coverage: 93.8% of statements

```

* Benchmark

```sh
➜  xxtea git:(master) ✗ go test -benchmem -bench=. ./... -run=none
goos: darwin
goarch: amd64
pkg: github.com/feiyuw/xxtea
BenchmarkEncrypt-8               3000000               531 ns/op             112 B/op          3 allocs/op
BenchmarkXXTeaGoEncrypt-8        2000000               650 ns/op             128 B/op          3 allocs/op
BenchmarkDecrypt-8               3000000               545 ns/op             112 B/op          3 allocs/op
BenchmarkXXTeaGoDecrypt-8        3000000               581 ns/op             128 B/op          3 allocs/op
PASS
ok  	github.com/feiyuw/xxtea	8.806s
```
