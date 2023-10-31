package board

import "errors"

var (
	ErrInvalidInputRows    = errors.New("board: invalid input for rows. Please try again with a number")
	ErrInvalidInputColumns = errors.New("board: invalid input for columns. Please try again with a number")
)
