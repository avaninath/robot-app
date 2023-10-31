package robot

import (
	"bufio"
	"errors"
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
			fmt.Println(ErrInvalidInputDirection)
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
func (r *Robot) MoveForward(board *board.Board) error {
	switch r.Direction {
	case DirectionNorth:
		r.Row--
		return outOfBoundError(r.Row, r.Column, board)
	case DirectionSouth:
		r.Row++
		return outOfBoundError(r.Row, r.Column, board)
	case DirectionEast:
		r.Column++
		return outOfBoundError(r.Row, r.Column, board)
	case DirectionWest:
		r.Column--
		return outOfBoundError(r.Row, r.Column, board)
	}
	return nil
}

// TurnLeft turns the robot once to the left
func (r *Robot) TurnLeft() {
	r.Direction = ToLeft(r.Direction)
}

// TurnRight turns the robot once to the right
func (r *Robot) TurnRight() {
	r.Direction = ToRight(r.Direction)
}

// Report prints the current position and direction of the robot and issues a warning if it is close to the edge or a corner
func (r *Robot) Report(b *board.Board) {
	// Check if robot is close to edge or corner and issue warning message
	corner := r.checkCornerPosition(b)
	if corner != "" {
		dangerDirections := getCornerDirections(corner)
		fmt.Printf("\n%s Danger in directions: %v ", warningMessage(), dangerDirections)
	} else if isCloseToEdge(*r, b) {
		fmt.Printf("\n%s Danger in direction: %v ", warningMessage(), r.Direction.ToString())
	}

	// Report back the position and direction of the robot
	fmt.Printf("\nThe robot is at position %v x %v facing %v", strconv.Itoa(r.Row), strconv.Itoa(r.Column), r.Direction.ToString())
}

// ExecuteCommand executes the set of commands given to the robot
func (r *Robot) ExecuteCommand(command string, b *board.Board) error {
	err := validateCommandInput(command)
	if err != nil {
		return err
	}

	command = strings.ToUpper(command)
	fmt.Printf("\nExecuting the commands %v ...", command)

	for _, c := range command {
		if c == 'L' {
			r.TurnLeft()
		} else if c == 'R' {
			r.TurnRight()
		} else if c == 'F' {
			err := r.MoveForward(b)
			if err != nil {
				return err
			}
		}
	}
	r.Report(b)
	return nil
}

func isCloseToEdge(rbt Robot, b *board.Board) bool {
	newRbt := Robot{
		Row:       rbt.Row,
		Column:    rbt.Column,
		Direction: rbt.Direction,
	}
	err := newRbt.MoveForward(b)
	fmt.Printf("\n%v", err)
	return errors.Is(err, ErrRobotFellOffBoard)
}

func (r *Robot) checkCornerPosition(b *board.Board) Corner {
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

func outOfBoundError(row, column int, b *board.Board) error {
	if row < 0 || row > b.MaxRows || column < 0 || column > b.MaxColumns {
		return fmt.Errorf("%w", ErrRobotFellOffBoard)
	}
	return nil
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
