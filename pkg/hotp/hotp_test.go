package hotp

import (
	"bytes"
	"testing"
	"encoding/hex"
	"github.com/fuele/totp-client/pkg/slices"
)

func TestComputeTruncatedBytes0(t *testing.T) {
	input := []byte{0xcc, 0x93, 0xcf, 0x18, 0x50, 0x8d, 0x94, 0x93, 0x4c, 0x64, 0xb6, 0x5d, 0x8b, 0xa7, 0x66, 0x7f, 0xb7, 0xcd, 0xe4, 0xb0}

	result := computeTruncatedBytes(input)

	if !bytes.Equal(result, []byte{0x4c, 0x93, 0xcf, 0x18}) {
		t.Fatal("Truncated Bytes are incorrect: ", hex.EncodeToString(result))
	}
}

func TestComputeTruncatedBytes1(t *testing.T) {
	input := []byte{0x75, 0xa4, 0x8a, 0x19, 0xd4, 0xcb, 0xe1, 0x00, 0x64, 0x4e, 0x8a, 0xc1, 0x39, 0x7e, 0xea, 0x74, 0x7a, 0x2d, 0x33, 0xab}

	result := computeTruncatedBytes(input)

	if !bytes.Equal(result, []byte{0x41, 0x39, 0x7e, 0xea}) {
		t.Fatal("Truncated Bytes are incorrect: ", hex.EncodeToString(result))
	}
}

func TestComputeTruncatedBytes2(t *testing.T) {
	input := []byte{0x0b, 0xac, 0xb7, 0xfa, 0x08, 0x2f, 0xef, 0x30, 0x78, 0x22, 0x11, 0x93, 0x8b, 0xc1, 0xc5, 0xe7, 0x04, 0x16, 0xff, 0x44}

	result := computeTruncatedBytes(input)

	if !bytes.Equal(result, []byte{0x8, 0x2f, 0xef, 0x30}) {
		t.Fatal("Truncated Bytes are incorrect: ", hex.EncodeToString(result))
	}
}

func TestComputeTruncatedBytes3(t *testing.T) {
	input := []byte{0x66, 0xc2, 0x82, 0x27, 0xd0, 0x3a, 0x2d, 0x55, 0x29, 0x26, 0x2f, 0xf0, 0x16, 0xa1, 0xe6, 0xef, 0x76, 0x55, 0x7e, 0xce}

	result := computeTruncatedBytes(input)

	if !bytes.Equal(result, []byte{0x66, 0xef, 0x76, 0x55}) {
		t.Fatal("Truncated Bytes are incorrect: ", hex.EncodeToString(result))
	}
}

func TestComputeTruncatedBytes4(t *testing.T) {
	input := []byte{0xa9, 0x04, 0xc9, 0x00, 0xa6, 0x4b, 0x35, 0x90, 0x98, 0x74, 0xb3, 0x3e, 0x61, 0xc5, 0x93, 0x8a, 0x8e, 0x15, 0xed, 0x1c}

	result := computeTruncatedBytes(input)

	if !bytes.Equal(result, []byte{0x61, 0xc5, 0x93, 0x8a}) {
		t.Fatal("Truncated Bytes are incorrect: ", hex.EncodeToString(result))
	}
}

func TestComputeReduction0(t *testing.T) {
	input := []byte{0x4c, 0x93, 0xcf, 0x18}
	result := computeReduction(input, 6)

	if !slices.SlicesAreEqual(result, []int{7,5,5,2,2,4}) {
		t.Fatal("incorrect Reduction: ", result)
	}
}
func TestComputeReduction1(t *testing.T) {
	input := []byte{0x41, 0x39, 0x7e, 0xea}
	result := computeReduction(input, 6)

	if !slices.SlicesAreEqual(result, []int{2,8,7,0,8,2}) {
		t.Fatal("incorrect Reduction: ", result)
	}
}

func TestComputeReduction2(t *testing.T) {
	input := []byte{0x8, 0x2f, 0xef, 0x30}
	result := computeReduction(input, 6)

	if !slices.SlicesAreEqual(result, []int{3,5,9,1,5,2}) {
		t.Fatal("incorrect Reduction: ", result)
	}
}

func TestComputeReduction3(t *testing.T) {
	input := []byte{0x66, 0xef, 0x76, 0x55}
	result := computeReduction(input, 6)

	if !slices.SlicesAreEqual(result, []int{9,6,9,4,2,9}) {
		t.Fatal("incorrect Reduction: ", result)
	}
}

func TestComputeReduction4(t *testing.T) {
	input := []byte{0x61, 0xc5, 0x93, 0x8a}
	result := computeReduction(input, 6)

	if !slices.SlicesAreEqual(result, []int{3,3,8,3,1,4}) {
		t.Fatal("incorrect Reduction: ", result)
	}
}

func TestComputeHotp0(t *testing.T) {
	hotp := Hotp{Algorithm: "sha1", Digits: 6}

	key := "GEZDGNBVGY3TQOJQGEZDGNBVGY3TQOJQ"
	counter := 0

	result, _ := hotp.ComputeHotp(key, counter)

	if !slices.SlicesAreEqual(result, []int{7,5,5,2,2,4}) {
		t.Fatal("Incorrect htop: ", result)
	}
}

func TestComputeHotp1(t *testing.T) {
	hotp := Hotp{Algorithm: "sha1", Digits: 6}

	key := "GEZDGNBVGY3TQOJQGEZDGNBVGY3TQOJQ"
	counter := 1

	result, _ := hotp.ComputeHotp(key, counter)

	if !slices.SlicesAreEqual(result, []int{2,8,7,0,8,2}) {
		t.Fatal("Incorrect htop: ", result)
	}
}

func TestComputeHotp2(t *testing.T) {
	hotp := Hotp{Algorithm: "sha1", Digits: 6}

	key := "GEZDGNBVGY3TQOJQGEZDGNBVGY3TQOJQ"
	counter := 2

	result, _ := hotp.ComputeHotp(key, counter)

	if !slices.SlicesAreEqual(result, []int{3,5,9,1,5,2}) {
		t.Fatal("Incorrect htop: ", result)
	}
}

func TestComputeHotp3(t *testing.T) {
	hotp := Hotp{Algorithm: "sha1", Digits: 6}

	key := "GEZDGNBVGY3TQOJQGEZDGNBVGY3TQOJQ"
	counter := 3

	result, _ := hotp.ComputeHotp(key, counter)

	if !slices.SlicesAreEqual(result, []int{9,6,9,4,2,9}) {
		t.Fatal("Incorrect htop: ", result)
	}
}

func TestComputeHotp4(t *testing.T) {
	hotp := Hotp{Algorithm: "sha1", Digits: 6}

	key := "GEZDGNBVGY3TQOJQGEZDGNBVGY3TQOJQ"
	counter := 4

	result, _ := hotp.ComputeHotp(key, counter)

	if !slices.SlicesAreEqual(result, []int{3,3,8,3,1,4}) {
		t.Fatal("Incorrect htop: ", result)
	}
}
