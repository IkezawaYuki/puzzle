package main

import (
	"fmt"
)

type Schedule struct {
	Start int
	End   int
}

type Schedules []*Schedule

func main2() {
	var sched = Schedules{
		&Schedule{
			Start: 7,
			End:   10,
		},
		&Schedule{
			Start: 8,
			End:   11,
		},
		&Schedule{
			Start: 9,
			End:   12,
		},
	}

	start := sched[0].Start
	end := sched[0].End

	count := celebrityDesity(sched, start, end)

	var maxcount int = 0
	var time int = 0
	for i := start; i <= end; i++ {
		if count[i] > maxcount {
			maxcount = count[i]
			time = i
		}
	}
	fmt.Printf("Best time to attend the party is at %v o'clock: %v celebrities will be attending!\n", time, maxcount)
}

func celebrityDesity(sched Schedules, start int, end int) (count []int) {
	count = make([]int, end+1)
	for i := start; i <= end; i++ {
		count[i] = 0
		for _, c := range sched {
			if (c.Start <= i) && (c.End > i) {
				count[i] += 1
			}
		}
	}
	return
}
