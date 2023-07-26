package logic

import (
	"log"

	"git/ssengerb/tetris-optimizer/valid"
)

func SolveTetris(data []string, board [][]string, boardSize int, index int, counter int) bool {
	if counter == len(data) {
		return true
	}
	for i := 0; i < len(data); i++ {
		flag := false
		letter := 'A' + rune(i)
		if letter > 'Z' {
			log.Fatalln("The limit of the alphabet has been exceeded")
		}

		for a := 0; a < boardSize; a++ {
			for b := 0; b < boardSize; b++ {
				if board[a][b] == string(letter) {
					flag = true
				}
			}
		}

		if flag {
			continue
		}

		slice, ok := valid.IsValidFormat(data[i])
		if !ok || !valid.IsTetromino(slice) {
			log.Fatal("ERROR\n\nBad Format\n")
		}

		if CanPlaceTetromino(board, slice, boardSize) {
			PlaceTetromino(board, slice, boardSize, letter)
			counter++

			if SolveTetris(data, board, boardSize, i, counter) {
				return true
			}

			DeleteTetromino(board, boardSize, letter)
			counter--
		}
	}
	if counter == len(data) {
		return true
	}
	return false
}
