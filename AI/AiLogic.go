package AI

import (
	"math/rand"
)

type Difficulty int

const (
	Easy Difficulty = iota
	Medium
	Hard
)

type FieldContainment int

const (
	None FieldContainment = iota
	Cross
	Circle
)

type Field [][]FieldContainment
type FieldRow []FieldContainment
type AI struct {
	Difficulty_ Difficulty
	Side_       FieldContainment
	MakeMove    func(GameState Field)
}

func CreateAI(Dif Difficulty, Side FieldContainment) *AI {
	var Ai AI
	Ai.Difficulty_ = Dif
	Ai.Side_ = Side
	switch Ai.Difficulty_ {
	case Easy:
		Ai.MakeMove = func(GameState Field) {
			possibleMoves := make([]uint, 0)
			for i, row := range GameState {
				for j, field := range row {

					if field == None {
						possibleMoves = append(possibleMoves, uint(j+3*i))
					}
				}
			}
			choice := possibleMoves[rand.Intn(len(possibleMoves))]
			GameState[choice/3][choice%3] = Ai.Side_
		}

	}
	return &Ai
}
