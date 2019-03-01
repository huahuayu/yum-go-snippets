package main

import (
	"github.com/kami-zh/go-capturer"
	"os"
	"strings"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}

func TestPrintDeck(t *testing.T) {
	d := newDeck()

	// capture stdout to variable out, so we can check it later
	out := capturer.CaptureStdout(func() {
		d.print()
	})

	ss := strings.Split(out,"\n")

	if ss[0] != "0 Ace of Spades"{
		t.Errorf("output is not right")
	}
}
