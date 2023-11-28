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

func TestIsGreaterThan5(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"ok", args{6}, true},
		{"ok2", args{4}, false},
		{"fail", args{6}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsGreaterThan5(tt.args.x); got != tt.want {
				t.Errorf("IsGreaterThan5() = %v, want %v", got, tt.want)
			}
		})
	}
}
