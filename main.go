package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/robotAssignment/robot"
)

var in = bufio.NewReader(os.Stdin)

func main() {
	fmt.Println("Welcome to the Robot Game")
	fmt.Println(strings.Repeat("-", 25))
	fmt.Println("\nPlease enter the size of the board")
	fmt.Println("Enter the number of rows of the room:")
	r, _ := in.ReadString('\n')
	fmt.Println("Enter the number of columns of the room:")
	c, _ := in.ReadString('\n')

	fmt.Printf("The size of the board is %v x %v \n", strings.TrimSpace(r), strings.TrimSpace(c))

	newRbt := robot.Initiate()

	fmt.Println("Please enter the commands for the robot")
	fmt.Println("Enter the commands:")
	commands, _ := in.ReadString('\n')

	fmt.Println("The commands for the robot are " + strings.TrimSpace(commands))

	finalDirection := robot.GetNextDirection(newRbt.Direction, strings.TrimSpace(commands))
	fmt.Println("The final direction of the robot is " + finalDirection)
}
