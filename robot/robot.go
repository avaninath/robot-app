package robot

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/robotAssignment/board"
)

var in = bufio.NewReader(os.Stdin)

type robot struct {
	Row       int
	Column    int
	Direction Direction
}

// Initiate creates a new robot with the starting position and returns it
func Initiate() robot {
	fmt.Println("Please enter the starting position of the robot")
	fmt.Println("Enter the row position:")
	rPos, _ := in.ReadString('\n')
	fmt.Println("Enter the column position:")
	cPos, _ := in.ReadString('\n')
	fmt.Println("Enter the direction (N, S, E, W):")
	stringDirection, _ := in.ReadString('\n')
	direction := ToDirection(strings.TrimSpace(stringDirection))

	rPosInt, _ := strconv.Atoi(strings.TrimSpace(rPos))
	cPosInt, _ := strconv.Atoi(strings.TrimSpace(cPos))

	robot := robot{
		Row:       rPosInt,
		Column:    cPosInt,
		Direction: direction,
	}

	return robot
}

// MoveForward moves the robot one step forward in the direction it is facing
func (r *robot) MoveForward(board *board.Board) error {
	switch r.Direction {
	case DirectionNorth:
		r.Row--
		if r.Row < 0 || r.Row > board.MaxRows {
			return errors.New("the robot fell off the board")
		}
	case DirectionSouth:
		r.Row++
		if r.Row < 0 || r.Row > board.MaxRows {
			return errors.New("the robot fell off the board")
		}
	case DirectionEast:
		r.Column++
		if r.Column < 0 || r.Column > board.MaxColumns {
			return errors.New("the robot fell off the board")
		}
	case DirectionWest:
		r.Column--
		if r.Column < 0 || r.Column > board.MaxColumns {
			return errors.New("the robot fell off the board")
		}
	}
	return nil
}

// TurnLeft turns the robot once to the left
func (r *robot) TurnLeft() {
	r.Direction = ToLeft(r.Direction)
}

// TurnRight turns the robot once to the right
func (r *robot) TurnRight() {
	r.Direction = ToRight(r.Direction)
}

// Report prints the current position and direction of the robot
func (r *robot) Report() {
	fmt.Println("The robot is at position " + strconv.Itoa(r.Row) + " x " + strconv.Itoa(r.Column) + " facing " + r.Direction.ToString())
}

// ExecuteCommand executes the set of commands given to the robot
func (r *robot) ExecuteCommand(command string, b *board.Board) error {
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
	r.Report()
	return nil
}
