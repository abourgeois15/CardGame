package game

import (
	"math/rand"
	"time"
)

// Constantes for game
const (
	CardDeckSize  = 52
	NbCardByColor = 13

	NbPlayersMin  = 2
	NbPlayersMax  = 4
	FirstPlayerID = 0

	FirstTurnID = 1
)

// Game is the struct with all the ongoing game info
type Game struct {
	NbPlayers   int
	CardDeck    []*Card
	Players     []*Player
	FirstPlayer *Player
	OngoingTurn *Turn
}

// NewGame creates a new game pointer with players and cards
func NewGame(nbPlayers int) *Game {
	game := new(Game)
	game.createCardDeck()
	game.NbPlayers = nbPlayers
	game.Players = make([]*Player, nbPlayers)
	for index := range game.Players {
		if index == FirstPlayerID {
			game.Players[index] = NewPlayer(index, Human)
		} else {
			game.Players[index] = NewPlayer(index, AI)
			game.Players[index-1].NextPlayer = game.Players[index]
		}
	}
	game.FirstPlayer = game.Players[FirstPlayerID]
	game.Players[nbPlayers-1].NextPlayer = game.Players[FirstPlayerID]
	game.OngoingTurn = NewTurn(FirstTurnID, game.FirstPlayer)
	return game
}

func (game *Game) createCardDeck() {
	game.CardDeck = make([]*Card, CardDeckSize)
	for index := range game.CardDeck {
		game.CardDeck[index] = NewCard(index, CardColor(index/NbCardByColor), CardNumber(index%NbCardByColor))
	}
	game.shuffleCardDeck()
}

func (game *Game) shuffleCardDeck() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(CardDeckSize, func(i, j int) { game.CardDeck[i], game.CardDeck[j] = game.CardDeck[j], game.CardDeck[i] })
}
