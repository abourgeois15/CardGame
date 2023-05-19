package game

import "strconv"

// PlayerType is a type for player types
type PlayerType int

// Enum of player types
const (
	Human PlayerType = iota
	AI
)

func (playerType PlayerType) String() string {
	switch playerType {
	case Human:
		return "Human"
	case AI:
		return "AI"
	default:
		return "Unknown"
	}
}

// Player is a struct for a player
type Player struct {
	ID         int
	Type       PlayerType
	Score      int
	Hand       []*Card
	NextPlayer *Player
}

func (player *Player) String() string {
	return "Player " + strconv.Itoa(player.ID) + " (" + player.Type.String() + ")"
}

// NewPlayer creates a new player
func NewPlayer(id int, pType PlayerType) *Player {
	player := new(Player)
	player.ID = id
	player.Type = pType
	player.Score = 0
	return player
}
