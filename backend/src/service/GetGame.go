package service

import (
	"net/http"
)

// GetGame init the game and returns the struct
func (s *Service) GetGame(w http.ResponseWriter, r *http.Request) {
	content := &Content[GameService]{
		Details: "Success",
	}
	defer content.Encode(w)

	//Return Game
	content.Payload = GameToService(s.Game)
	w.WriteHeader(http.StatusOK)
}
