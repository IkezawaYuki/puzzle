package main

import (
	"bufio"
	"fmt"
	"os"
)

var deck = []string{"A_C", "A_D", "A_H", "A_S", "2_C", "2_D", "2_H", "2_S", "3_C", "3_D", "3_H", "3_S",
	"4_C", "4_D", "4_H", "4_S", "5_C", "5_D", "5_H", "5_S", "6_C", "6_D", "6_H", "6_S",
	"7_C", "7_D", "7_H", "7_S", "8_C", "8_D", "8_H", "8_S", "9_C", "9_D", "9_H", "9_S",
	"10_C", "10_D", "10_H", "10_S", "J_C", "J_D", "J_H", "J_S",
	"Q_C", "Q_D", "Q_H", "Q_S", "K_C", "K_D", "K_H", "K_S"}

func index(d []string, card string) (index int) {
	for index, value := range d {
		if value == card {
			return index
		}
	}
	return -1
}

func main() {
	fmt.Println("hello")
	fmt.Printf("order is: %v\n", deck)

	cards := []string{}
	cind := []int{}
	cardsuits := []int{}
	cnumbers := []int{}

	numsuits := make([]int, 4)

	var pairsuit int
	for i := 0; i < 5; i++ {
		fmt.Printf("Please give card %v ", i+1)
		scan := bufio.NewScanner(os.Stdin)
		scan.Scan()
		card := scan.Text()
		cards = append(cards, card)
		n := index(deck, card)
		cind = append(cind, n)
		cardsuits = append(cardsuits, n%4)
		cnumbers = append(cnumbers, n/4)
		numsuits[n%4]++
		if numsuits[n%4] > 1 {
			pairsuit = n % 4
		}
	}

	cardh := []int{}
	for i := 0; i < 5; i++ {
		if cardsuits[i] == pairsuit {
			cardh = append(cardh, i)
		}
	}
	hidden, other, encode := outputFirstCard(cnumbers, cardh, cards)

	remindices := []int{}
	for i := 0; i < 5; i++ {
		remindices = append(remindices, cind[i])
	}

	remindices = sortList(remindices)

}

func outputFirstCard(numbers []int, oneTwo []int, cards []string) (hidden int, other int, encode int) {
	encode = (numbers[oneTwo[0]] - numbers[oneTwo[1]]) % 13
	if encode > 0 && encode <= 6 {
		hidden = oneTwo[0]
		other = oneTwo[1]
	} else {
		hidden = oneTwo[1]
		other = oneTwo[0]
		encode = (numbers[oneTwo[1]] - numbers[oneTwo[0]]) % 13
	}
	fmt.Printf("First card is: %v\n", cards[other])
	return
}

func sortList(tlist []int) (slist []int) {
	for i := 0; i < len(tlist)-1; i++ {
		ismall := i
		for j := i; j < len(tlist); j++ {
			if tlist[ismall] > tlist[j] {
				ismall = j
			}
		}
		tlist[i], tlist[ismall] = tlist[ismall], tlist[i]
	}
	slist = tlist
	return
}

func MagicGuessCard() {

}
