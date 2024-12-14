package connector

import (
	"TicTacToe/AI"
	"TicTacToe/model"
	"TicTacToe/session"
	"encoding/json"
	"fmt"
	"net/http"
)

func StartGame(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Неподдерживаемый метод", http.StatusMethodNotAllowed)
		return
	}
	var req model.StartRequestModel
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Неправильный запрос", http.StatusBadRequest)
		return
	}
	var dif AI.Difficulty
	switch req.Difficulty {
	case "easy":
		dif = AI.Easy
	case "medium":
		dif = AI.Medium
	case "hard":
		dif = AI.Hard
	}
	var side AI.FieldContainment
	switch req.Side {
	case "X":
		side = AI.Cross
	case "O":
		side = AI.Circle
	}

	man := session.GetManager()
	ThisSessionId := man.GenerateUniqueID()
	man.Sessions[ThisSessionId].Init(dif, side)
	resp := model.StartResponseModel{SessionId: ThisSessionId}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(resp)
	fmt.Printf("Началась игра: сложность=%s, сторона=%s, сессия=%s\n", req.Difficulty, req.Side, ThisSessionId)
}
