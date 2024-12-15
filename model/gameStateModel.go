package model

import "TicTacToe/AI"

type GameStateToJson struct {
	Board  AI.Field `json:"board"`
	Status string   `json:"status"`
}
