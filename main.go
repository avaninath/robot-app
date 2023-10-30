package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/robotAssignment/board"
	"github.com/robotAssignment/robot"
)

var in = bufio.NewReader(os.Stdin)

func main() {
	fmt.Println("Welcome to the Robot Game")
	fmt.Println(strings.Repeat("-", 25))

	// Create a new board
	newBoard := board.Initiate()
	fmt.Printf("\nThe size of the board is %v x %v \n", newBoard.Rows, newBoard.Columns)

	// Create a new robot and assign initial position
	newRbt := robot.Initiate()
	fmt.Printf("\nThe initial position of the robot is %v, %v facing %v \n", newRbt.Row, newRbt.Column, newRbt.Direction)

	fmt.Println("Please enter the commands for the robot: (eg: FFRFFLFRRF)")
	inputCommand, _ := in.ReadString('\n')
	command := strings.TrimSpace(inputCommand)

	fmt.Println("The commands for the robot are " + command)

	newRbt.ExecuteCommand(command)
}
