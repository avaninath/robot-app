package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/robotAssignment/board"
	"github.com/robotAssignment/robot"
)

var in = bufio.NewReader(os.Stdin)

func Initiate() {
	fmt.Println("Welcome to the Robot Game")
	fmt.Println(strings.Repeat("-", 25))

loop:
	for {
		fmt.Println("Please select an option")
		fmt.Println("1) Start a new game")
		fmt.Println("q) quit")
		choice, _ := in.ReadString('\n')

		switch strings.TrimSpace(choice) {
		case "1":
			// Create a new board
			newBoard := board.Initiate()
			fmt.Printf("\nThe size of the board is %v x %v \n", newBoard.MaxRows, newBoard.MaxColumns)

			// Create a new robot and assign initial position
			newRbt := robot.Initiate()
			fmt.Printf("\nThe initial position of the robot is %v, %v facing %v \n", newRbt.Row, newRbt.Column, newRbt.Direction)

			// Get the commands from the user
			for {
				fmt.Println("\nPlease enter the commands for the robot: (eg: FFRFFLFRRF)")
				inputCommand, _ := in.ReadString('\n')
				command := strings.TrimSpace(inputCommand)

				fmt.Println("The commands for the robot are " + command)
				err := newRbt.ExecuteCommand(command, &newBoard)
				if err != nil {
					fmt.Println(err)
					break
				}
			}
		case "q":
			break loop
		default:
			fmt.Println("Unknown option")
		}
	}
}
