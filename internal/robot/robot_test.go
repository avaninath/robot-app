package robot

import (
	"testing"

	"github.com/robotAssignment/internal/board"
)

func TestRobot_MoveForward(t *testing.T) {
	tests := []struct {
		skip     bool
		name     string
		newRobot *Robot
		wantErr  bool
	}{
		{
			skip: false,
			name: "MoveForward North - OK",
			newRobot: &Robot{
				Row:       1,
				Column:    1,
				Direction: DirectionNorth,
			},
			wantErr: false,
		},
		{
			skip: false,
			name: "MoveForward South - OK",
			newRobot: &Robot{
				Row:       1,
				Column:    1,
				Direction: DirectionSouth,
			},
			wantErr: false,
		},
		{
			skip: false,
			name: "MoveForward East - OK",
			newRobot: &Robot{
				Row:       1,
				Column:    1,
				Direction: DirectionEast,
			},
			wantErr: false,
		},
		{
			skip: false,
			name: "MoveForward West - OK",
			newRobot: &Robot{
				Row:       1,
				Column:    1,
				Direction: DirectionWest,
			},
			wantErr: false,
		},
		{
			skip: false,
			name: "MoveForward North - Out of bound",
			newRobot: &Robot{
				Row:       0,
				Column:    1,
				Direction: DirectionNorth,
			},
			wantErr: true,
		},
	}
	for _, test := range tests {
		if test.skip {
			continue
		}
		t.Run(test.name, func(t *testing.T) {
			newBoardoard := board.Board{
				MaxRows:    5,
				MaxColumns: 5,
			}
			if err := test.newRobot.MoveForward(&newBoardoard); (err != nil) != test.wantErr {
				t.Errorf("Robot.MoveForward() error = %v, wantErr %v", err, test.wantErr)
			}
		})
	}
}

func Test_validateCommandInput(t *testing.T) {
	tests := []struct {
		skip    bool
		name    string
		command string
		wantErr bool
		err     error
	}{
		{
			skip:    false,
			name:    "ValidateCommandInput - OK",
			command: "LRF",
			wantErr: false,
			err:     nil,
		},
		{
			skip:    false,
			name:    "ValidateCommandInput - Return ErrInvalidCommandInput when command is invalid",
			command: "LRFK",
			wantErr: true,
			err:     ErrInvalidCommandInput,
		},
	}
	for _, test := range tests {
		if test.skip {
			continue
		}
		t.Run(test.name, func(t *testing.T) {
			if err := validateCommandInput(test.command); (err != nil) != test.wantErr {
				t.Errorf("validateCommandInput() error = %v, wantErr %v", err, test.wantErr)
			}
		})
	}
}

func TestRobot_ExecuteCommand(t *testing.T) {
	tests := []struct {
		skip    bool
		name    string
		command string
		wantErr bool
	}{
		{
			skip:    false,
			name:    "ExecuteCommand - OK",
			command: "LRF",
			wantErr: false,
		},
		{
			skip:    false,
			name:    "ExecuteCommand - Return ErrInvalidCommandInput when command is invalid",
			command: "LRFK",
			wantErr: true,
		},
		{
			skip:    false,
			name:    "ExecuteCommand - Return ErrRobotFellOffBoard when robot is out of bound",
			command: "FFFFF",
			wantErr: true,
		},
	}
	for _, test := range tests {
		if test.skip {
			continue
		}
		t.Run(test.name, func(t *testing.T) {
			newBoardoard := board.Board{
				MaxRows:    5,
				MaxColumns: 5,
			}
			newRobot := Robot{
				Row:       1,
				Column:    1,
				Direction: DirectionNorth,
			}

			if err := newRobot.ExecuteCommand(test.command, &newBoardoard); (err != nil) != test.wantErr {
				t.Errorf("Robot.ExecuteCommand() error = %v, wantErr %v", err, test.wantErr)
			}
		})
	}
}

func TestRobot_Report(t *testing.T) {
	type fields struct {
		Row       int
		Column    int
		Direction Direction
	}
	type args struct {
		b *board.Board
	}
	tests := []struct {
		skip   bool
		name   string
		fields fields
		args   args
	}{
		{
			skip: false,
			name: "Report - OK",
			fields: fields{
				Row:       1,
				Column:    1,
				Direction: DirectionNorth,
			},
			args: args{
				b: &board.Board{
					MaxRows:    5,
					MaxColumns: 5,
				},
			},
		},
		{
			skip: false,
			name: "Report when corner - OK",
			fields: fields{
				Row:       0,
				Column:    0,
				Direction: DirectionNorth,
			},
			args: args{
				b: &board.Board{
					MaxRows:    5,
					MaxColumns: 5,
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := &Robot{
				Row:       test.fields.Row,
				Column:    test.fields.Column,
				Direction: test.fields.Direction,
			}
			r.Report(test.args.b)
		})
	}
}
