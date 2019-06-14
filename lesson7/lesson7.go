package main

import "fmt"

func main() {
	Ls := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	lsearch(Ls, 43)
	bsearch(Ls, 43)
}

func lsearch(l []int, target int) {
	for index, value := range l {
		if value == target {
			fmt.Printf("location is %v\n", index)
			return
		}
	}
	fmt.Println("Could not find the value ", target)
}

func bsearch(l []int, target int) {
	lo, hi := 0, len(l)-1
	for lo <= hi {
		mid := (lo + hi) / 2
		if l[mid] < target {
			lo = mid + 1
		} else if l[mid] > target {
			hi = mid - 1
		} else {
			fmt.Println("found at location", mid)
			return
		}
	}
	fmt.Println("Could not find the value", target)
}
