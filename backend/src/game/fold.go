package game

// Fold is a struct to describe the ongoing fold
type Fold struct {
	FirstPlayer   *Player
	WinningPlayer *Player
	WinningCard   *Card
	AskedColor    CardColor
}
