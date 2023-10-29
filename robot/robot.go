package robot

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var in = bufio.NewReader(os.Stdin)

type Robot struct {
	Row       int
	Column    int
	Direction Direction
}

func Initiate() Robot {
	fmt.Println("Please enter the starting position of the robot")
	fmt.Println("Enter the row position:")
	rPos, _ := in.ReadString('\n')
	fmt.Println("Enter the column position:")
	cPos, _ := in.ReadString('\n')
	fmt.Println("Enter the direction (N, S, E, W):")
	stringDirection, _ := in.ReadString('\n')
	direction := ToDirection(strings.TrimSpace(stringDirection))

	fmt.Println("The starting position of the robot is " + strings.TrimSpace(rPos) + " x " + strings.TrimSpace(cPos) + " facing " + direction.ToString())

	// Create a variable of type Robot and save the values for row, column and direction input in that variable
	rPosInt, _ := strconv.Atoi(strings.TrimSpace(rPos))
	cPosInt, _ := strconv.Atoi(strings.TrimSpace(cPos))
	robot := Robot{Row: rPosInt, Column: cPosInt, Direction: direction}

	return robot
}
