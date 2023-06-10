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

type CreateGameTestSuite struct {
	suite.Suite
	Handler http.Handler
	service *service.Service
}

func TestCreateEmployeeTestSuite(t *testing.T) {
	suite.Run(t, new(CreateGameTestSuite))
}

func (s *CreateGameTestSuite) SetupTest() {
	s.service = service.NewService()
	s.Handler = system.SetupRouter(s.service)
	go http.ListenAndServe(":3001", s.Handler)
}

func (s *CreateGameTestSuite) TestCreateGame() {
	var err error
	nbPlayers := game.NbPlayersMin + rand.Intn(game.NbPlayersMax-game.NbPlayersMin)

	// Create request
	var request *http.Request
	if request, err = http.NewRequest("POST", "/game/"+strconv.Itoa(nbPlayers), nil); !s.Nil(err) {
		s.FailNow(err.Error())
	}
	recorder := httptest.NewRecorder()
	s.T().Log("Serve endpoints")
	s.Handler.ServeHTTP(recorder, request)
	s.Equal(http.StatusOK, recorder.Code)
	s.Equal(nbPlayers, s.service.Game.NbPlayers)

	// Check the body
	var content service.Content[struct{}]
	if err := json.NewDecoder(recorder.Body).Decode(&content); err != nil {
		s.FailNow(err.Error())
	}
	s.T().Log(content.Payload)
	s.Equal("Success", content.Details)
}
