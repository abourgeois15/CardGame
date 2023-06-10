package system

import (
	"api/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

// SetupRouter initializes the router's routes and returns it
func SetupRouter(s *service.Service) chi.Router {
	var r chi.Router = chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST"},
		ExposedHeaders: []string{"Content-Types"},
	}))
	// Setting up routes
	r.Route("/game", func(r chi.Router) {
		r.Post("/{NbPlayers}", s.CreateGame)
		r.Get("/", s.GetGame)
	})
	return r
}
