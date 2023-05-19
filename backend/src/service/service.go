package service

import "api/game"

// Service is the structure that execute all handle function
type Service struct {
	Game *game.Game
}

// NewService initializes the Service
func NewService() *Service {
	return &Service{}
}
