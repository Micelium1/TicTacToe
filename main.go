package main

import (
	"TicTacToe/connector"
	"net/http"
)

func main() {

	http.Handle("/", http.HandlerFunc(connector.MainScreenHandler))
	http.Handle("/start-game", http.HandlerFunc(connector.StartGame))

	http.ListenAndServe(":8080", nil)
}
