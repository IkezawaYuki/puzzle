package main

import (
	"fmt"
	"math"
)

func main() {
	dislikePairs := [][]string{{"Alice", "Bob"}, {"Bob", "Eve"}}
	guestList := []string{"Alice", "Bob", "Cleo", "Don", "Eve"}
	inviteDinner(guestList, dislikePairs)
}

func inviteDinner(guestList []string, dislikePairs [][]string) {
	allCombList := Combinations(len(guestList), guestList)
	allGoodCombinations := removeBadCombinations(allCombList, dislikePairs)
	invite := []string{}
	for _, v := range allGoodCombinations {
		if len(v) > len(invite) {
			invite = v[:]
		}
		fmt.Println("Optinum Solution", invite)
	}
}

func Contains(target string, group []string) bool {
	for _, v := range group {
		if v == target {
			return true
		}
	}
	return false
}

func removeBadCombinations(allCombList [][]string, dislikePairs [][]string) (allGoodCombinations [][]string) {
	allGoodCombinations = [][]string{}
	for _, v := range allCombList {
		good := true
		for _, j := range dislikePairs {
			if Contains(j[0], v) && Contains(j[1], v) {
				good = false
			}
		}
		if good {
			allGoodCombinations = append(allGoodCombinations, v)
			fmt.Println(allGoodCombinations)
		}
	}
	return
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
