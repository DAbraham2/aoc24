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
		{[]int{8, 1, 3, 6, 7, 9}, true},
		{[]int{1, 5, 6, 7}, true},
		{[]int{2, 1, 2, 3, 4}, true},
		{[]int{48, 46, 47, 49, 51, 54, 56}, true},
		{[]int{1, 1, 2, 3, 4, 5}, true},
		{[]int{1, 2, 3, 4, 5, 5}, true},
		{[]int{5, 1, 2, 3, 4, 5}, true},
		{[]int{1, 4, 3, 2, 1}, true},
		{[]int{1, 6, 7, 8, 9}, true},
		{[]int{1, 2, 3, 4, 3}, true},
		{[]int{9, 8, 7, 6, 7}, true},
		{[]int{7, 10, 8, 10, 11}, true},
		{[]int{29, 28, 27, 25, 26, 25, 22, 20}, true},
		{[]int{7, 10, 8, 10, 11}, true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.input)
		t.Run(testname, func(t *testing.T) {
			ans := SafeDampener(&tt.input)
			if ans != tt.want {
				t.Errorf("got %t, want %t", ans, tt.want)
			}
		})
	}
}

func TestDampener(t *testing.T) {
	in := []int{1, 3, 2, 4, 5}
	ans := SafeDampener(&in)

	if ans != true {
		t.Errorf("got %t, want %t", ans, true)
	}
}
