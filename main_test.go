package main

import (
	"testing"

	nydigits "github.com/timgluz/nydigits/pkg"
)

// TestRun does sanity tests on the run function to ensure basic errors are handled
// and after that it would check solutions for some real examples from the game.
func TestRun(t *testing.T) {
	tests := []struct {
		name     string
		target   int
		digits   []int
		wantErr  bool
		solution nydigits.Solution
	}{
		{name: "null target", target: 0, digits: []int{1}, wantErr: true, solution: nydigits.Solution{}},
		{name: "null digits", target: 1, digits: []int{}, wantErr: true, solution: nydigits.Solution{}},
		{
			name:    "April14, 1st",
			target:  59,
			digits:  []int{2, 3, 5, 11, 15, 25},
			wantErr: false,
			solution: nydigits.Solution{
				Value: 59,
				Operations: []nydigits.OperationStep{
					{Op: nydigits.Plus, Digit: 3, PrevValue: 0, Value: 3},
					{Op: nydigits.Plus, Digit: 15, PrevValue: 3, Value: 45},
					{Op: nydigits.Plus, Digit: 25, PrevValue: 45, Value: 70},
					{Op: nydigits.Minus, Digit: 11, PrevValue: 70, Value: 59},
				},
			},
		},
		{
			name:    "April14, 2nd",
			target:  133,
			digits:  []int{4, 5, 8, 11, 15, 20},
			wantErr: false,
			solution: nydigits.Solution{
				Value: 133,
				Operations: []nydigits.OperationStep{
					{Op: nydigits.Plus, Digit: 20, PrevValue: 0, Value: 20},
					{Op: nydigits.Minus, Digit: 4, PrevValue: 20, Value: 16},
					{Op: nydigits.Times, Digit: 8, PrevValue: 16, Value: 128},
					{Op: nydigits.Plus, Digit: 5, PrevValue: 128, Value: 133},
				},
			},
		},
		{
			name:    "April14, 3rd",
			target:  218,
			digits:  []int{4, 5, 7, 9, 11, 20},
			wantErr: false,
			solution: nydigits.Solution{
				Value: 218,
				Operations: []nydigits.OperationStep{
					{Op: nydigits.Plus, Digit: 5, PrevValue: 0, Value: 5},
					{Op: nydigits.Plus, Digit: 20, PrevValue: 5, Value: 25},
					{Op: nydigits.Times, Digit: 9, PrevValue: 25, Value: 225},
					{Op: nydigits.Minus, Digit: 7, PrevValue: 225, Value: 218},
				},
			},
		},
		{
			name:    "April14, 4th",
			target:  388,
			digits:  []int{3, 5, 9, 20, 23, 25},
			wantErr: false,
			solution: nydigits.Solution{
				Value: 388,
				Operations: []nydigits.OperationStep{
					{Op: nydigits.Plus, Digit: 25, PrevValue: 0, Value: 25},
					{Op: nydigits.Minus, Digit: 9, PrevValue: 25, Value: 16},
					{Op: nydigits.Times, Digit: 23, PrevValue: 16, Value: 368},
					{Op: nydigits.Plus, Digit: 20, PrevValue: 368, Value: 388},
				},
			},
		},
		{
			name:    "April14, 5th",
			target:  462,
			digits:  []int{3, 5, 9, 19, 20, 25},
			wantErr: false,
			solution: nydigits.Solution{
				Value: 462,
				Operations: []nydigits.OperationStep{
					{Op: nydigits.Plus, Digit: 3, PrevValue: 0, Value: 3},
					{Op: nydigits.Plus, Digit: 20, PrevValue: 3, Value: 23},
					{Op: nydigits.Times, Digit: 19, PrevValue: 23, Value: 437},
					{Op: nydigits.Plus, Digit: 25, PrevValue: 437, Value: 462},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := run(tt.target, tt.digits)
			if (err != nil) != tt.wantErr {
				t.Errorf("run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got.Value != tt.solution.Value {
				t.Errorf("run() got = %v, want %v", got, tt.solution)
			}
		})
	}
}
