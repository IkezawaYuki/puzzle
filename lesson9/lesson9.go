package main

import (
	"fmt"
	"math"
)

func Good(Comb []string, candList []string, candTalents [][]string, AllTalents []string) bool {
	for i := 0; i < len(AllTalents); i++ {
		cover := false
		tal := AllTalents[i]
		for j := 0; j < len(Comb); j++ {
			cand := Comb[j]
			index := index(candList, cand)
			candTal := candTalents[index]
			if contains(tal, candTal) {
				cover = true
			}
		}
		if cover == false {
			return false
		}
	}
	return true
}

func Hire4Show(candList []string, candTalents [][]string, talentList []string) {
	n := len(candList)
	hire := candList[:]
	p := math.Pow(2, float64(n))
	for i := 0; i < int(p); i++ {
		combination := []string{}
		num := i
		for j := 0; j < n; j++ {
			if num%2 == 1 {
				combination = append([]string{candList[n-1-j]}, combination...)
			}
			num = num / 2
		}
		if Good(combination, candList, candTalents, talentList) {
			if len(hire) > len(combination) {
				hire = combination
			}
		}
	}
	fmt.Println("Optimum Solution:", hire)
}

func main() {
	var talents = []string{"Sing", "Dance", "Magic", "Act", "Flex", "Code"}
	var candidates = []string{"Aly", "Bob", "Cal", "Don", "Eve", "Fay"}
	var candidatesTalents = [][]string{{"Flex", "Code"}, {"Dance", "Magic"}, {"Sing", "Magic"}, {"Sing", "Dance"}, {"Dance", "Act", "Code"}, {"Act", "Code"}}
	Hire4Show(candidates, candidatesTalents, talents)
}

func contains(target string, group []string) bool {
	for _, v := range group {
		if v == target {
			return true
		}
	}
	return false
}

func index(d []string, card string) (index int) {
	for index, value := range d {
		if value == card {
			return index
		}
	}
	return -1
}
