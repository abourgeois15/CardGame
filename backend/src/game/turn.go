package game

// Turn is a struct to describe the ongoing turn
type Turn struct {
	ID          int
	NbCard      int
	TrumpColor  CardColor
	FirstPlayer *Player
	OngoingFold *Fold
}
