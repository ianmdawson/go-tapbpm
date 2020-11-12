package main

import (
	"testing"
	"time"
)

func TestTapTrackerBPM(t *testing.T) {
	testCases := []struct {
		lastTapTime      time.Time
		duration         time.Duration
		taps             int
		expectedBpmValue float64
	}{
		{time.Now(), time.Minute * 1, 60, float64(60)},
		{time.Now(), time.Minute * 2, 120, float64(60)},
		{time.Now(), time.Second * 30, 60, float64(120)},
		{time.Now(), time.Second, 0, float64(0)},
	}

	for _, testCase := range testCases {
		trkr := tapTracker{&testCase.lastTapTime, testCase.duration, testCase.taps}

		bpm := trkr.bpm()
		expectedValue := float64(60)
		if bpm != testCase.expectedBpmValue {
			t.Errorf("Expected BPM to be %v, got: %v", expectedValue, bpm)
		}
	}
}

func TestTapTrackerBpmString(t *testing.T) {
	now := time.Now()
	minute := time.Minute * 1
	trkr := tapTracker{&now, minute, 2}

	bpmString := trkr.bpmString()
	expectedValue := "2.00"
	if bpmString != expectedValue {
		t.Errorf("Expected BPM to be %v, got: %v", expectedValue, bpmString)
	}
}

func TestTapTrackerFirstBeat(t *testing.T) {
	minute := time.Minute * 1
	trkr := tapTracker{nil, minute, 1}

	bpm := trkr.bpm()
	expectedValue := float64(0)
	if bpm != expectedValue {
		t.Errorf("Expected BPM to be %v, got: %v", expectedValue, bpm)
	}
}

func TestTapTrackerReset(t *testing.T) {
	now := time.Now()
	trkr := tapTracker{&now, time.Minute * 1, 1}
	trkr.reset()

	if trkr.lastTapTime != nil {
		t.Errorf("Expected lastTapTime to be nil, got: %v", trkr.lastTapTime)
	}

	if trkr.totalTime != 0 {
		t.Errorf("Expected totalTime to be 0, got: %v", trkr.totalTime)
	}

	if trkr.numberOfTaps != 0 {
		t.Errorf("Expected numberOfTaps to be 0, got: %v", trkr.numberOfTaps)
	}
}

func TestTapTrackerTap(t *testing.T) {
	startingTotalTime := time.Nanosecond - 1
	trkr := tapTracker{nil, startingTotalTime, 0}

	newTime := time.Now()
	trkr.tap(newTime)

	if trkr.lastTapTime.Equal(newTime) != true {
		t.Errorf("Expected numberOfTaps to be %v, got: %v", newTime, trkr.lastTapTime)
	}

	expectedNumberOfTaps := 0
	if trkr.numberOfTaps != expectedNumberOfTaps {
		t.Errorf("Expected numberOfTaps to be %v, got: %v", expectedNumberOfTaps, trkr.numberOfTaps)
	}

	if trkr.totalTime < startingTotalTime {
		t.Errorf("Expected startingTotalTime (%v) to less than ending total time %v", startingTotalTime, trkr.numberOfTaps)
	}

	trkr.tap(newTime)
	expectedNumberOfTaps = 1
	if trkr.numberOfTaps != expectedNumberOfTaps {
		t.Errorf("Expected numberOfTaps to be %v, got: %v", expectedNumberOfTaps, trkr.numberOfTaps)
	}
}
