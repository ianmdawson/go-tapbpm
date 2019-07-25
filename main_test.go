package main

import (
	"bytes"
	"testing"
	"time"

	"github.com/gosuri/uilive"
)

func initTestWriter() *uilive.Writer {
	testWriter := uilive.New()
	// start listening for updates and render

	testWriterBuffer := &bytes.Buffer{}
	testWriter.Out = testWriterBuffer
	testWriter.Start()
	return testWriter
}

func TestHandleInputReset(t *testing.T) {
	now := time.Now()
	testTapTracker := tapTracker{&now, time.Minute * 1, 60}
	testWriter := initTestWriter()

	handleInput(&testTapTracker, rune(114), 0, testWriter)
	if testTapTracker.totalTime != 0 {
		t.Errorf("Expected totalTime to be reset to 0, got: %v", testTapTracker.totalTime)
	}

	// TODO: test output
	testWriter.Stop()
}

func TestHandleInputTap(t *testing.T) {
	now := time.Now()
	testTotalTime := time.Minute * 1
	testTapTracker := tapTracker{&now, testTotalTime, 60}
	testWriter := initTestWriter()

	handleInput(&testTapTracker, rune(5), 0, testWriter)
	if testTapTracker.totalTime <= testTotalTime {
		t.Errorf("Expected totalTime to be greater than %v, got: %v", testTotalTime, testTapTracker.totalTime)
	}

	// TODO: test output

	testWriter.Stop()
}
