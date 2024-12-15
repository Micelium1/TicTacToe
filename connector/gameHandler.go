package connector

import (
	"TicTacToe/Session"
	"TicTacToe/model"
	"encoding/json"
	"net/http"
)

func Game(w http.ResponseWriter, r *http.Request) {
	var req model.PlayRequest
	json.NewDecoder(r.Body).Decode(&req)
	man := Session.GetManager()
	thisSession := man.Sessions[req.ID]
	thisSession.GameState_[req.Row][req.Col] = thisSession.Side

	CurrentState := Session.CheckGameState(thisSession.GameState_)
	var resp model.PlayResponse
	if CurrentState != "Игра продолжается" {
		resp = model.PlayResponse{
			Board:   thisSession.GameState_,
			Status:  CurrentState,
			EndGame: true,
		}
		json.NewEncoder(w).Encode(resp)
		return
	}
	thisSession.Ai.MakeMove(thisSession.GameState_)
	CurrentState = Session.CheckGameState(thisSession.GameState_)
	resp = model.PlayResponse{
		Board:   thisSession.GameState_,
		Status:  CurrentState,
		EndGame: CurrentState != "Игра продолжается",
	}
	json.NewEncoder(w).Encode(resp)
}
