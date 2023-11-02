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
				newRobotPosition, err := robot.ExecuteCommand(inputCommand, &newRbt, &newBoard)
				if errors.Is(err, robot.ErrInvalidCommandInput) {
					fmt.Println(invalidCommandInputMessageString)
					continue
				} else if checkRobotPosition(newRobotPosition, &newBoard) != nil {
					fmt.Println("\nOops! The robot fell off the board. Better luck next time!")
					fmt.Printf("\nThe final position of the robot was %v, %v facing %v \n", newRobotPosition.Row, newRobotPosition.Column, newRobotPosition.Direction)
					break
				}
				robot.IssueWarning(newRobotPosition, &newBoard)
				fmt.Printf("\n\nREPORT:\nThe final position of the robot is %v, %v facing %v \n", newRobotPosition.Row, newRobotPosition.Column, newRobotPosition.Direction)

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

func checkRobotPosition(rbt *robot.Robot, brd *board.Board) error {
	if rbt.Row < 0 || rbt.Row >= brd.MaxRows || rbt.Column < 0 || rbt.Column >= brd.MaxColumns {
		return robot.ErrRobotFellOffBoard
	}
	return nil
}
