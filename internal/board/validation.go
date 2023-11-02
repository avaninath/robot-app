package board

import (
	"strconv"
	"strings"
)

func validateUserInput(inputString string) (int, error) {
	inputInt, err := strconv.Atoi(strings.TrimSpace(inputString))
	if err != nil || inputInt <= 0 {
		return 0, ErrInvalidInput
	}
	return inputInt, nil
}
