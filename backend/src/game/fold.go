package game

import "strconv"

// Fold is a struct to describe the ongoing fold
type Fold struct {
	ID            int
	NbCards       int
	FirstPlayer   *Player
	WinningPlayer *Player
	WinningCard   *Card
	AskedColor    CardColor
}

// NewFold creates a new fold and refolds the pointer
func NewFold(id int, nbCards int, firstPlayer *Player) *Fold {
	fold := new(Fold)
	fold.ID = id
	fold.NbCards = nbCards
	fold.FirstPlayer = firstPlayer
	fold.WinningPlayer = firstPlayer
	return fold
}

func (fold *Fold) String() string {
	return "Fold " + strconv.Itoa(fold.ID) + " (" + strconv.Itoa(fold.NbCards) + ")"
}
