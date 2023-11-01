package robot

import "errors"

var (
	ErrRobotFellOffBoard      = errors.New("robot fell off the board")
	ErrInvalidInputRows       = errors.New("robot: invalid input for rows. Please try again with a number")
	ErrInvalidInputColumns    = errors.New("robot: invalid input for columns. Please try again with a number")
	ErrOutOfBoundInputRows    = errors.New("robot: out of bound input for rows. The initial position cannot be outside the board")
	ErrOutOfBoundInputColumns = errors.New("robot: out of bound input for columns. The initial position cannot be outside the board")
	ErrInvalidCommandInput    = errors.New("invalid command input, allowed characters are LRF")
)
