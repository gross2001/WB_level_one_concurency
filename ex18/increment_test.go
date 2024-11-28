package ex18

import "testing"

func Test_run(t *testing.T) {
	tests := []struct {
		name             string
		howManyIncrement int
		want             int
	}{
		{
			name:             "1 times",
			howManyIncrement: 1,
			want:             1,
		},
		{
			name:             "1000 times",
			howManyIncrement: 1000,
			want:             1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.howManyIncrement); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
