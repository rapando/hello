package main

import "testing"

func TestAddNumbers(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"ok:1", args{2, 3}, 5},
		{"ok:2", args{3, 3}, 6},
		{"ok:3", args{30, 3}, 33},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddNumbers(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("AddNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
