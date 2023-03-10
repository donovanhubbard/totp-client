package totp

import "time"

type Clock interface {
	Now() time.Time
}

type realClock struct{}

func (realClock) Now() time.Time { return time.Now() }

type testClock struct{
	currentTime time.Time
}

func (t testClock) Now() time.Time { return  t.currentTime}