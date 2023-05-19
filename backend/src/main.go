package main

import (
	"api/service"
	"fmt"
)

const (
	port = "3001"
)

func main() {
	service := service.NewService()
	for _, card := range service.Game.CardDeck {
		fmt.Println(card.ID, ": ", card, " value ", card.Value)
	}
	service.Game.CreateGame(4)
	for _, player := range service.Game.Players {
		fmt.Println(player, ", score: ", player.Score, ", next player: ", player.NextPlayer)
	}
	// r := system.SetupRouter(service)
	// if err := http.ListenAndServe(":"+port, r); err != nil {
	// 	log.Fatal(err)
	// }
}
