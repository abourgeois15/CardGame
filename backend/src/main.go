package main

import (
	"api/game"
	"api/service"
	"api/system"
	"fmt"
	"log"
	"net/http"
)

const (
	port = "3001"
)

func main() {
	service := service.NewService()
	service.Game = game.NewGame(4)
	for _, card := range service.Game.CardDeck {
		fmt.Println(card.ID, ": ", card, " value ", card.Value)
	}
	for _, player := range service.Game.Players {
		fmt.Println(player, ", score: ", player.Score, ", next player: ", player.NextPlayer)
	}
	fmt.Println(service.Game.OngoingTurn)
	fmt.Println(service.Game.OngoingTurn.OngoingFold)
	r := system.SetupRouter(service)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
