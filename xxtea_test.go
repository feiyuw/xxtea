package xxtea

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestXXTeaEncryptAndDecrypt(t *testing.T) {
	data := []byte("gvhaerutq vnp3h 7-q324bv571 5adhfadddf")
	key := []byte("1231241adfsdfh3456sadfasdf")

	encryptedData := []byte{256 - 88, 256 - 91, 256 - 58, 9, 256 - 109, 256 - 54, 117, 61, 22, 111, 29, 51, 20, 115, 256 - 5, 256 - 50, 256 - 34, 124, 256 - 85, 10, 45, 256 - 85, 52, 256 - 29, 256 - 125, 256 - 113, 52, 107, 93, 256 - 68, 87, 256 - 58, 113, 123, 122, 86, 63, 114, 109, 256 - 46, 256 - 117, 256 - 116, 256 - 106, 11}

	Convey("test asUint32Array", t, func() {
		So(asUint32Array(data, false), ShouldResemble, []uint32{1634236007, 1953854053, 1853235313, 543699824, 863055159, 1986147378, 540096309, 1751408949, 1684300134, 26212})
		So(asUint32Array(data, true), ShouldResemble, []uint32{1634236007, 1953854053, 1853235313, 543699824, 863055159, 1986147378, 540096309, 1751408949, 1684300134, 26212, uint32(len(data))})
	})

	Convey("test asKey", t, func() {
		So(asKey(key), ShouldResemble, []uint32{825438769, 1630614578, 1685284452, 875784294})
		So(asKey([]byte("12312")), ShouldResemble, []uint32{825438769, 50, 0, 0})
	})

	Convey("asByteArray without length", t, func() {
		data, err := asByteArray([]uint32{1634236007}, false)
		So(err, ShouldBeNil)
		So(data, ShouldResemble, []byte{103, 118, 104, 97})
	})

	Convey("encrypt data", t, func() {
		result, err := Encrypt(data, key)
		So(err, ShouldBeNil)
		So(result, ShouldResemble, encryptedData)
	})

	Convey("decrypt data", t, func() {
		result, err := Decrypt(encryptedData, key)
		So(err, ShouldBeNil)
		So(result, ShouldResemble, data)
	})
}

func BenchmarkEncrypt(b *testing.B) {
	data := []byte("gvhaerutq vnp3h 7-q324bv571 5adhfadddf")
	key := []byte("1231241adfsdfh3456sadfasdf")

	for i := 0; i < b.N; i++ {
		Encrypt(data, key)
	}
}

func BenchmarkDecrypt(b *testing.B) {
	encryptedData := []byte{256 - 88, 256 - 91, 256 - 58, 9, 256 - 109, 256 - 54, 117, 61, 22, 111, 29, 51, 20, 115, 256 - 5, 256 - 50, 256 - 34, 124, 256 - 85, 10, 45, 256 - 85, 52, 256 - 29, 256 - 125, 256 - 113, 52, 107, 93, 256 - 68, 87, 256 - 58, 113, 123, 122, 86, 63, 114, 109, 256 - 46, 256 - 117, 256 - 116, 256 - 106, 11}
	key := []byte("1231241adfsdfh3456sadfasdf")

	for i := 0; i < b.N; i++ {
		Decrypt(encryptedData, key)
	}
}
