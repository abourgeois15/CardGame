package game

import "strconv"

const (
	halfGameTurnID = 12
	nbCardsMax     = 12

	firstFoldID = 1
)

// Turn is a struct to describe the ongoing turn
type Turn struct {
	ID          int
	NbCards     int
	TrumpColor  CardColor
	FirstPlayer *Player
	OngoingFold *Fold
}

// NewTurn creates a new turn and returns the pointer
func NewTurn(id int, firstPlayer *Player) *Turn {
	turn := new(Turn)
	turn.ID = id
	turn.NbCards = IDToNbCards(id)
	turn.FirstPlayer = firstPlayer
	turn.OngoingFold = NewFold(firstFoldID, turn.NbCards, firstPlayer)
	return turn
}

// IDToNbCards returns the number of card for the id of the turn
func IDToNbCards(id int) int {
	if id <= halfGameTurnID {
		return id
	}
	return nbCardsMax + halfGameTurnID + 1 - id
}

func (turn *Turn) String() string {
	return "Turn " + strconv.Itoa(turn.ID) + " (" + strconv.Itoa(turn.NbCards) + ")"
}
