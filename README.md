# Robot App
​
This project is a simple implementation of a robot game. The game consists of a board, a robot, and a set of instructions that the robot can follow to move around the board. The Robot gives us back its final position at the end of the game.
​
The input for the game consists of a set of instructions that the robot can follow. The instructions are provided as a string, and can include the following commands:
​
- `PLACE R,C,F`: Places the robot on the board at position `(Row, Column)` facing direction `(N | W | S | E)`.
- `FORWARD`: Moves the robot one unit in the direction it is facing.
- `LEFT`: Rotates the robot 90 degrees to the left.
- `RIGHT`: Rotates the robot 90 degrees to the right.
​
## Board Creation
​
The board is created as a `nxn` grid of squares. The input is validated to ensure that it is in the correct format.
- Integer row size, e.g - 10
- Integer column size, e.g - 10
​
## Robot Initial Position
​
The input is validated to ensure that it is in the correct format. The `PLACE` command works in the following steps:
- Integer row input , e.g - 5
- Integer column input, e.g - 5
- Initial direction the robot is facing, e.g - `N`
​
## Robot Movement
​
The robot can move around the board using the `FORWARD`, `LEFT`, and `RIGHT` commands. The `FORWARD` command moves the robot one unit in the direction it is facing. The `LEFT` and `RIGHT` commands rotate the robot 90 degrees to the left or right, respectively.
## Tech Stack
​
**App:** Written in Go
​
**Testing:** [Testing in-built package in Go](https://pkg.go.dev/testing)
​
## Run App Locally
**Option 1** 
​
*Note:* *This requires Docker to be installed on the machine*
​
- Pull docker image from DockerHub:
```docker pull avaninath/robot-app:manifest-multi-arch```
(*Architecture supported are linux/amd64 and linux/arm64*)
​
- Verify docker image is downloaded:
```docker image ls | grep avaninath/robot-app```
​
- Start docker container and play the game:
```docker run -it avaninath/robot-app:manifest-multi-arch```
​
**Option 2**
*Note:* *This requires Go & dependencies to be installed on the machine*
​
- Clone the project from GitHub
​
```git clone https://link-to-project```
​
Go to the project directory
​
```cd robotAssignmentDevoTeam```
​
Start the app locally and play the game:
​
```go run cmd/main.go```
## Running Tests
​
To run tests, run the following command
​
```bash
  cd my-project && make test
```
## Screenshots
