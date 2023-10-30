package board

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var in = bufio.NewReader(os.Stdin)

type Board struct {
	MaxRows    int
	MaxColumns int
}

// Initiate creates a new board with the size and returns it
func Initiate() Board {
	fmt.Println("Please enter the size of the board")
	fmt.Println("Enter the number of rows of the room:")
	r, _ := in.ReadString('\n')
	fmt.Println("Enter the number of columns of the room:")
	c, _ := in.ReadString('\n')

	rInt, _ := strconv.Atoi(strings.TrimSpace(r))
	cInt, _ := strconv.Atoi(strings.TrimSpace(c))

	board := Board{
		MaxRows:    rInt,
		MaxColumns: cInt,
	}

	return board
}
