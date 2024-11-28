package ex03

import "testing"

func Test_sumOfSquare(t *testing.T) {
	tests := []struct {
		name    string
		src     []int
		wantSum int
	}{
		{
			name:    "from Task",
			src:     []int{2, 4, 6, 8, 10},
			wantSum: 220,
		},
		{
			name:    "one value in src",
			src:     []int{2},
			wantSum: 4,
		},
		{
			name:    "empty src",
			src:     []int{},
			wantSum: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := sumOfSquare(tt.src); gotSum != tt.wantSum {
				t.Errorf("sumOfSquare() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}
