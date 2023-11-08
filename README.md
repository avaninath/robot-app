# Robot Game

This project is a simple yet efficient implementation of a robot game. The game consists of a board, a robot, and a set of instructions that the robot can follow to move around the board. The Robot gives us back its final position at the end of the game.

The input for the game consists of a set of instructions that the robot can follow. The prompts for the game allow user to enter the inputs for the game to initiate.

## Game Features
  - Board creation
    - number of rows
    - number of columns

  - Robot creation
    - starting row position (int)
    - starting column position (int)
    - initial direction (`L R F`)
  
  - Reporting initial position of the robot on the board

  - Allows user to input commands in a single string format to move the robot around the board
    - `L` - rotate the robot 90 degrees to the left
    - `R` - rotate the robot 90 degrees to the right
    - `F` - move the robot one unit forward in the direction it is facing

  - Reporting final position of the robot on the board
    - The final position of the robot is reported as coordinates on the board and the direction it is facing.
    - If the robot is close to any edge or at a corner, it will issue an warning to the user to be careful with the next move, alongwith the direction which will cause the robot to fall off.
    - If the robot falls off the board, the game is over and the final position is reported as coordinates that were out of bounds of the board.

  - Allows user to choose to continue playing by entering more commands for the robot or to quit the game gracefully as prompted.

  - Allows user to quit the game or start a new game once the robot has fallen off the board.


## Board Creation
The board is created as a `n x n` grid of squares. The input is validated to ensure that it is in the correct format.
- Integer row size, e.g - 10
- Integer column size, e.g - 10

#### User Input and Validation
The user input for the board is taken as number of rows and columns, prompted by the game. 

Input is validated and the following are not allowed:
  - Non-integer input
  - Negative integers
  - `0` as the `row` or `column` input 


## Robot Creation and Positioning
Once the board is created, the user is prompted to create a robot and position it on the board. The robot is positioned on the board using the following inputs:
- Integer row position, e.g - 1
- Integer column position, e.g - 1
- Direction, e.g - `NORTH`

#### User Input and Validation
The input is validated to ensure that it is in the correct format.
- Row/Column - not allowed:
  - Non-integer input
  - Negative integers
  - Position greater than the board size

- Direction
  - Any input other than `N`, `S`, `E`, `W` is not allowed
  - Input is not case sensitive


## Robot Movement
The robot can be moved on the board using a single string command input, for e.g `LRFLR`.

#### User Input and Validation
- Any input other than `L`, `R`, `F` is not allowed
- Input is not case sensitive


## Warning
The robot is programmed to issue a warning to the user if it is close to any edge or at a corner, and the next move will cause it to fall off the board. The warning is issued alongwith the direction which will cause the robot to fall off.

#### Cases where warning is issued
- Robot is at any edge of the board and the next move in the same direction will cause it to fall off
  - A random warning message will be printed to the user and the direction which will cause the robot to fall off

- Robot is at one of the four corners of the board
  - A warning message will be displayed with both the directions that will cause the robot to fall off


## Tech Stack
**App:** Written in Go

**Testing:** [Testing in-built package in Go](https://pkg.go.dev/testing)

## Run App Locally
**Option 1** 

*Note:* *This requires Docker to be installed on the machine*

- Pull docker image from DockerHub:
```docker pull avaninath/robot-app:manifest-multi-arch```
(*Architecture supported are linux/amd64 and linux/arm64*)

- Verify docker image is downloaded:
```docker image ls | grep avaninath/robot-app```

- Start docker container and play the game:
```docker run -it avaninath/robot-app:manifest-multi-arch```


**Option 2**

*Note:* *This requires Go & dependencies to be installed on the machine*

- Clone the project from GitHub

```git clone https://github.com/avaninath/robot-app.git```

Go to the project directory

```cd robot-app```

Start the app locally and play the game:

```go run cmd/main.go```


## Running Tests​
To run tests, run the following command

```cd robot-app && make test```


## Screenshots

![robot](https://github.com/avaninath/robotAssignmentDevoTeam/assets/55435998/b1e0916d-0e0d-401b-b6ad-8f5df7366ecf)


## DemoPlayList


https://github.com/avaninath/robotAssignmentDevoTeam/assets/55435998/a11f3ad0-186d-46ea-abf0-9f9c17f3cb03



https://github.com/avaninath/robotAssignmentDevoTeam/assets/55435998/4d9bd241-c29d-4165-8688-5ef817b7e3d6




https://github.com/avaninath/robotAssignmentDevoTeam/assets/55435998/3c49be26-bc73-458d-8586-48c8dbad5016


