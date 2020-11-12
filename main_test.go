package main

import (
	"bytes"
	"testing"
	"time"

	"github.com/gosuri/uilive"
)

func TestHandleInputReset(t *testing.T) {
	now := time.Now()
	testTapTracker := tapTracker{&now, time.Minute * 1, 60}
	testWriter := initTestWriter()

	handleInput(&testTapTracker, rune(114), 0, testWriter)
	if testTapTracker.totalDuration != 0 {
		t.Errorf("Expected totalDuration to be reset to 0, got: %v", testTapTracker.totalDuration)
	}

	testWriter.Stop()
}

func TestHandleInputTap(t *testing.T) {
	now := time.Now()
	testTotalDuration := time.Minute * 1
	testTapTracker := tapTracker{&now, testTotalDuration, 60}
	testWriter := initTestWriter()

	handleInput(&testTapTracker, rune(5), 0, testWriter)
	if testTapTracker.totalDuration <= testTotalDuration {
		t.Errorf("Expected totalDuration to be greater than %v, got: %v", testTotalDuration, testTapTracker.totalDuration)
	}

	testWriter.Stop()
}

// Helpers

func initTestWriter() *uilive.Writer {
	testWriter := uilive.New()
	// start listening for updates and render

	testWriterBuffer := &bytes.Buffer{}
	testWriter.Out = testWriterBuffer
	testWriter.Start()
	return testWriter
}
