package main

import (
	"fmt"
	"github.com/deck"
	"strings"
)

type Hand []deck.Card

func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := range h {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ", ")
}

func (h Hand) Score() int {
	minScore := h.MinScore()
	if minScore > 11 {
		return minScore
	}
	for _, c := range h {
		if c.Rank == deck.Ace {
			return minScore + 10
		}
	}
	return minScore
}

func (h Hand) MinScore() int {
	var score = 0
	for _, c := range h {
		score += min(int(c.Rank), 10)
	}
	return score
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	cards := deck.NewDeck(deck.Deck(1), deck.Shuffle)
	var card deck.Card
	var player, dealer Hand
	for i := 0; i < 2; i++ {
		for _, hand := range []*Hand{&player, &dealer} {
			card, cards = draw(cards)
			*hand = append(*hand, card)
		}
	}
	var input string
	for input != "s" {
		fmt.Println("player: ", player)
		fmt.Println("dealer: ", dealer.dealerString())
		fmt.Println("what will you do? (h)it or (s)tand")
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h": {
			card, cards = draw(cards)
			player = append(player, card)
		}
		}
	}
	if dealer.Score() <= 16 || (dealer.Score() == 17 && dealer.MinScore() != 17) {
		card, cards = draw(cards)
		dealer = append(dealer, card)
	}
	pScore, dScore := player.Score(), dealer.Score()
	fmt.Println("-------final hands-------")
	fmt.Println("player: ", player, ", score: ", pScore)
	fmt.Println("dealer: ", dealer, ", score: ", dScore)
	switch {
	case pScore > 21 : {
		fmt.Println("You bursted")
	}
	case dScore > 21: {
		fmt.Println("Dealer bursted")
	}
	case pScore > dScore: {
		fmt.Println("You win!")
	}
	case pScore < dScore: {
		fmt.Println("You lose!")
	}
	case pScore == dScore: {
		fmt.Println("Draw")
	}
	}
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

func (h Hand) dealerString() string {
	return h[0].String() + ", ----hidden----"
}