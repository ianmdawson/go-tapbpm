package main

import (
	"fmt"
	"time"

	"github.com/eiannone/keyboard"
)

func inputIsQuitKey(char rune, key keyboard.Key) bool {
	if string(char) == "q" || key == keyboard.KeyCtrlC || key == keyboard.KeyEsc {
		return true
	}
	return false
}

func handleInput(trkr *tapTracker, char rune, key keyboard.Key) {
	if string(char) == "r" {
		fmt.Println("Resetting counts...")
		trkr.reset()
		return
	}

	t := time.Now()
	trkr.tap(t)
	println(trkr.numberOfTaps)
	println("bpm: ", trkr.bpmString())
	return
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

		if inputIsQuitKey(char, key) == true {
			fmt.Println("Goodbye...")
			break
		}

		handleInput(&trkr, char, key)
	}
}
