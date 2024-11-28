package main

import "testing"

func Test_defineType(t *testing.T) {
	tests := []struct {
		name string
		i    any
		want string
	}{
		{
			name: "int",
			i:    -100,
			want: "int",
		},
		{
			name: "String",
			i:    "HELLO WORLD",
			want: "string",
		},
		{
			name: "func",
			i:    func(int) {},
			want: "func(int)",
		},
		{
			name: "chan int",
			i:    make(chan int, 1),
			want: "chanel chan int",
		},
		{
			name: "slice int",
			i:    make([]int, 1),
			want: "[]int",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := defineType(tt.i); got != tt.want {
				t.Errorf("defineType() = %v, want %v", got, tt.want)
			}
		})
	}
}
