package robot

type Direction string

const (
	DirectionNorth Direction = "N"
	DirectionSouth Direction = "S"
	DirectionEast  Direction = "E"
	DirectionWest  Direction = "W"
)

func (d Direction) ToString() string {
	return string(d)
}

func ToLeft(direction Direction) Direction {
	switch direction {
	case DirectionNorth:
		return DirectionWest
	case DirectionSouth:
		return DirectionEast
	case DirectionEast:
		return DirectionNorth
	case DirectionWest:
		return DirectionSouth
	default:
		return direction
	}
}

func ToRight(direction Direction) Direction {
	switch direction {
	case DirectionNorth:
		return DirectionEast
	case DirectionSouth:
		return DirectionWest
	case DirectionEast:
		return DirectionSouth
	case DirectionWest:
		return DirectionNorth
	default:
		return direction
	}
}

func ToDirection(direction string) Direction {
	switch direction {
	case "N":
		return DirectionNorth
	case "S":
		return DirectionSouth
	case "E":
		return DirectionEast
	case "W":
		return DirectionWest
	default:
		return ""
	}
}

func GetNextDirection(direction Direction, commands string) Direction {
	for _, command := range commands {
		switch command {
		case 'L':
			direction = ToLeft(direction)
		case 'R':
			direction = ToRight(direction)
		}
	}
	return direction
}
