package board

import "testing"

func TestCreateBoard(t *testing.T) {
	tests := []struct {
		skip    bool
		name    string
		rows    int
		columns int
		want    Board
		wantErr bool
	}{
		{
			skip:    false,
			name:    "CreateBoard - OK",
			rows:    5,
			columns: 5,
			want: Board{
				MaxRows:    5,
				MaxColumns: 5,
			},
			wantErr: false,
		},
		{
			skip:    false,
			name:    "CreateBoard - invalid input 1",
			rows:    0,
			columns: 0,
			want:    Board{},
			wantErr: true,
		},
		{
			skip:    false,
			name:    "CreateBoard - invalid input 2",
			rows:    -1,
			columns: 5,
			want:    Board{},
			wantErr: true,
		},
	}
	for _, test := range tests {
		if test.skip {
			continue
		}
		t.Run(test.name, func(t *testing.T) {
			got, err := createBoard(test.rows, test.columns)
			if (err != nil) != test.wantErr {
				t.Errorf("CreateBoard() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if got.MaxRows != test.want.MaxRows {
				t.Errorf("CreateBoard() got = %v, want %v", got.MaxRows, test.want.MaxRows)
			}
			if got.MaxColumns != test.want.MaxColumns {
				t.Errorf("CreateBoard() got = %v, want %v", got.MaxColumns, test.want.MaxColumns)
			}
		})
	}
}
