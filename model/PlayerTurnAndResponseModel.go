package model

import "TicTacToe/AI"

type PlayRequest struct {
	ID  string `json:"id"`
	Row int    `json:"row"`
	Col int    `json:"col"`
}

type PlayResponse struct {
	Board   AI.Field `json:"board"`
	Status  string   `json:"status"`
	EndGame bool     `json:"endGame"`
}
