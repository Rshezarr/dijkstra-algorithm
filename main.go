package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	table := [][]int{
		{0, 7, 9, 0, 0, 14},
		{7, 0, 10, 15, 0, 0},
		{9, 10, 0, 11, 0, 2},
		{0, 15, 11, 0, 6, 0},
		{0, 0, 0, 6, 0, 9},
		{14, 0, 2, 0, 9, 0},
	}
	res, err := djk(table, 0, 5)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(res)
}

func djk(table [][]int, start int, dest int) ([]int, error) {
	var index int = start
	var value int
	// var sum int
	var path []int
	var err error

	for index != dest {
		index, value, err = findMinWeight(table[index], value)
		if err != nil {
			return nil, err
		}
		// sum += value
		path = append(path, value)
	}

	return path, nil
}

func findMinWeight(node []int, prev int) (int, int, error) {
	var index int
	var value int
	for i, v := range node {
		if v != prev && v != 0 {
			value = v
			index = i
			break
		}
	}

	for i := index; i < len(node); i++ {
		if node[i] < 0 {
			return 0, 0, errors.New("found negative value")
		}

		if node[i] == prev || node[i] == 0 {
			continue
		}
		if node[i] < value {
			value = node[i]
			index = i
		}
	}
	return index, value, nil
}
