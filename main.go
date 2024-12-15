package main

import (
	"TicTacToe/connector"
	"net/http"
)

func main() {

	http.Handle("/", http.HandlerFunc(connector.MainScreenHandler))
	http.Handle("/start-game", http.HandlerFunc(connector.StartGame))
	http.Handle("/game", http.HandlerFunc(connector.GameInit))
	http.Handle("/game/state", http.HandlerFunc(connector.GameStateHandler))
	http.Handle("/game/play", http.HandlerFunc(connector.Game))

	http.ListenAndServe(":8080", nil)
}
