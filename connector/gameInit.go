package connector

import (
	"TicTacToe/AI"
	session "TicTacToe/Session"
	"net/http"
	"text/template"
)

func GameInit(w http.ResponseWriter, r *http.Request) {
	game, err := template.ParseFiles("web/Game.html")
	if err != nil {
		http.Error(w, "Ошибка загрузки фронта", http.StatusInternalServerError)
		return
	}
	query := r.URL.Query()

	sessionId := query.Get("sessionId")
	man := session.GetManager()
	thisSession := man.Sessions[sessionId]
	if thisSession.Ai.Side_ == AI.Cross {
		thisSession.Ai.MakeMove(thisSession.GameState_)
	}
	game.Execute(w, nil)
}
