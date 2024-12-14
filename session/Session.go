package session

import (
	"TicTacToe/AI"
)

type Session struct {
	SessionId_ string
	Ai         AI.AI
	GameState_ AI.Field
}

func (s *Session) Init(dif AI.Difficulty, side AI.FieldContainment) {

	if side == AI.Cross {
		s.Ai = *AI.CreateAI(dif, AI.Circle)
	} else {
		s.Ai = *AI.CreateAI(dif, AI.Cross)
	}

	s.GameState_ = make(AI.Field, 3)
	for i := range s.GameState_ {
		s.GameState_[i] = make(AI.FieldRow, 3)
	}

}
