/*
Package main implements go-tapbpm command-line tool
*/
package main

import (
	"fmt"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/gosuri/uilive"
)

// inputIsQuitKey determines if the character or key combination should quit the program
func inputIsQuitKey(char rune, key keyboard.Key) bool {
	if string(char) == "q" || key == keyboard.KeyCtrlC || key == keyboard.KeyEsc {
		return true
	}
	return false
}

// handleInput handles user input from the keyboard, records key presses, and updates the UI
func handleInput(trkr *tapTracker, char rune, key keyboard.Key, writer *uilive.Writer) {
	if string(char) == "r" {
		fmt.Fprintln(writer, "Resetting counts...")
		trkr.reset()
		return
	}

	t := time.Now()
	trkr.tap(t)

	if trkr.numberOfTaps == 0 {
		fmt.Fprintf(writer, "First tap...\n")
	} else {
		fmt.Fprintf(writer, "bpm: %v\n", trkr.bpmString())
	}

	return
}

// initTerminalWriter initializes the terminal writer
func initTerminalWriter() *uilive.Writer {
	// TODO: return a writer that doesn't update as frequently as every millisecond to improve performance
	writer := uilive.New()
	// start listening for updates and render
	writer.Start()
	return writer
}

// main loops awaiting user input until receiving quit signal
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
