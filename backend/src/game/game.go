package game

// Game is the struct with all the ongoing game info
type Game struct {
	NbPlayer    int
	Players     []*Player
	OngoingTurn Turn
}
