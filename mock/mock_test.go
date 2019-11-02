package main

import (
    "bytes"
    "reflect"
    "testing"
    "time"
)

type SpySleeper struct {
    Calls int
}

func (s *SpySleeper) Sleep() {
    s.Calls++
}

type CountdownOperationsSpy struct {
    Calls []string
}

func (s *CountdownOperationsSpy) Sleep() {
    s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
    s.Calls = append(s.Calls, write)
    return
}

type SpyTime struct {
    durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
    s.durationSlept = duration
}

const write = "write"
const sleep = "sleep"

func TestCountdown(t *testing.T) {
    t.Run("print 3 to Go!", func(t *testing.T) {
        buffer := &bytes.Buffer{}
        spySleeper := &SpySleeper{}

        Countdown(buffer, spySleeper)

        actual := buffer.String()
        expected := `3
2
1
Go!`

        if actual != expected {
            t.Errorf("Expected %q but got actual %q", expected, actual)
        }

        if spySleeper.Calls != 4 {
            t.Errorf("Not enough calls to sleeper, expected 4 got %d", spySleeper.Calls)
        }
    })

    t.Run("sleep before every print", func(t *testing.T) {
        spySleepPrinter := &CountdownOperationsSpy{}
        Countdown(spySleepPrinter, spySleepPrinter)

        expected := []string{
            sleep,
            write,
            sleep,
            write,
            sleep,
            write,
            sleep,
            write,
        }

        if !reflect.DeepEqual(expected, spySleepPrinter.Calls) {
            t.Errorf("Expected call sequence %v got actual %v", expected, spySleepPrinter.Calls)
        }
    })
}

func TestConfigurableSleeper(t *testing.T) {
    sleepTime := 5 * time.Second

    spyTime := &SpyTime{}
    sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
    sleeper.Sleep()

    if spyTime.durationSlept != sleepTime {
        t.Errorf("Expected slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
    }

}