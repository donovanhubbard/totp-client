package hmac

import (
	"encoding/hex"
	"testing"
)

func TestHmacSHA1(t *testing.T) {
	key := []byte("key")
	message := []byte("The quick brown fox jumps over the lazy dog")

	hmac := ComputeHmacSha1(key, message)
	encodedString := hex.EncodeToString(hmac)

	if encodedString != "de7c9b85b8b78aa6bc8a7a36f70a90701c9db4d9" {
		t.Fatal("sha1 hmac not correct: ", encodedString)
	}
}

func TestHmacSHA256(t *testing.T) {
	key := []byte("key")
	message := []byte("The quick brown fox jumps over the lazy dog")

	hmac := ComputeHmacSha256(key, message)
	encodedString := hex.EncodeToString(hmac)

	if encodedString != "f7bc83f430538424b13298e6aa6fb143ef4d59a14946175997479dbc2d1a3cd8" {
		t.Fatal("sha256 hmac not correct: ", encodedString)
	}
}

func TestHmacSHA512(t *testing.T) {
	key := []byte("key")
	message := []byte("The quick brown fox jumps over the lazy dog")

	hmac := ComputeHmacSha512(key, message)
	encodedString := hex.EncodeToString(hmac)

	if encodedString != "b42af09057bac1e2d41708e48a902e09b5ff7f12ab428a4fe86653c73dd248fb82f948a549f7b791a5b41915ee4d1ec3935357e4e2317250d0372afa2ebeeb3a" {
		t.Fatal("sha256 hmac not correct: ", encodedString)
	}
}
