//go:build (go1.18 || go1.19 || go2) && feature && tradingcard

package main

type TradingCard struct {
	CollectableName string
}

func NewTradingCard(collectableName string) *TradingCard {
	return &TradingCard{CollectableName: collectableName}
}

func (c *TradingCard) String() string {
	return c.CollectableName
}

func (c *TradingCard) Name() string {
	return c.String()
}

type TradingCardDeckOptions struct {
	RandSeed     int64
	Collectables []string
}

func NewTradingCardDeck(options ...TradingCardDeckOptions) *Deck[*TradingCard] {
	option := TradingCardDeckOptions{
		RandSeed: 99,
	}

	if len(options) > 0 {
		option = options[0]
	}

	deck := NewDeck[*TradingCard](DeckOptions{
		RandSeed: option.RandSeed,
	})

	collectables := []string{"Sammy", "Droplets", "Spaces", "App Platform"}
	if len(option.Collectables) > 0 {
		collectables = option.Collectables
	}

	for _, collectable := range collectables {
		deck.AddCard(NewTradingCard(collectable))
	}

	return deck
}
