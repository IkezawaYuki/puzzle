package main

import (
	"fmt"
)

type Schedule struct {
	Start int
	End   int
}

func bestTimeToParty(schedules []Schedule){
	start := schedules[0].Start
	end := schedules[0].End

	for _, c := range schedules{
		start = min(c.Start, start)
		end = max(c.End, end)
	}

	count := celebrityDesity(schedules, start, end)
	maxCount := 0
	var time int = 0
	for i := start; i <= end; i++ {
		if count[i] > maxCount {
			maxCount = count[i]
			time = i
		}
	}
	fmt.Printf("Best time to attend the party is at %v o'clock: %v celebrities will be attending!\n", time, maxCount)

}

func max(a, b int)int{
	if a > b{
		return a
	}
	return b
}

func min(a, b int)int{
	if a < b{
		return a
	}
	return b
}

func celebrityDesity(sched []Schedule, start, end int) (count []int) {
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
