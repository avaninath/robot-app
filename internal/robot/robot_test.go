package robot

import (
	"reflect"
	"testing"

	"github.com/robotAssignment/internal/board"
)

func TestInitiate(t *testing.T) {
	type args struct {
		board *board.Board
	}
	tests := []struct {
		name string
		args args
		want Robot
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Initiate(tt.args.board); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Initiate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRobot_MoveForward(t *testing.T) {
	type fields struct {
		Row       int
		Column    int
		Direction Direction
	}
	type args struct {
		board *board.Board
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Robot{
				Row:       tt.fields.Row,
				Column:    tt.fields.Column,
				Direction: tt.fields.Direction,
			}
			if err := r.MoveForward(tt.args.board); (err != nil) != tt.wantErr {
				t.Errorf("Robot.MoveForward() error = %v, wantErr %v", err, tt.wantErr)
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
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Robot{
				Row:       tt.fields.Row,
				Column:    tt.fields.Column,
				Direction: tt.fields.Direction,
			}
			r.Report(tt.args.b)
		})
	}
}

func TestRobot_ExecuteCommand(t *testing.T) {
	type fields struct {
		Row       int
		Column    int
		Direction Direction
	}
	type args struct {
		command string
		b       *board.Board
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Robot{
				Row:       tt.fields.Row,
				Column:    tt.fields.Column,
				Direction: tt.fields.Direction,
			}
			if err := r.ExecuteCommand(tt.args.command, tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("Robot.ExecuteCommand() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_isCloseToEdge(t *testing.T) {
	type args struct {
		rbt Robot
		b   *board.Board
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCloseToEdge(tt.args.rbt, tt.args.b); got != tt.want {
				t.Errorf("isCloseToEdge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRobot_checkCornerPosition(t *testing.T) {
	type fields struct {
		Row       int
		Column    int
		Direction Direction
	}
	type args struct {
		b *board.Board
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Corner
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Robot{
				Row:       tt.fields.Row,
				Column:    tt.fields.Column,
				Direction: tt.fields.Direction,
			}
			if got := r.checkCornerPosition(tt.args.b); got != tt.want {
				t.Errorf("Robot.checkCornerPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateCommandInput(t *testing.T) {
	type args struct {
		command string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateCommandInput(tt.args.command); (err != nil) != tt.wantErr {
				t.Errorf("validateCommandInput() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
