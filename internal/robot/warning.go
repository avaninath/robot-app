package robot

import (
	"fmt"
	"math/rand"

	"github.com/robotAssignment/internal/board"
)

type cornerCoordinates struct {
	row    int
	column int
}

type Corner string

const (
	TopLeft     Corner = "TopLeft"
	TopRight    Corner = "TopRight"
	BottomLeft  Corner = "BottomLeft"
	BottomRight Corner = "BottomRight"
)

var warning = []string{
	"Lookout Chief -You are about to fall off the board!",
	"Choose another direction Chief - You are dangerously close to the edge!",
	"Arghh, so close to the edge Chief - Careful!",
	"Another direction perhaps Mr. Robot? - almost at the cliff!",
	"You can do better Mr. Robot, don't fall off now!!",
}

func warningMessage() string {
	return warning[randInt(0, len(warning)-1)]
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func getCornerDirections(c Corner) string {
	switch c {
	case TopLeft:
		return fmt.Sprintf("%s and %s", DirectionNorth, DirectionWest)
	case TopRight:
		return fmt.Sprintf("%s and %s", DirectionNorth, DirectionEast)
	case BottomLeft:
		return fmt.Sprintf("%s and %s", DirectionSouth, DirectionWest)
	case BottomRight:
		return fmt.Sprintf("%s and %s", DirectionSouth, DirectionEast)
	default:
		return ""
	}
}

func getBoardCornerCoordinates(b *board.Board) map[Corner][]cornerCoordinates {
	return map[Corner][]cornerCoordinates{
		TopLeft: {
			{row: 0, column: 0},
		},
		TopRight: {
			{row: 0, column: b.MaxColumns - 1},
		},
		BottomLeft: {
			{row: b.MaxRows - 1, column: 0},
		},
		BottomRight: {
			{row: b.MaxRows - 1, column: b.MaxColumns - 1},
		},
	}
}
