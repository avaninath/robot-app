package board

import (
	"reflect"
	"testing"
)

func TestInitiate(t *testing.T) {
	tests := []struct {
		name string
		want Board
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Initiate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Initiate() = %v, want %v", got, tt.want)
			}
		})
	}
}
