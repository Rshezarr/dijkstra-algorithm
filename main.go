package main

import "fmt"

func main() {
	table := [][]int{
		{0, 7, 9, 0, 0, 14},
		{7, 0, 10, 15, 0, 0},
		{9, 10, 0, 11, 0, 2},
		{0, 15, 11, 0, 6, 0},
		{0, 0, 0, 6, 0, 9},
		{14, 0, 2, 0, 9, 0},
	}
	fmt.Println(djk(table, 0, 5))
}

func djk(table [][]int, start int, dest int) []int {
	var index int = start
	var value int
	var sum int
	var path []int

	for index != dest {
		index, value = findMin(table[index], value)
		sum += value
		path = append(path, sum)
	}

	return path
}

func findMin(node []int, prev int) (int, int) {
	index := 0
	val := 0
	for i, v := range node {
		if v != prev && v != 0 {
			val = v
			index = i
			break
		}
	}

	for i := index; i < len(node); i++ {
		if node[i] == prev || node[i] == 0 {
			continue
		}
		if node[i] < val {
			val = node[i]
			index = i
		}
	}
	return index, val
}
