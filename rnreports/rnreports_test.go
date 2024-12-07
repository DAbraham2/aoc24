package rnreports

import (
	"fmt"
	"testing"
)

func TestSafe(t *testing.T) {
	var tests = []struct {
		input []int
		want  bool
	}{
		{[]int{7, 6, 4, 2, 1}, true},
		{[]int{1, 2, 7, 8, 9}, false},
		{[]int{9, 7, 6, 2, 1}, false},
		{[]int{1, 3, 2, 4, 5}, false},
		{[]int{8, 6, 4, 4, 1}, false},
		{[]int{1, 3, 6, 7, 9}, true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.input)
		t.Run(testname, func(t *testing.T) {
			ans := Safe(&tt.input)
			if ans != tt.want {
				t.Errorf("got %t, want %t", ans, tt.want)
			}
		})
	}
}

func TestSafeDampener(t *testing.T) {
	var tests = []struct {
		input []int
		want  bool
	}{
		{[]int{7, 6, 4, 2, 1}, true},
		{[]int{1, 2, 7, 8, 9}, false},
		{[]int{9, 7, 6, 2, 1}, false},
		{[]int{1, 3, 2, 4, 5}, true},
		{[]int{8, 6, 4, 4, 1}, true},
		{[]int{1, 3, 6, 7, 9}, true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.input)
		t.Run(testname, func(t *testing.T) {
			ans := Safe(&tt.input)
			if ans != tt.want {
				t.Errorf("got %t, want %t", ans, tt.want)
			}
		})
	}
}
