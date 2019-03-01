package main

import "fmt"

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	d1,d2 := deal(cards,8)
	fmt.Println(d1)
	fmt.Println("______________________")
	fmt.Println(d2)
}
