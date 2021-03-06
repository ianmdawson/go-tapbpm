package main

import (
	"strconv"
	"time"
)

type tapTracker struct {
	lastTapTime   *time.Time
	totalDuration time.Duration
	numberOfTaps  int
}

// reset all the tapTracker fields to initial values
func (trkr *tapTracker) reset() {
	trkr.totalDuration = time.Duration(0)
	trkr.numberOfTaps = 0
	trkr.lastTapTime = nil
}

// record a new tap
// The very first tap initializes the start time to (lastTapTime), but
// doesn't record a tap in numberOfTaps for ease of calculation
func (trkr *tapTracker) tap(newTime time.Time) {
	prevTime := trkr.lastTapTime
	trkr.lastTapTime = &newTime
	if prevTime != nil {
		trkr.totalDuration = trkr.totalDuration + newTime.Sub(*prevTime)
	}
	if prevTime == nil {
		return
	}
	trkr.numberOfTaps++
}

// Calculate the BPM: numberOfTaps/(totalDuration in minutes)
// returns 0 if BPM can't be calculated
func (trkr *tapTracker) bpm() float64 {
	if trkr.lastTapTime == nil || trkr.numberOfTaps == 0 {
		return float64(0)
	}

	if trkr.numberOfTaps == 0 {
		return trkr.totalDuration.Minutes()
	}

	bpm := (float64(trkr.numberOfTaps) / trkr.totalDuration.Minutes())
	return bpm
}

func (trkr *tapTracker) bpmString() string {
	return strconv.FormatFloat(trkr.bpm(), 'f', 2, 64)
}
