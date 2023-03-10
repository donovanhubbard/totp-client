package totp

import (
	"github.com/fuele/totp-client/pkg/hotp"
)

type Totp struct {
	Digits    int
	Algorithm string
	TimeZero  int64
	TimeStep  int64
	clock     Clock
}

func (t Totp) ComputeTotp(key string) ([]int, error) {
	timeSteps := t.computeTimeSteps()
	h := hotp.Hotp{
		Digits:    t.Digits,
		Algorithm: t.Algorithm,
	}
	return h.ComputeHotp(key, int(timeSteps))
}

func (t Totp) computeTimeSteps() int64 {
	if t.clock == nil {
		t.clock = realClock{}
	}
	currentTime := t.clock.Now().Unix()

	return (currentTime - t.TimeZero) / t.TimeStep
}
