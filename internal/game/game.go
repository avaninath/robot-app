package game

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/robotAssignment/internal/board"
	"github.com/robotAssignment/internal/robot"
)

var in = bufio.NewReader(os.Stdin)

/*
Initiate is the entry point of the game.
It displays the options available for the player and takes the input
*/
func Initiate() {
	fmt.Println("\nWelcome to the Robot Game")
	fmt.Println(strings.Repeat("-", 25))

gameLoop:
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
			newRbt := robot.Initiate(&newBoard)
			fmt.Printf("\nThe initial position of the robot is %v, %v facing %v \n", newRbt.Row, newRbt.Column, newRbt.Direction)

			// Get the commands from the user
		commandLoop:
			for {
				fmt.Println("\nPlease enter the commands for the robot: (eg: FFRFFLFRRF)")
				inputCommand, _ := in.ReadString('\n')
				inputCommand = strings.TrimSpace(inputCommand)

				// Validates and execute the input commands
				err := newRbt.ExecuteCommand(inputCommand, &newBoard)
				if errors.Is(err, robot.ErrRobotFellOffBoard) {
					fmt.Println("\nOops! The robot fell off the board. Better luck next time!")
					break
				} else if errors.Is(err, robot.ErrInvalidCommandInput) {
					fmt.Println(invalidCommandInputMessageString)
					continue
				}

				for {
					fmt.Println("\nWould you like to enter more commands for the robot? (y/n)")
					input, _ := in.ReadString('\n')
					input = strings.TrimSpace(input)
					if input != "y" && input != "n" {
						fmt.Println("Invalid input. Please try again")
						continue
					}
					if input == "y" {
						continue commandLoop
					} else {
						break commandLoop
					}
				}
			}
		case "q":
			break gameLoop
		default:
			fmt.Println(unknownGameOptionString)
		}
	}
}
