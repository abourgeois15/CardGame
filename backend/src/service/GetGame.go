package service

import (
	"api/game"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// GetGame init the game and returns the struct
func (s *Service) GetGame(w http.ResponseWriter, r *http.Request) {
	content := &Content[GameService]{
		Details: "Success",
	}
	defer content.Encode(w)

	// Extracting data
	var nbPlayers int
	nbStr := chi.URLParam(r, "NbPlayers")
	if _, err := fmt.Sscanf(nbStr, "%d", &nbPlayers); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		content.Details = "Couldn't read URL parameter"
		fmt.Println(content.Details, err)
		return
	}

	//Create Game
	s.Game = game.NewGame(nbPlayers)
	content.Payload = GameToService(s.Game)
	w.WriteHeader(http.StatusOK)
}
