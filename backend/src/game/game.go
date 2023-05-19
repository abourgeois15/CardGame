package game

const (
	cardDeckSize  = 52
	nbCardByColor = 13
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
	return game
}

func (game *Game) createCardDeck() {
	game.CardDeck = make([]*Card, cardDeckSize)
	for index := range game.CardDeck {
		game.CardDeck[index] = NewCard(index, CardColor(index/nbCardByColor), CardNumber(index%nbCardByColor))
	}
}
