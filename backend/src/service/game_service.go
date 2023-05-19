package service

import "api/game"

// PlayerService is a player struct to return to the client
type PlayerService struct {
	ID           int `json:"id"`
	Type         int `json:"type"`
	Score        int `json:"score"`
	NextPlayerID int `json:"nextPlayerID"`
}

// PlayerToService is a function that returns the player service struct from a player
func PlayerToService(player *game.Player) PlayerService {
	return PlayerService{
		ID:           player.ID,
		Type:         int(player.Type),
		Score:        player.Score,
		NextPlayerID: player.NextPlayer.ID,
	}
}

// GameService is a game struct ro return to the client
type GameService struct {
	NbPlayers     int             `json:"nbPlayer"`
	Players       []PlayerService `json:"players"`
	FirstPlayerID int             `json:"firstPlayerID"`
}

// GameToService is a function that returns the game service struct from a game
func GameToService(game *game.Game) GameService {
	players := make([]PlayerService, game.NbPlayers)
	for index := range players {
		players[index] = PlayerToService(game.Players[index])
	}
	return GameService{
		NbPlayers:     game.NbPlayers,
		Players:       players,
		FirstPlayerID: game.FirstPlayer.ID,
	}
}
