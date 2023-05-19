package game

import (
	"math/rand"
	"time"
)

const (
	cardDeckSize  = 52
	nbCardByColor = 13

	nbPlayersMin  = 2
	firstPlayerID = 0
)

// Game is the struct with all the ongoing game info
type Game struct {
	NbPlayer    int
	CardDeck    []*Card
	Players     []*Player
	OngoingTurn *Turn
}

// NewGame creates a new game pointer with players and cards
func NewGame() *Game {
	game := new(Game)
	game.createCardDeck()
	game.shuffleCardDeck()
	return game
}

func (game *Game) createCardDeck() {
	game.CardDeck = make([]*Card, cardDeckSize)
	for index := range game.CardDeck {
		game.CardDeck[index] = NewCard(index, CardColor(index/nbCardByColor), CardNumber(index%nbCardByColor))
	}
}

func (game *Game) shuffleCardDeck() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(cardDeckSize, func(i, j int) { game.CardDeck[i], game.CardDeck[j] = game.CardDeck[j], game.CardDeck[i] })
}

// CreateGame will create the players and ongoing turn of the game
func (game *Game) CreateGame(nbPlayer int) {
	game.NbPlayer = nbPlayer
	game.Players = make([]*Player, nbPlayer)
	for index := range game.Players {
		if index == firstPlayerID {
			game.Players[index] = NewPlayer(index, Human)
		} else {
			game.Players[index] = NewPlayer(index, AI)
			game.Players[index-1].NextPlayer = game.Players[index]
		}
	}
	game.Players[nbPlayer-1].NextPlayer = game.Players[firstPlayerID]
}
