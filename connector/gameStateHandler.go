package connector

import (
	session "TicTacToe/Session"
	"TicTacToe/model"
	"encoding/json"
	"net/http"
)

func GameStateHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	sessionId := query.Get("sessionId")
	man := session.GetManager()
	thisSession := man.Sessions[sessionId]

	State := model.GameStateToJson{
		Board:  thisSession.GameState_,
		Status: session.CheckGameState(thisSession.GameState_),
	}
	json.NewEncoder(w).Encode(&State)
}
