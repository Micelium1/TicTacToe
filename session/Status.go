package Session

import "TicTacToe/AI"

func CheckGameState(field AI.Field) string {
	const size = 3

	// Проверяем строки и столбцы.
	for i := 0; i < size; i++ {
		// Проверка строки.
		if field[i][0] != AI.None && field[i][0] == field[i][1] && field[i][1] == field[i][2] {
			if field[i][0] == AI.Cross {
				return "Крестики выиграли"
			}
			return "Нолики выиграли"
		}

		// Проверка столбца.
		if field[0][i] != AI.None && field[0][i] == field[1][i] && field[1][i] == field[2][i] {
			if field[0][i] == AI.Cross {
				return "Крестики выиграли"
			}
			return "Нолики выиграли"
		}
	}

	// Проверяем диагонали.
	if field[0][0] != AI.None && field[0][0] == field[1][1] && field[1][1] == field[2][2] {
		if field[0][0] == AI.Cross {
			return "Крестики выиграли"
		}
		return "Нолики выиграли"
	}

	if field[0][2] != AI.None && field[0][2] == field[1][1] && field[1][1] == field[2][0] {
		if field[0][2] == AI.Cross {
			return "Крестики выиграли"
		}
		return "Нолики выиграли"
	}

	// Проверяем, остались ли свободные клетки.
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if field[i][j] == AI.None {
				return "Игра продолжается"
			}
		}
	}

	// Если дошли сюда, то свободных клеток нет, и победителя тоже нет — ничья.
	return "Ничья"
}
