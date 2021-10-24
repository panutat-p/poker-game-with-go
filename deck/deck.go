package deck

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string // user-defined type

func Run() {
	d1 := newDeck()
	d1.print()

	h1 := d1.deal(40)
	h1.print() // go to hand1
	d1.print() // remaining

	d1.shuffle()
	d1.print()
}

// method of type deck,
// deck is a receiver
func (d deck) print() {
	fmt.Printf("\n%T\n", d)
	fmt.Println("len:", len(d))
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

// function
// create a new deck
func newDeck() deck {
	var d deck

	s := []string{"Spades", "Diamonds", "Hearts", "Clubs"}                               // card suite
	v := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13"} // card value

	for _, suit := range s {
		for _, value := range v {
			d = append(d, fmt.Sprintf("%s of %s", suit, value))
		}
	}

	return d
}

// pointer receiver
// d is &d, *d is d
// h: hand from the deck
// r: remaining in the deck
func (d *deck) deal(size int) deck {
	h := (*d)[:size]
	*d = (*d)[size:]
	return h
}

// write a file to disk
func (d deck) saveToFile(filename string) error {
	bsl := []byte(d.toString()) // string to byte-slice
	err := os.WriteFile(filename, bsl, 0666)
	return err
}

// d is already a string-slice
// remove redundant conversion, "[]byte(d)"
func (d deck) toString() string {
	return strings.Join(d, ",")
}

// read a file from disk
func newDeckFromFile(filename string) deck {
	bsl, err := os.ReadFile(filename) // get a file in byte-slice format
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	s := string(bsl)             // string of cards
	ssl := strings.Split(s, ",") // string-slice
	d := deck(ssl)               // deck type
	return d
}

// In this case, d is the same as (*d)
// shuffle by swapping 2 elements by its index
// randomize index in range [0, len(d)-1] >> [0, 51] for full deck
// do len(d) times >> 52 times for full deck
func (d *deck) shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(*d), func(i int, j int) {
		(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
	})
}
