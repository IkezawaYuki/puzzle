package main

import "testing"

func TestBestTimeToParty(t *testing.T) {
	var sched = []Schedule{
		Schedule{
			Start: 7,
			End:   10,
		},
		Schedule{
			Start: 8,
			End:   11,
		},
		Schedule{
			Start: 9,
			End:   12,
		},
	}
	bestTimeToParty(sched)
}
