// Package deck provides functions to create a deck of cards
package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Suit represents the 4 suits of deck
type Suit int8

// There are 4 [Suit] in a deck: Spade, Diamond, Club, and Heart 
const (
    Spade Suit = iota
    Diamond
    Club
    Heart
    Joker // Joker is not a suit but is included for special case
)

var suits = []Suit{Spade, Diamond, Club, Heart}

// Rank represents the 13 cards in a suit 
type Rank int8

// There are 13 [Rank] cards in a [Suit]: Ace, Two, Three, ..., King
const (
    _ Rank = iota // to skip the 0th value
    Ace
    Two
    Three
    Four
    Five
    Six
    Seven
    Eight
    Nine
    Ten
    Jack
    Queen
    King
)

const (
    minRank = Ace
    maxRank = King
)

// Card represents a card with a [Suit] and [Rank]
type Card struct{
    Suit
    Rank
}

// String is used to print a string version of the Card
//
// Card{Suit: Spade, Rank: Ace} will give "Ace of Spades"
func (c Card) String() string {
    if c.Suit == Joker {
        return c.Suit.String() 
    }
    return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

// New returns a new deck of cards which are modified by the
// functional options opts
func New(opts ...func(c []Card)[]Card) []Card {
 cards := []Card{}

 for _, suit := range suits {
    for rank := minRank; rank <= maxRank; rank++ {
        cards = append(cards, Card{Suit: suit, Rank: rank})
    }
 }

 for _, opt := range opts {
    cards = opt(cards)
 }

 return cards
}

// Less is a helper function for sorting cards based on their absolute [Rank]
func Less(cards []Card) func(i, j int) bool {
   return func(i, j int) bool {
    return absRank(cards[i]) < absRank(cards[j])
   } 
}

// DefaultSort returns a deck of cards after sorting the cards
func DefaultSort(cards []Card) []Card {
   sort.Slice(cards, Less(cards)) 
   return cards
}

// Sort returns a sorting function based on the provided comparison function
func Sort(less func(cards []Card) func(i, j int) bool) func (cards []Card) []Card {
   return func(cards []Card) []Card {
       sort.Slice(cards, less(cards)) 
       return cards
   } 
}

// absRank calculates the absolute rank of a card by combining the suit and rnak
func absRank(c Card) int{
    return int(c.Suit) * int(maxRank) + int(c.Rank); 
}

var shuffleRand = rand.New(rand.NewSource(time.Now().Unix()))

// Shuffle returns a deck of cards shuffled in a random order
func Shuffle(cards []Card) []Card {
    ret := make([]Card, len(cards))
    perm := shuffleRand.Perm(len(cards))
    for i, j := range perm {
        ret[i] = cards[j]
    }

    return ret
}

// Jokers returns a deck of cards with n number of [Joker]
func Jokers(n int) func (cards []Card) []Card {
    return func(cards []Card) []Card {
        for i := 0; i<n; i++ {
            cards = append(cards, Card{
                Rank: Rank(i),
                Suit: Joker,
            })
        } 
        return cards
    }
}

// Filter returns a deck of cards which is filtered by the condition f
func Filter(f func(card Card) bool) func(cards []Card) []Card {
    return func (cards []Card) []Card {
        var ret []Card

        for _, card := range cards {
            if !f(card) {
                ret = append(ret, card)
            }
        }

        return ret
    }
}

// Deck returns a deck of cards which contains the existing deck n times
func Deck(n int) func(cards []Card) []Card {
    return func (cards []Card) []Card {
        var ret []Card
        for i:=0; i<n; i++ {
            ret = append(ret, cards...)
        }
        return ret
    }
}


