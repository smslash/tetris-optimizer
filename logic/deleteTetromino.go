package logic

func DeleteTetromino(board [][]string, boardSize int, letter rune) {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if board[i][j] == string(letter) {
				board[i][j] = "."
			}
		}
	}
}
