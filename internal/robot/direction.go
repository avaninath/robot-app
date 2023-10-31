package robot

// Direction is the type for the direction of the robot
type Direction string

// The possible directions of the robot
const (
	DirectionNorth Direction = "N"
	DirectionSouth Direction = "S"
	DirectionEast  Direction = "E"
	DirectionWest  Direction = "W"
)

// ToString converts a Direction to a string
func (d Direction) ToString() string {
	return string(d)
}

// IsValid checks if the input for direction is valid
func (d Direction) IsValid() bool {
	switch d {
	case DirectionNorth:
		return true
	case DirectionSouth:
		return true
	case DirectionEast:
		return true
	case DirectionWest:
		return true
	default:
		return false
	}
}

// ToLeft turns the robot once to the left
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

// ToRight turns the robot once to the right
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

// ToDirection converts a string to a Direction
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
		return "Invalid input"
	}
}
