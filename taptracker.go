package main

import (
	"strconv"
	"time"
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
	if prevTime == nil {
		return
	}
	trkr.numberOfTaps++
}

func (trkr *tapTracker) bpm() float64 {
	if trkr.trackedTime == nil {
		return float64(0)
	}

	if trkr.numberOfTaps == 0 {
		return trkr.totalTime.Minutes()
	}

	bpm := (float64(trkr.numberOfTaps) / trkr.totalTime.Minutes())
	return bpm
}

func (trkr *tapTracker) bpmString() string {
	return strconv.FormatFloat(trkr.bpm(), 'f', 2, 64)
}
