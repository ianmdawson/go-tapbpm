package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/eiannone/keyboard"
)

type tapTracker struct {
	trackedTime  *time.Time // TODO: rename this previous tap time?
	totalTime    time.Duration
	numberOfTaps int
}

func (trkr *tapTracker) reset() {
	trkr.totalTime = time.Nanosecond - 1 // TODO: is there a cleaner way to do this?
	trkr.numberOfTaps = 0
	trkr.trackedTime = nil
}

func (trkr *tapTracker) tap(newTime time.Time) {
	prevTime := trkr.trackedTime
	trkr.trackedTime = &newTime
	if prevTime != nil {
		trkr.totalTime = trkr.totalTime + newTime.Sub(*prevTime)
	}
	trkr.numberOfTaps++
}

func (trkr *tapTracker) bpm() float64 {
	if trkr.trackedTime == nil {
		return float64(0)
	}
	bpm := (float64(trkr.numberOfTaps) / trkr.totalTime.Minutes())
	return bpm
}

func (trkr *tapTracker) bpmString() string {
	return strconv.FormatFloat(trkr.bpm(), 'f', 2, 64)
}

func main() {
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")

	trkr := tapTracker{nil, 0, 0}
	fmt.Println("Press ESC to quit")
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		} else if string(char) == "q" || char == '\x00' || key == keyboard.KeyEsc {
			fmt.Println("Goodbye...")
			break
		} else if string(char) == "r" {
			fmt.Println("Resetting counts...")
			trkr.reset()
		} else {
			t := time.Now()
			trkr.tap(t)
			println(trkr.numberOfTaps)
			bpm := strconv.FormatFloat(trkr.bpm(), 'f', 2, 64)
			println("bpm: ", bpm)
		}
	}
}
