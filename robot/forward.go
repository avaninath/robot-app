package robot

func MoveForward(r, c int, direction Direction) (int, int) {
	switch direction {
	case DirectionNorth:
		return r - 1, c
	case DirectionSouth:
		return r + 1, c
	case DirectionEast:
		return r, c + 1
	case DirectionWest:
		return r, c - 1
	default:
		return r, c
	}
}
