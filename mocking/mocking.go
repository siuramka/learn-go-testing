package mocking

import (
	"fmt"
	"io"
	"time"
)

const countdownStart = 3
const finalWord = "GO!"

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (m *SpySleeper) Sleep() {
	m.Calls++
}

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

type ConfigurableSleeper struct {
	Duration  time.Duration
	SleepFunc func(time.Duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func (c *ConfigurableSleeper) Sleep() {
	c.SleepFunc(c.Duration)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		_, _ = fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	_, _ = fmt.Fprintln(out, finalWord)
}
