package deck

import (
	"math/rand"
	"testing"
)

func TestCards(t *testing.T) {
	aceOfHearts := Card{Rank: Ace, Suit: Heart}
	twoOfClub := Card{Rank: Two, Suit: Club}
	jackOfDiamond := Card{Rank: Jack, Suit: Diamond}
	kingOfSpade := Card{Rank: King, Suit: Spade}
	joker := Card{Suit: Joker}

	if aceOfHearts.String() != "Ace of Hearts" {
		t.Errorf("got '%s' , want 'Ace of Hearts'", aceOfHearts.String())
	}
	if twoOfClub.String() != "Two of Clubs" {
		t.Errorf("got '%s' , want 'Two of Clubs'", twoOfClub.String())
	}
	if jackOfDiamond.String() != "Jack of Diamonds" {
		t.Errorf("got '%s' , want 'Jack of Diamonds'", jackOfDiamond.String())
	}
	if kingOfSpade.String() != "King of Spades" {
		t.Errorf("got '%s' , want 'King of Spades'", kingOfSpade.String())
	}
	if joker.String() != "Joker" {
		t.Errorf("got '%s' , want 'Joker'", joker.String())
	}
}

func TestNew(t *testing.T) {
	deck := New()

	if len(deck) != 13*4 {
		t.Errorf("got '%d' , want 52", len(deck))
	}
}

func TestDefaultSort(t *testing.T) {
	deck := New(DefaultSort)

	expectedFirstCard := Card{Suit: Spade, Rank: Ace}
	if deck[0] != expectedFirstCard {
		t.Errorf("got: %s, expected: %s", deck[0], expectedFirstCard)
	}

	expectedSecondCard := Card{Suit: Heart, Rank: King}
	if deck[51] != expectedSecondCard {
		t.Errorf("got: %s, expected: %s", deck[51], expectedSecondCard)
	}
}

func TestSort(t *testing.T) {
	deck := New(Sort(Less))

	expectedFirstCard := Card{Suit: Spade, Rank: Ace}
	if deck[0] != expectedFirstCard {
		t.Errorf("got: %s, expected: %s", deck[0], expectedFirstCard)
	}

	expectedSecondCard := Card{Suit: Heart, Rank: King}
	if deck[51] != expectedSecondCard {
		t.Errorf("got: %s, expected: %s", deck[51], expectedSecondCard)
	}
}

func TestShuffle(t *testing.T) {
	shuffleRand = rand.New(rand.NewSource(0))

	orig := New()
	first := orig[40]
	second := orig[35]

	deck := New(Shuffle)
	if deck[0] != first {
		t.Errorf("got: %s, expected: %s", deck[0], first)
	}
	if deck[1] != second {
		t.Errorf("got: %s, expected: %s", deck[1], second)
	}
}

func TestJokers(t *testing.T) {
	expectedJokers := 3
	deck := New(Jokers(expectedJokers))
	count := 0

	for _, card := range deck {
		if card.Suit == Joker {
			count++
		}
	}

	if count != expectedJokers {
		t.Errorf("got: %d, expected: %d", count, expectedJokers)
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}

	deck := New(Filter(filter))
	for _, card := range deck {
		if card.Rank == Two || card.Rank == Three {
			t.Error("found two and threes which should have been filtered")
		}
	}
}

func TestDeck(t *testing.T) {
	deck := New(Deck(3))
	if 4*13*3 != len(deck) {
		t.Errorf("got: %d, expected: %d", len(deck), 4*14*3)
	}
}
