//go:build (go1.18 || go1.19 || go2) && feature && playingcard

package main

type PlayingCard struct {
	Suit string
	Rank string
}

func NewPlayingCard(suit, rank string) *PlayingCard {
	return &PlayingCard{Suit: suit, Rank: rank}
}

func (c *PlayingCard) String() string {
	return c.Rank + " of " + c.Suit
}

func (c *PlayingCard) Name() string {
	return c.String()
}

type PlayingCardDeckOptions struct {
	RandSeed int64
	Suits    []string
	Ranks    []string
}

func NewPlayingCardDeck(options ...PlayingCardDeckOptions) *Deck[*PlayingCard] {
	option := PlayingCardDeckOptions{
		RandSeed: 99,
	}

	if len(options) > 0 {
		option = options[0]
	}

	suits := []string{"Diamonds", "Hearts", "Clubs", "Spades"}
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	if len(option.Suits) > 0 {
		suits = option.Suits
	}

	if len(option.Ranks) > 0 {
		ranks = option.Ranks
	}

	deck := NewDeck[*PlayingCard](DeckOptions{
		RandSeed: option.RandSeed,
	})

	// Iterate over each suit and rank to create a new card.
	// And add the card to the Deck.
	for _, suit := range suits {
		for _, rank := range ranks {
			deck.AddCard(NewPlayingCard(suit, rank))
		}
	}

	return deck
}
