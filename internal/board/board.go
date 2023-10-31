package board

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var in = bufio.NewReader(os.Stdin)

// Board is a struct that holds the size of the board
type Board struct {
	MaxRows    int
	MaxColumns int
}

// Initiate creates a new board with the size and returns it
func Initiate() Board {
	var (
		rInt, cInt int
		err        error
	)
	fmt.Println("\nLet's create a new board!")

	for {
		fmt.Println("\nEnter the number of rows of the room:")
		r, _ := in.ReadString('\n')
		rInt, err = strconv.Atoi(strings.TrimSpace(r))
		if err == nil {
			break
		}
		fmt.Println(ErrInvalidInputRows)
	}

	for {
		fmt.Println("\nEnter the number of columns of the room:")
		c, _ := in.ReadString('\n')
		cInt, err = strconv.Atoi(strings.TrimSpace(c))
		if err == nil {
			break
		}
		fmt.Println(ErrInvalidInputColumns)
	}

	board := Board{
		MaxRows:    rInt,
		MaxColumns: cInt,
	}

	return board
}
