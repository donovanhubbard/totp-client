package hmac

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
)

func computeHash(algorithm func() hash.Hash, key []byte, message []byte) []byte {
	mac := hmac.New(algorithm, key)
	mac.Write(message)
	return mac.Sum(nil)
}

func ComputeHmacSha1(key, message []byte) []byte {
	return computeHash(sha1.New, key, message)
}

func ComputeHmacSha256(key, message []byte) []byte {
	return computeHash(sha256.New, key, message)
}

func ComputeHmacSha512(key, message []byte) []byte {
	return computeHash(sha512.New, key, message)
}
