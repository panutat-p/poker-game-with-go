package deck

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Spades of Ace" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deck-testing.txt")

	deck := newDeck()

	deck.saveToFile("_deck-testing.txt")

	loadedDeck := newDeckFromFile("_deck-testing.txt")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, but got %v", len(loadedDeck))
	}

	os.Remove("_deck-testing.txt")
}
