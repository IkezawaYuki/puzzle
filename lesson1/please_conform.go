package main

import "fmt"

type Interval struct {
	Start int
	End   int
	Caps  string
}

func pleaseConform(caps []string){
	start := 0
	forward := 0
	backward := 0

	intervals := []*Interval{}
	caps = append(caps, "end")

	for i := 1; i < len(caps); i++ {
		if caps[start] != caps[i] {
			var inerval = &Interval{
				Start: start,
				End:   i - 1,
				Caps:  caps[start],
			}
			intervals = append(intervals, inerval)

			if caps[start] == "F" {
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


