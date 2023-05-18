package system

import (
	"api/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

// SetupRouter initializes the router's routes and returns it
func SetupRouter(service *service.Service) chi.Router {
	var r chi.Router = chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST"},
		ExposedHeaders: []string{"Content-Types"},
	}))
	// Setting up routes
	return r

}
