package integration

import (
	"api/game"
	"api/service"
	"api/system"
	"encoding/json"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/suite"
)

type GetGameTestSuite struct {
	suite.Suite
	Handler http.Handler
	service *service.Service
}

func TestGetEmployeeTestSuite(t *testing.T) {
	suite.Run(t, new(GetGameTestSuite))
}

func (s *GetGameTestSuite) SetupTest() {
	s.service = service.NewService()
	s.Handler = system.SetupRouter(s.service)
	go http.ListenAndServe(":3001", s.Handler)
}

func (s *GetGameTestSuite) TestGetGame() {
	var err error
	nbPlayers := game.NbPlayersMin + rand.Intn(game.NbPlayersMax-game.NbPlayersMin)

	// Create request
	var request *http.Request
	if request, err = http.NewRequest("GET", "/game/"+strconv.Itoa(nbPlayers), nil); !s.Nil(err) {
		s.FailNow(err.Error())
	}
	recorder := httptest.NewRecorder()
	s.T().Log("Serve endpoints")
	s.Handler.ServeHTTP(recorder, request)
	s.Equal(http.StatusOK, recorder.Code)
	s.Equal(nbPlayers, s.service.Game.NbPlayers)

	// Check the body
	var content service.Content[service.GameService]
	if err := json.NewDecoder(recorder.Body).Decode(&content); err != nil {
		s.FailNow(err.Error())
	}
	s.T().Log(content.Payload)
	expectedGame := s.service.Game
	actualGame := content.Payload
	s.Equal(expectedGame.NbPlayers, actualGame.NbPlayers)
	s.Equal(expectedGame.FirstPlayer.ID, actualGame.FirstPlayerID)
	s.Equal(len(expectedGame.Players), len(actualGame.Players))
	for index, player := range expectedGame.Players {
		s.Equal(player.ID, actualGame.Players[index].ID)
		s.Equal(player.Type, game.PlayerType(actualGame.Players[index].Type))
		s.Equal(player.Score, actualGame.Players[index].Score)
		s.Equal(player.NextPlayer.ID, actualGame.Players[index].NextPlayerID)
	}
}
