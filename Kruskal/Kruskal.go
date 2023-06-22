package main

import (
	"fmt"
	"math"
	"sort"
)

const N = 10000000

type vertex struct {
	coord1   int
	coord2   int
	distance float64
}

func solve() (result float64) {
	result = 0.0
	var n int
	fmt.Scan(&n)
	arr := [N][2]int{}
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i][0], &arr[i][1])
	}
	graph := make([]vertex, 0)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			graph = append(graph, vertex{coord1: i, coord2: j, distance: math.Sqrt(math.Pow((float64(arr[i][0]-arr[j][0])), 2) + math.Pow((float64(arr[i][1]-arr[j][1])), 2))})
		}
	}
	sort.Slice(graph, func(i, j int) bool {
		return graph[i].distance < graph[j].distance
	})
	tree_id := make([]int, n)
	for i := 0; i < n; i++ {
		tree_id[i] = i
	}
	for i := 0; i < len(graph); i++ {
		if tree_id[graph[i].coord1] != tree_id[graph[i].coord2] {
			result += graph[i].distance
			old_id := tree_id[graph[i].coord2]
			new_id := tree_id[graph[i].coord1]
			for j := 0; j < n; j++ {
				if tree_id[j] == old_id {
					tree_id[j] = new_id
				}
			}
		}
	}
	return
}
func main() {
	fmt.Printf("%.2f", solve())
}
