package game

// Player is a struct for a player
type Player struct {
	ID         int
	Score      int
	Hand       []*Card
	nextPlayer *Player
}
