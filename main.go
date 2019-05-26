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

func inputIsQuitKey(char rune, key keyboard.Key) bool {
	if string(char) == "q" || key == keyboard.KeyCtrlC || key == keyboard.KeyEsc {
		return true
	}
	return false
}

func handleInput(trkr *tapTracker, char rune, key keyboard.Key) (bool, error) {
	if inputIsQuitKey(char, key) == true {
		return false, nil
	}

	if string(char) == "r" {
		fmt.Println("Resetting counts...")
		trkr.reset()
		return true, nil
	}

	t := time.Now()
	trkr.tap(t)
	println(trkr.numberOfTaps)
	println("bpm: ", trkr.bpmString())
	return true, nil
}

func main() {
	fmt.Println("Tap BPM, tap a letter to track bpm -- 'r' to reset -- 'q' or ESC to quit")
	fmt.Println("---------------------")

	trkr := tapTracker{nil, 0, 0}

	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		continueLoop, err := handleInput(&trkr, char, key)
		if err != nil {
			panic(err)
		}

		if continueLoop == false {
			fmt.Println("Goodbye...")
			break
		}
	}
}
