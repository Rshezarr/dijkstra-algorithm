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
	index, val1 := findMin(table[start], 0)
	fmt.Println(index)
	fmt.Println(val1)
	fmt.Println("=============")
	if index == dest {
		fmt.Println("first check")
		return []int{4, 0, 9, 3}
	}

	new2, val2 := findMin(table[index], val1)
	fmt.Println(new2)
	fmt.Println(val2)
	fmt.Println("=============")

	if new2 == dest {
		fmt.Println("sec check")
		return []int{4, 0, 9, 3}
	}

	new3, val3 := findMin(table[new2], val2)
	fmt.Println(new3)
	fmt.Println(val3)
	fmt.Println("=============")

	if new3 == dest {
		fmt.Println("third check")
		fmt.Println(table[new2])
		return []int{4, 0, 9, 3}
	}

	new4, _ := findMin(table[new3], val3)
	fmt.Println(new4)
	fmt.Println("=============")

	if new4 == dest {
		fmt.Println("last check")
		return []int{4, 0, 9, 3}
	}

	return nil
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
