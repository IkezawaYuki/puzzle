package main

import "fmt"

type Interval struct {
	Start int
	End   int
	Caps  string
}
type Intervals []*Interval

func main() {
	caps1 := []string{"F", "F", "B", "B", "F", "F", "B", "B"}
	//caps2 := []string{"F", "B", "B", "B", "F", "B", "B", "F"}

	start := 0
	forward := 0
	backward := 0
	var intervals Intervals
	caps1 = append(caps1, "end")

	for i := 1; i < len(caps1); i++ {
		if caps1[start] != caps1[i] {
			var inerval = &Interval{
				Start: start,
				End:   i - 1,
				Caps:  caps1[start],
			}
			intervals = append(intervals, inerval)

			if caps1[start] == "F" {
				forward++
			} else {
				backward++
			}
			start = i
		}
	}
	var flip string
	if forward < backward {
		flip = "F"
	} else {
		flip = "B"
	}
	for _, inter := range intervals {
		if inter.Caps == flip {
			fmt.Printf("People in positions %v throwgh %v flip your caps\n", inter.Start, inter.End)
		}
	}
}
