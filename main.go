package main

import (
	"fmt"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/gosuri/uilive"
)

func inputIsQuitKey(char rune, key keyboard.Key) bool {
	if string(char) == "q" || key == keyboard.KeyCtrlC || key == keyboard.KeyEsc {
		return true
	}
	return false
}

func handleInput(trkr *tapTracker, char rune, key keyboard.Key, writer *uilive.Writer) {
	if string(char) == "r" {
		fmt.Fprintln(writer, "Resetting counts...")
		trkr.reset()
		return
	}

	t := time.Now()
	trkr.tap(t)

	if trkr.trackedTime != nil {
		fmt.Fprintf(writer, "bpm: %v\n", trkr.bpmString())
	} else {
		fmt.Fprintf(writer, "First tap...\n")
	}

	return
}

func initTerminalWriter() *uilive.Writer {
	writer := uilive.New()
	// start listening for updates and render
	writer.Start()
	return writer
}

func main() {
	fmt.Println("Tap BPM, tap a letter to track bpm -- 'r' to reset -- 'q' or ESC to quit")

	writer := initTerminalWriter()
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

		handleInput(&trkr, char, key, writer)
	}

	writer.Stop() // flush output and stop rendering
}
