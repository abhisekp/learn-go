//go:build go1.18 || go1.19 || go2

package main

import "math/rand"

type Deck[C any] struct {
	Cards    []C
	randSeed int64
}

type DeckOptions struct {
	RandSeed int64
}

// NewDeck creates a new deck of cards.
// If options is nil, the default options will be used.
// Considers only first and only one argument, DeckOptions.
func NewDeck[C Card](options ...DeckOptions) *Deck[C] {
	option := DeckOptions{
		RandSeed: 99,
	}
	if len(options) > 0 {
		option = options[0]
	}

	randSeed := option.RandSeed
	if randSeed == 0 {
		randSeed = 99
	}

	return &Deck[C]{
		randSeed: randSeed,
	}
}

func (d *Deck[C]) AddCard(card C) *Deck[C] {
	d.Cards = append(d.Cards, card)
	return d
}

func (d *Deck[C]) DrawCard() (ret C) {
	if d.Cards == nil || len(d.Cards) == 0 {
		return
	}

	// Optimization for single card deck.
	if len(d.Cards) == 1 {
		card := d.Cards[0]
		d.Cards = []C{}
		return card
	}

	r := rand.New(rand.NewSource(d.randSeed)) //nolint:gosec

	cardIdx := r.Intn(len(d.Cards))
	card := d.Cards[cardIdx]

	// Remove the picked card from the deck.
	d.Cards = append(d.Cards[:cardIdx], d.Cards[cardIdx+1:]...)
	return card
}
