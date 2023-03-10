package totp

import (
	"testing"
	"time"
	"github.com/fuele/totp-client/pkg/slices"
)

func TestComputeTotp0(t *testing.T) {
	key := "GEZDGNBVGY3TQOJQGEZDGNBVGY3TQOJQ"
	c := testClock { currentTime: time.Date(1970, 1, 1, 0, 0, 59, 0, time.UTC)}
	totp := Totp{
		Algorithm: "sha1",
		Digits:    8,
		TimeZero:  0,
		TimeStep:  30,
		clock:     c,
	}

	result, err := totp.ComputeTotp(key)

	if err != nil {
		t.Fatal("error is not null: ", err)
	}

	if !slices.SlicesAreEqual(result, []int{9,4,2,8,7,0,8,2}) {
		t.Fatal("slices are not equal. Result: ", result)
	}
}

func TestComputeTotp1(t *testing.T) {
	key := "GEZDGNBVGY3TQOJQGEZDGNBVGY3TQOJQ"
	c := testClock { currentTime: time.Date(2005, 3, 18, 1, 58, 29, 0, time.UTC)}
	totp := Totp{
		Algorithm: "sha1",
		Digits:    8,
		TimeZero:  0,
		TimeStep:  30,
		clock:     c,
	}

	result, err := totp.ComputeTotp(key)

	if err != nil {
		t.Fatal("error is not null: ", err)
	}

	if !slices.SlicesAreEqual(result, []int{0,7,0,8,1,8,0,4}) {
		t.Fatal("slices are not equal. Result: ", result)
	}
}

func TestComputeTotp2(t *testing.T) {
	key := "GEZDGNBVGY3TQOJQGEZDGNBVGY3TQOJQ"
	c := testClock { currentTime: time.Date(2005, 3, 18, 1, 58, 31, 0, time.UTC)}
	totp := Totp{
		Algorithm: "sha1",
		Digits:    8,
		TimeZero:  0,
		TimeStep:  30,
		clock:     c,
	}

	result, err := totp.ComputeTotp(key)

	if err != nil {
		t.Fatal("error is not null: ", err)
	}

	if !slices.SlicesAreEqual(result, []int{1,4,0,5,0,4,7,1}) {
		t.Fatal("slices are not equal. Result: ", result)
	}
}

func TestComputeTotp3(t *testing.T) {
	key := "GEZDGNBVGY3TQOJQGEZDGNBVGY3TQOJQ"
	c := testClock { currentTime: time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC)}
	totp := Totp{
		Algorithm: "sha1",
		Digits:    8,
		TimeZero:  0,
		TimeStep:  30,
		clock:     c,
	}

	result, err := totp.ComputeTotp(key)

	if err != nil {
		t.Fatal("error is not null: ", err)
	}

	if !slices.SlicesAreEqual(result, []int{8,9,0,0,5,9,2,4}) {
		t.Fatal("slices are not equal. Result: ", result)
	}
}
