package hotp

import (
	"math"
	"math/big"
	"github.com/fuele/totp-client/pkg/hmac"
)

func computeDynamicTruncation(hash []byte, digits int) int {
	truncatedBytes := computeTruncatedBytes(hash)
	return computeReduction(truncatedBytes, digits)
}

func computeTruncatedBytes(hash []byte) []byte {
	lastByte := hash[len(hash)-1]
	offset := lastByte & 0xe
	offsetIndex := int(big.NewInt(0).SetBytes([]byte{offset}).Uint64())
	return hash[offsetIndex : offsetIndex+4]
}

func computeReduction(sbytes []byte, digits int) int {
	decimal := int(big.NewInt(0).SetBytes(sbytes).Uint64())
	modulo := math.Pow(10, float64(digits))
	return decimal % int(modulo)
}

// TODO: Write up a config object for computing HOTP that includes the
//			hash algorithm and the number of digits to return
func ComputeHotp(key, counter []byte) int{
	mac:=hmac.ComputeHmacSha1(key,counter)
	return computeDynamicTruncation(mac,6)
}
