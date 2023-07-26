package logic

func PlaceTetromino(board, slice [][]string, boardSize int, letter rune) {
	xx := []int{0, 0}
	xy := []int{0, 0}
	yx := []int{0, 0}
	yy := []int{0, 0}
	counter := 0
	i_min := 4
	j_min := 4
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if slice[i][j] == "#" {
				if counter == 0 {
					xx = []int{i, j}
					if i < i_min {
						i_min = i
					}
					if j < j_min {
						j_min = j
					}
				} else if counter == 1 {
					xy = []int{i, j}
					if i < i_min {
						i_min = i
					}
					if j < j_min {
						j_min = j
					}
				} else if counter == 2 {
					yx = []int{i, j}
					if i < i_min {
						i_min = i
					}
					if j < j_min {
						j_min = j
					}
				} else if counter == 3 {
					yy = []int{i, j}
					if i < i_min {
						i_min = i
					}
					if j < j_min {
						j_min = j
					}
				}
				counter++
			}
		}
	}
	if xx[0]-i_min >= 0 {
		xx[0] = xx[0] - i_min
	}
	if xx[1]-j_min >= 0 {
		xx[1] = xx[1] - j_min
	}
	if xy[0]-i_min >= 0 {
		xy[0] = xy[0] - i_min
	}
	if xy[1]-j_min >= 0 {
		xy[1] = xy[1] - j_min
	}
	if yx[0]-i_min >= 0 {
		yx[0] = yx[0] - i_min
	}
	if yx[1]-j_min >= 0 {
		yx[1] = yx[1] - j_min
	}
	if yy[0]-i_min >= 0 {
		yy[0] = yy[0] - i_min
	}
	if yy[1]-j_min >= 0 {
		yy[1] = yy[1] - j_min
	}
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if board[i][j] == "." {
				if i+xx[0] < boardSize && j+xx[1] < boardSize && board[i+xx[0]][j+xx[1]] == "." && i+xy[0] < boardSize && j+xy[1] < boardSize && board[i+xy[0]][j+xy[1]] == "." && i+yx[0] < boardSize && j+yx[1] < boardSize && board[i+yx[0]][j+yx[1]] == "." && i+yy[0] < boardSize && j+yy[1] < boardSize && board[i+yy[0]][j+yy[1]] == "." {
					board[i+xx[0]][j+xx[1]] = string(letter)
					board[i+xy[0]][j+xy[1]] = string(letter)
					board[i+yx[0]][j+yx[1]] = string(letter)
					board[i+yy[0]][j+yy[1]] = string(letter)
					return
				}
			}
		}
	}
}
