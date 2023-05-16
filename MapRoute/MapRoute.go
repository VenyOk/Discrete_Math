package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Node struct {
	x, y int
	dist int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	node := x.(*Node)
	*pq = append(*pq, node)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[:n-1]
	return node
}

func dijkstra(matrix [][]int) int {
	n := len(matrix)

	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dist[i][j] = math.MaxInt32
		}
	}

	dist[0][0] = matrix[0][0]

	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	pq := make(PriorityQueue, 0)
	heap.Push(&pq, &Node{0, 0, dist[0][0]})

	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*Node)
		x, y, d := node.x, node.y, node.dist

		for _, dir := range directions {
			newX, newY := x+dir[0], y+dir[1]

			if newX >= 0 && newX < n && newY >= 0 && newY < n {
				newDist := d +matrix[newX][newY]

				if newDist < dist[newX][newY] {
					dist[newX][newY] = newDist
					heap.Push(&pq, &Node{newX, newY, newDist})
				}
			}
		}
	}

	return dist[n-1][n-1]
}


func main() {
	var n int
	fmt.Scan(&n)

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}
	result := dijkstra(matrix)
	fmt.Println(result)
}
