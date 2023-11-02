package game

import (
	"testing"

	"github.com/robotAssignment/internal/board"
	"github.com/robotAssignment/internal/robot"
)

func Test_checkRobotPosition(t *testing.T) {
	type args struct {
		rbt *robot.Robot
		brd *board.Board
	}
	tests := []struct {
		skip    bool
		name    string
		args    args
		wantErr bool
	}{
		{
			skip: false,
			name: "checkRobotPosition - OK",
			args: args{
				rbt: &robot.Robot{
					Row:       1,
					Column:    1,
					Direction: robot.DirectionNorth,
				},
				brd: &board.Board{
					MaxRows:    5,
					MaxColumns: 5,
				},
			},
			wantErr: false,
		},
		{
			skip: false,
			name: "checkRobotPosition - Error",
			args: args{
				rbt: &robot.Robot{
					Row:       1,
					Column:    6,
					Direction: robot.DirectionNorth,
				},
				brd: &board.Board{
					MaxRows:    5,
					MaxColumns: 5,
				},
			},
			wantErr: true,
		},
	}
	for _, test := range tests {
		if test.skip {
			continue
		}
		t.Run(test.name, func(t *testing.T) {
			if err := checkRobotPosition(test.args.rbt, test.args.brd); (err != nil) != test.wantErr {
				t.Errorf("checkRobotPosition() error = %v, wantErr %v", err, test.wantErr)
			}
		})
	}
}
