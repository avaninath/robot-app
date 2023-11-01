package board

import "testing"

func Test_validateUserInput(t *testing.T) {
	tests := []struct {
		skip    bool
		name    string
		input   string
		want    int
		wantErr bool
	}{
		{
			skip:    false,
			name:    "validateUserInput - OK",
			input:   "5",
			want:    5,
			wantErr: false,
		},
		{
			skip:    false,
			name:    "validateUserInput - invalid input",
			input:   "a",
			want:    0,
			wantErr: true,
		},
	}
	for _, test := range tests {
		if test.skip {
			continue
		}
		t.Run(test.name, func(t *testing.T) {
			got, err := validateUserInput(test.input)
			if (err != nil) != test.wantErr {
				t.Errorf("validateUserInput() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if got != test.want {
				t.Errorf("validateUserInput() got = %v, want %v", got, test.want)
			}
		})
	}
}
