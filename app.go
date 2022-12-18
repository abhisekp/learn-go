package main

import (
	"fmt"
	"time"
)

func main() {
	playingCardDeck := NewPlayingCardDeck(PlayingCardDeckOptions{
		RandSeed: time.Now().UnixNano(),
	})

	tradingCardDeck := NewTradingCardDeck(TradingCardDeckOptions{
		RandSeed: time.Now().UnixNano(),
	})

	fmt.Println("--- drawing playing card ---")
	playingCard := playingCardDeck.DrawCard()
	fmt.Println("Drew playing card:", playingCard)

	/*
	   playingCard, ok := card.(*PlayingCard)
	   	if !ok {
	   		fmt.Println("Not a playing card")
	   		return
	   	} else {
	   		fmt.Println("Playing card:", playingCard)
	   	}
	*/

	fmt.Println("--- drawing trading card ---")
	tradingCard := tradingCardDeck.DrawCard()
	fmt.Println("Drew trading card:", tradingCard)

	fmt.Println("--- printing card ---")
	printCard[*PlayingCard](playingCard)
	printCard(tradingCard)
}

func printCard[C Card](card C) {
	fmt.Println("Card name:", card.Name())
}
