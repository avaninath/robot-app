package robot

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/robotAssignment/internal/board"
)

var in = bufio.NewReader(os.Stdin)

type Robot struct {
	Row       int
	Column    int
	Direction Direction
}

// Initiate creates a new robot with the starting position and returns it
func Initiate(board *board.Board) Robot {
	var (
		rPosInt, cPosInt int
		direction        Direction
		err              error
	)

	fmt.Println("\nLet's create a new robot and assign it a position on the board!")

	// Get row position
	for {
		fmt.Println("\nEnter the row position:")
		rPos, _ := in.ReadString('\n')
		rPosInt, err = strconv.Atoi(strings.TrimSpace(rPos))
		if err != nil {
			fmt.Println(ErrInvalidInputRows)
			continue
		}
		if rPosInt < 0 || rPosInt > board.MaxRows {
			fmt.Println(ErrOutOfBoundInputRows)
			continue
		}
		break
	}

	// Get column position
	for {
		fmt.Println("\nEnter the column position:")
		cPos, _ := in.ReadString('\n')
		cPosInt, err = strconv.Atoi(strings.TrimSpace(cPos))
		if err != nil {
			fmt.Println(ErrInvalidInputColumns)
			continue
		}
		if cPosInt < 0 || cPosInt > board.MaxColumns {
			fmt.Println(ErrOutOfBoundInputColumns)
			continue
		}
		break
	}

	// Get direction
	for {
		fmt.Println("\nEnter the direction (N, S, E, W):")
		stringDirection, _ := in.ReadString('\n')
		direction = ToDirection(strings.TrimSpace(strings.ToUpper(stringDirection)))
		if !direction.IsValid() {
			fmt.Println(invalidDirectionMessageString)
			continue
		}
		break
	}

	newRobot := Robot{
		Row:       rPosInt,
		Column:    cPosInt,
		Direction: direction,
	}
	return newRobot
}

// MoveForward moves the robot one step forward in the direction it is facing
func MoveForward(r *Robot, board *board.Board) {
	switch r.Direction {
	case DirectionNorth:
		r.Row--
	case DirectionSouth:
		r.Row++
	case DirectionEast:
		r.Column++
	case DirectionWest:
		r.Column--
	}
}

// Report prints the current position and direction of the robot and issues a warning if it is close to the edge or a corner
func IssueWarning(r *Robot, b *board.Board) {
	// Check if robot is close to edge or corner and issue warning message
	corner := checkCornerPosition(r, b)
	if corner != "" {
		dangerDirections := getCornerDirections(corner)
		fmt.Printf("\n%s Danger in directions: %v ", warningMessage(), dangerDirections)
	} else if isCloseToEdge(*r, b) {
		fmt.Printf("\n%s Danger in direction: %v ", warningMessage(), r.Direction.ToString())
	}
}

// ExecuteCommand executes the set of commands given to the robot
func ExecuteCommand(command string, r *Robot, b *board.Board) (*Robot, error) {
	err := validateCommandInput(command)
	if err != nil {
		return nil, err
	}

	command = strings.ToUpper(command)
	fmt.Printf("\nExecuting the commands %v ...", command)

	for _, c := range command {
		if c == 'L' {
			r.Direction = ToLeft(r.Direction)
		} else if c == 'R' {
			r.Direction = ToRight(r.Direction)
		} else if c == 'F' {
			MoveForward(r, b)
		}
	}
	// issueWarning(r, b)
	return r, nil
}

func isCloseToEdge(rbt Robot, b *board.Board) bool {
	newRbt := Robot{
		Row:       rbt.Row,
		Column:    rbt.Column,
		Direction: rbt.Direction,
	}
	MoveForward(&newRbt, b)
	if newRbt.Row < 0 || newRbt.Row >= b.MaxRows || newRbt.Column < 0 || newRbt.Column >= b.MaxColumns {
		return true
	}
	return false

}

func checkCornerPosition(r *Robot, b *board.Board) Corner {
	cornerCoordinates := getBoardCornerCoordinates(b)
	for boardCorner, boardCorners := range cornerCoordinates {
		for _, coordinates := range boardCorners {
			if r.Row == coordinates.row && r.Column == coordinates.column {
				return boardCorner
			}
		}
	}
	return ""
}

func validateCommandInput(command string) error {
	allowedChars := "LRF"
	command = strings.ToUpper(command)
	for _, char := range command {
		if !strings.Contains(allowedChars, string(char)) {
			return ErrInvalidCommandInput
		}
	}
	return nil
}
