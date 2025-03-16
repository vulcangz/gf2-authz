package ctime

import (
	"time"
)

type StaticClock struct{}

var (
	initialTime = time.Date(2100, time.January, 1, 1, 0, 0, 0, time.UTC)
)

func (c *StaticClock) Now() time.Time {
	time.Local = time.UTC

	return initialTime
}

func NewStaticClock() *StaticClock {
	return &StaticClock{}
}
