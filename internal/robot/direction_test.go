package robot

import "testing"

func TestDirection_IsValid(t *testing.T) {
	tests := []struct {
		name string
		d    Direction
		want bool
	}{
		{
			name: "valid - North",
			d:    DirectionNorth,
			want: true,
		},
		{
			name: "valid - South",
			d:    DirectionSouth,
			want: true,
		},
		{
			name: "valid - East",
			d:    DirectionEast,
			want: true,
		},
		{
			name: "valid - West",
			d:    DirectionWest,
			want: true,
		},
		{
			name: "invalid direction 1",
			d:    Direction("invalid"),
			want: false,
		},
		{
			name: "invalid direction 2",
			d:    Direction(""),
			want: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := test.d.IsValid(); got != test.want {
				t.Errorf("Direction.IsValid() = %v, want %v", got, test.want)
			}
		})
	}
}
