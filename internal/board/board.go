package board

import (
	"bufio"
	"fmt"
	"os"
)

var in = bufio.NewReader(os.Stdin)

// Board is a struct that holds the size of the board
type Board struct {
	MaxRows    int
	MaxColumns int
}

// Initiate is the entry point of the board.
func Initiate() Board {
	rows, columns, _ := readInput()
	newBoard, _ := createBoard(rows, columns)
	return newBoard
}

func readInput() (int, int, error) {
	var (
		rInt, cInt int
		err        error
	)
	fmt.Println("\nLet's create a new board!")

	for {
		fmt.Println("\nEnter the number of rows of the room:")
		r, _ := in.ReadString('\n')
		rInt, err = validateUserInput(r)
		if err == nil {
			break
		}
		fmt.Println(err)
	}

	for {
		fmt.Println("\nEnter the number of columns of the room:")
		c, _ := in.ReadString('\n')
		cInt, err = validateUserInput(c)
		if err == nil {
			break
		}
		fmt.Println(err)
	}
	return rInt, cInt, nil
}

// createBoard creates a new board with the given size and returns it
func createBoard(rows int, columns int) (Board, error) {
	if rows <= 0 || columns <= 0 {
		return Board{}, ErrInvalidInput
	}
	board := Board{
		MaxRows:    rows,
		MaxColumns: columns,
	}
	return board, nil
}
