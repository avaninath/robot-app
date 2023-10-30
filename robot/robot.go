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

// Initiate creates a new robot with the starting position and returns it
func Initiate() Robot {
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

	robot := Robot{
		Row:       rPosInt,
		Column:    cPosInt,
		Direction: direction,
	}

	return robot
}

// MoveForward moves the robot one step forward in the direction it is facing
func (r *Robot) MoveForward() {
	switch r.Direction {
	case DirectionNorth:
		r.Row--
	case DirectionSouth:
		r.Row++
	case DirectionEast:
		r.Column++
	case DirectionWest:
		r.Column--
	default:
	}
}

// TurnLeft turns the robot once to the left
func (r *Robot) TurnLeft() {
	r.Direction = ToLeft(r.Direction)
}

// TurnRight turns the robot once to the right
func (r *Robot) TurnRight() {
	r.Direction = ToRight(r.Direction)
}

// Report prints the current position and direction of the robot
func (r *Robot) Report() {
	fmt.Println("The robot is at position " + strconv.Itoa(r.Row) + " x " + strconv.Itoa(r.Column) + " facing " + r.Direction.ToString())
}

// ExecuteCommand executes the set of commands given to the robot
func (r *Robot) ExecuteCommand(command string) {
	for _, c := range command {
		switch c {
		case 'L':
			r.TurnLeft()
		case 'R':
			r.TurnRight()
		case 'F':
			r.MoveForward()
		default:
		}
	}
	r.Report()
}
