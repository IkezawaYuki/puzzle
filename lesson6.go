package main

import "fmt"

func main() {
	coinsList := []int{10, 10, 10, 10, 10, 10, 10, 11, 10, 10, 10}
	coinComparison(coinsList)
}

func calcSum(group []int) (sum int) {
	for _, v := range group {
		sum += v
	}
	return
}

func compare(group1, group2 []int) (result string) {
	if calcSum(group1) > calcSum(group2) {
		result = "left"
	} else if calcSum(group1) < calcSum(group2) {
		result = "right"
	} else {
		result = "equal"
	}
	return
}

func findFakeGroup(group1, group2, group3 []int) (fakeGroup []int) {
	result1and2 := compare(group1, group2)
	if result1and2 == "left" {
		fakeGroup = group1
	} else if result1and2 == "right" {
		fakeGroup = group2
	} else {
		fakeGroup = group3
	}
	return
}

func index(group []int, fake int) (index int) {
	for index, value := range group {
		if value == fake {
			return index
		}
	}
	return
}

func coinComparison(coinsLists []int) {
	counter := 0
	currList := coinsLists[:]
	for 1 < len(currList) {
		group1, group2, group3 := splitCoins(currList)
		currList = findFakeGroup(group1, group2, group3)
		fmt.Println(currList)
		counter++
	}
	fake := currList[0]
	i := index(coinsLists, fake)
	fmt.Printf("The fake coin %v in the original list\n", i)
	fmt.Printf("Number of weighings: %v\n", counter)
}

func splitCoins(coinsList []int) (group1, group2, group3 []int) {
	length := len(coinsList)
	group1 = coinsList[:length/3]
	group2 = coinsList[length/3 : length/3*2]
	group3 = coinsList[length/3*2:]
	return
}
