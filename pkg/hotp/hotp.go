package hotp

import (
	"encoding/base32"
	"encoding/binary"
	"errors"
	"math"
	"math/big"

	"github.com/fuele/totp-client/pkg/hmac"
	"github.com/fuele/totp-client/pkg/slices"
)

type Hotp struct {
	Digits    int
	Algorithm string
}

func computeDynamicTruncation(hash []byte, digits int) []int {
	truncatedBytes := computeTruncatedBytes(hash)
	return computeReduction(truncatedBytes, digits)
}

func computeTruncatedBytes(hash []byte) []byte {
	lastByte := hash[len(hash)-1]
	offset := lastByte & 0xf
	offsetIndex := int(big.NewInt(0).SetBytes([]byte{offset}).Uint64())
	bits32 := hash[offsetIndex : offsetIndex+4]
	bits31 := make([]byte,4)
	bits31[0] = bits32[0] & 0x7F
	bits31[1] = bits32[1]
	bits31[2] = bits32[2]
	bits31[3] = bits32[3]
	return bits31
}

func computeReduction(sbytes []byte, digits int) []int {
	decimal := int(big.NewInt(0).SetBytes(sbytes).Uint64())
	modulo := math.Pow(10, float64(digits))
	num := decimal % int(modulo)
	reduction := slices.IntToSlicePadded(num, digits)
	return reduction
}

func (h Hotp) ComputeHotp(key string, counter int) ([]int, error) {
	keyData, err := base32.StdEncoding.DecodeString(key)

	if err != nil {
		return []int{}, errors.New("failed to base32 decode key. " + err.Error())
	}

	counterBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(counterBytes, uint64(counter))
	var mac []byte
	if h.Algorithm == "sha1" {
		mac = hmac.ComputeHmacSha1(keyData, counterBytes)
	} else if h.Algorithm == "sha256" {
		mac = hmac.ComputeHmacSha256(keyData, counterBytes)
	} else if h.Algorithm == "sha512" {
		mac = hmac.ComputeHmacSha512(keyData, counterBytes)
	} else {
		return []int{}, errors.New("unsupported Algorithm type: " + h.Algorithm)
	}

	if h.Digits < 6 {
		return []int{}, errors.New("unsupported number of digits. must be at least 6")
	}

	return computeDynamicTruncation(mac, h.Digits), nil
}
