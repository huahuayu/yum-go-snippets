package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// 初始化一副牌
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// 打印牌
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// 发牌
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	// deck转化为stringslice 每个元素用逗号分隔后join起来成为string,再把string转化为[]byte保存
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	// 读文件获得byteslice
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - Log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// 将byteslice转化为string用逗号拆分为stringslice
	s := strings.Split(string(bs), ",")

	// 将string slice强制转化为deck type 返回
	return deck(s)
}

// 洗牌
func (d deck) shuffle() {
	// 把时间戳作为随机数种子
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
