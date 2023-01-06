# Dijkstra Algorithm

This package contains an implementation of the Dijkstra algorithm for finding the shortest path between two nodes in a graph.

## Usage

```go
import "github.com/user/dijkstra"

// Create a 2D slice representing the graph, with the weights of the edges between the nodes.
// A value of 0 indicates that there is no edge between the nodes.
matrix := [][]int{
	{0, 7, 9, 0, 0, 14},
	{7, 0, 10, 15, 0, 0},
	{9, 10, 0, 11, 0, 2},
	{0, 15, 11, 0, 6, 0},
	{0, 0, 0, 6, 0, 9},
	{14, 0, 2, 0, 9, 0},
}

// Find the shortest path between the start node (0) and the end node (5).
path, err := dijkstra.FindShortestPath(matrix, 0, 5)
if err != nil {
	log.Fatal(err)
}

fmt.Println(path)  // [7, 10, 2]
```

## Complexity

The complexity of the Dijkstra algorithm is `O(n^2)`, where `n` is the number of nodes in the graph. This is because the algorithm processes each node and its edges once, and the number of edges is proportional to the number of nodes.

## Notes

-   The graph is represented as a 2D slice, where each element `table[i][j]` represents the weight of the edge between node `i` and node `j`.
-   The `FindShortestPath` function returns an error if it encounters a negative value in the graph, as the Dijkstra algorithm does not support negative edge weights.
