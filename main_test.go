package main

import (
	"testing"
	"time"
)

func TestHandleInputReset(t *testing.T) {
	now := time.Now()
	testTapTracker := tapTracker{&now, time.Minute * 1, 60}

	handleInput(&testTapTracker, rune(114), 0)
	if testTapTracker.totalTime != 0 {
		t.Errorf("Expected totalTime to be reset to 0, got: %v", testTapTracker.totalTime)
	}
}

func TestHandleInputTap(t *testing.T) {
	now := time.Now()
	testTotalTime := time.Minute * 1
	testTapTracker := tapTracker{&now, testTotalTime, 60}

	handleInput(&testTapTracker, rune(5), 0)
	if testTapTracker.totalTime <= testTotalTime {
		t.Errorf("Expected totalTime to be greater than %v, got: %v", testTotalTime, testTapTracker.totalTime)
	}
}
