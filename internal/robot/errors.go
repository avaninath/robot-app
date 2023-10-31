package robot

import "errors"

var (
	ErrRobotFellOffBoard      = errors.New("robot fell off the board")
	ErrInvalidInputRows       = errors.New("robot: invalid input for rows. Please try again")
	ErrInvalidInputColumns    = errors.New("robot: invalid input for columns. Please try again")
	ErrOutOfBoundInputRows    = errors.New("robot: out of bound input for rows. Please try again")
	ErrOutOfBoundInputColumns = errors.New("robot: out of bound input for columns. Please try again")
	ErrInvalidInputDirection  = errors.New("robot: invalid input for direction. Please try again")
	ErrInvalidCommandInput    = errors.New("invalid command input, allowed characters are LRF")
)
