package robot

import "math/rand"

var warning = []string{
	"Lookout Chief -You are about to fall off the board!",
	"Choose another direction Chief - You are dangerously close to the edge!",
	"Arghh, so close to the edge Chief - Careful!",
	"Another direction perhaps Mr. Robot? - almost at the cliff!",
	"You can do better Mr. Robot, don't fall off now!!",
}

// Warning returns a random warning message
func Warning() string {
	return warning[randInt(0, len(warning)-1)]
}

// randInt returns a random integer between min and max
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
