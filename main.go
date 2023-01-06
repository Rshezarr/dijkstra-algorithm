package main

import (
	"errors"
)

func FindShortestPath(matrix [][]int, start int, dest int) ([]int, error) {
	var value int
	var path []int
	var err error

	// Start at the starting node
	index := start

	// Loop until we reach the destination node
	for index != dest {
		// Find the next node with the minimum weight
		index, value, err = findMinWeight(matrix[index], value)
		if err != nil {
			return nil, err
		}

		// Add the weight to the path
		path = append(path, value)
	}

	return path, nil
}

func findMinWeight(node []int, prev int) (int, int, error) {
	var index int
	var value int

	// Find the first non-zero, non-prev value in the node
	for i, v := range node {
		if v != prev && v != 0 {
			value = v
			index = i
			break
		}
	}

	// Find the minimum value that is not prev or 0
	for i := index; i < len(node); i++ {
		// Check for negative values
		if node[i] < 0 {
			return 0, 0, errors.New("found negative value")
		}

		// Skip prev and 0 values
		if node[i] == prev || node[i] == 0 {
			continue
		}

		// Update the minimum value and index
		if node[i] < value {
			value = node[i]
			index = i
		}
	}

	return index, value, nil
}
