package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	howHardIsTheCrystal(10, 2)
}

func howHardIsTheCrystal(n int, d int) {
	radix := 1.0
	for math.Pow(radix, float64(d)) <= 0 {
		radix = radix + 1.0
	}
	r := int(radix)
	fmt.Printf("Radix chosen is %v\n", radix)

	newd := d
	for math.Pow(float64(radix), float64(newd-1)) > float64(n) {
		newd--
	}
	if newd < d {
		fmt.Println("Using only", newd, "balls")
	}
	d = newd

	numDrops := 0
	floorNoBreak := make([]int, d)
	scan := bufio.NewScanner(os.Stdin)
	for i := 0; i < d; i++ {
		for j := 0; j < r; j++ {
			floorNoBreak[i]++
			Floor := convertToDecimal(r, d, floorNoBreak)
			if Floor > n {
				floorNoBreak[i]--
				break
			}
			fmt.Printf("drop ball %v from floor\n", Floor)
			fmt.Print("Did the ball break (yes/no)? :")
			scan.Scan()

			yes := scan.Text()
			numDrops++
			if yes == "yes" {
				floorNoBreak[i]--
				break
			}
		}
	}
	hardness := convertToDecimal(r, d, floorNoBreak)
	fmt.Println("hardness cofficient is", hardness)
	fmt.Println("total number of drops is", numDrops)
}

func convertToDecimal(r int, d int, rep []int) (number int) {
	number = 0
	for i := 0; i < d-1; i++ {
		number = (number + rep[i]) * r
	}
	number += rep[d-1]
	return
}
