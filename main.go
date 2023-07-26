package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"git/ssengerb/tetris-optimizer/logic"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Print("Error: Incorrect number of arguments\n\nUsage: go run . [FILE]\n\nExample: go run . <fileName.txt>\n")
		return
	}

	file := os.Args[1]

	if len(file) < 4 {
		fmt.Print("Error: <fileName.txt> is short\n\nUsage: go run . [FILE]\n\nExample: go run . <fileName.txt>\n")
		return
	} else if len(file) == 4 && file == ".txt" {
		fmt.Print("Error: <fileName.txt> must have a name\n\nUsage: go run . [FILE]\n\nExample: go run . <fileName.txt>\n")
		return
	} else if file[len(file)-4:] != ".txt" {
		fmt.Print("Error: <fileName.txt> must be <.txt> format\n\nUsage: go run . [FILE]\n\nExample: go run . <fileName.txt>\n")
		return
	}

	data0, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	data0 = append(data0, 10)
	data1 := strings.ReplaceAll(string(data0), "\n\n", " ")
	data2 := strings.ReplaceAll(data1, "\n", "")
	data := strings.Fields(data2)

	for boardSize := logic.SmallestSquare(len(data)); boardSize <= 11; boardSize++ {
		board := logic.CreateBoard(boardSize)
		counter := 0
		for i := 0; i < len(data); i++ {
			if logic.SolveTetris(data, board, boardSize, i, counter) {
				for a := 0; a < boardSize; a++ {
					for b := 0; b < boardSize; b++ {
						fmt.Print(board[a][b])
					}
					fmt.Println()
				}
				return
			}
		}
	}
}
