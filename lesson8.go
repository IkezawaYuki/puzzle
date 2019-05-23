package main

import "math"

func main() {
	dislikePairs := [][]string{{"Alice", "Bob"}, {"Bob", "Eve"}}
	guestList := []string{"Alice", "Bob", "Cleo", "Don", "Eve"}
	inviteDinner(guestList, dislikePairs)
}

func inviteDinner(guestList []string, dislikePairs [][]string) {
	allCombList := Combinations(len(guestList), guestList)
}

func Combinations(n int, guestList []string) (allCombList [][]string) {
	a := math.Pow(2, float64(n))
	for i := 0; i < int(a); i++ {
		num := i
		cList := []string{}
		for j := 0; j < n; j++ {
			if num%2 == 1 {
				cList = append([]string{guestList[n-1-j]}, cList...)
			}
			num = num / 2
		}
		allCombList = append(allCombList, cList)
	}
	return
}
