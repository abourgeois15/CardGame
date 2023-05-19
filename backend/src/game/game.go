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

	firstTurnID = 1
)

// Game is the struct with all the ongoing game info
type Game struct {
	NbPlayer    int
	CardDeck    []*Card
	Players     []*Player
	firstPlayer *Player
	OngoingTurn *Turn
}

// NewGame creates a new game pointer with players and cards
func NewGame(nbPlayer int) *Game {
	game := new(Game)
	game.createCardDeck()
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
	game.firstPlayer = game.Players[firstPlayerID]
	game.Players[nbPlayer-1].NextPlayer = game.Players[firstPlayerID]
	game.OngoingTurn = NewTurn(firstTurnID, game.firstPlayer)
	return game
}

func (game *Game) createCardDeck() {
	game.CardDeck = make([]*Card, cardDeckSize)
	for index := range game.CardDeck {
		game.CardDeck[index] = NewCard(index, CardColor(index/nbCardByColor), CardNumber(index%nbCardByColor))
	}
	game.shuffleCardDeck()
}

func (game *Game) shuffleCardDeck() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(cardDeckSize, func(i, j int) { game.CardDeck[i], game.CardDeck[j] = game.CardDeck[j], game.CardDeck[i] })
}
