package main

import (
	"reflect"
	"testing"
)

func TestFindShortestPath(t *testing.T) {
	matrix := [][]int{
		{0, 7, 9, 0, 0, 14},
		{7, 0, 10, 15, 0, 0},
		{9, 10, 0, 11, 0, 2},
		{0, 15, 11, 0, 6, 0},
		{0, 0, 0, 6, 0, 9},
		{14, 0, 2, 0, 9, 0},
	}

	// Test a simple path
	path, err := FindShortestPath(matrix, 0, 3)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(path, []int{7, 10, 2, 9, 6}) {
		t.Errorf("unexpected path: %v", path)
	}

	// Test a path that has a negative value
	matrix[1][3] = -1
	_, err = FindShortestPath(matrix, 0, 3)
	if err == nil {
		t.Errorf("expected error but got none")
	}
}
