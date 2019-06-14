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

func Index(d []string, card string) (index int) {
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
		n := Index(deck, card)
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
		if i != hidden && i != other {
			remindices = append(remindices, cind[i])
		}
	}

	remindices = sortList(remindices)
	outputNext3Cards(encode, remindices)

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

func outputNext3Cards(code int, ind []int) (second int, third int, fourth int) {
	switch code {
	case 1:
		second = ind[0]
		third = ind[1]
		fourth = ind[2]
	case 2:
		second = ind[0]
		third = ind[2]
		fourth = ind[1]
	case 3:
		second = ind[1]
		third = ind[0]
		fourth = ind[2]
	case 4:
		second = ind[1]
		third = ind[2]
		fourth = ind[0]
	case 5:
		second = ind[2]
		third = ind[0]
		fourth = ind[1]
	case 6:
		second = ind[2]
		third = ind[1]
		fourth = ind[0]
	}
	print("second card is:", deck[second])
	print("third card is:", deck[third])
	print("fourth card is:", deck[fourth])
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

func MagicianGuessesCard() {
	print("ordering is :", deck)
	cards := []string{}
	cind := []int{}
	var suit int
	var number int
	var encode int
	for i := 0; i < 5; i++ {
		scan := bufio.NewScanner(os.Stdin)
		scan.Scan()
		card := scan.Text()
		cards = append(cards, card)
		n := Index(deck, card)
		cind = append(cind, n)
		if i == 0 {
			suit = n % 4
			number = n / 4
		}
	}
	if cind[1] < cind[2] && cind[1] > cind[3] {
		if cind[2] < cind[3] {
			encode = 1
		} else {
			encode = 2
		}
	} else if (cind[1] < cind[2] && cind[1] > cind[3]) || (cind[1] > cind[2] && cind[2] < cind[3]) {
		if cind[2] < cind[3] {
			encode = 3
		} else {
			encode = 4
		}
	} else if cind[1] > cind[2] && cind[1] > cind[3] {
		if cind[2] < cind[3] {
			encode = 5
		} else {
			encode = 6
		}
	}

	var hiddennumber int = (number + encode) % 13
	var index int = hiddennumber*4 + suit
	fmt.Printf("hidden card is:", deck[index])
}

func ComputerAssistant() {
	fmt.Printf("ordering is:", deck)
	cards := []string{}
	cind := []int{}
	cardsuits := []int{}
	cnumbers := []int{}
	numsuits := make([]int, 4)
	number := 0
	var pairsuit int

	for {
		scan := bufio.NewScanner(os.Stdin)
		scan.Scan()
		input := scan.Text()
		if len(input) > 5 {
			break
		}
	}
	for i := 0; i < 5; i++ {
		number = number * (i + 1) / (i + 2)
		n := number % 52
		cards = append(cards, deck[n])
		cind = append(cind, n)
		cardsuits = append(cardsuits, n%4)
		cnumbers = append(cnumbers, n/4)
		numsuits[n%4] += 1
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
		if i != hidden && i != other {
			remindices = append(remindices, cind[i])
		}
	}

	remindices = sortList(remindices)
	outputNext3Cards(encode, remindices)

	scan := bufio.NewScanner(os.Stdin)
	fmt.Print("What is the hidden card?")
	guess := scan.Text()
	if guess == cards[hidden] {
		fmt.Println("you are a mind reader extraordinaire!")
	} else {
		fmt.Println("sorry, not impressed!")
	}
}
