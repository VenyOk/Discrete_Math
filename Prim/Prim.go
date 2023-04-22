package main

import (
	"container/heap"
	"fmt"
)

type edge struct {
	to   int
	cost int
}

type node struct {
	edges   []edge
	visited bool
}

type priorityQueue []edge

func (pq priorityQueue) Len() int           { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool { return pq[i].cost < pq[j].cost }
func (pq priorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *priorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(edge))
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func prim(graph []node) int {
	q := priorityQueue{}
	heap.Init(&q)

	for _, e := range graph[0].edges {
		heap.Push(&q, e)
	}

	graph[0].visited = true
	cost := 0

	for len(q) > 0 {
		e := heap.Pop(&q).(edge)
		if graph[e.to].visited {
			continue
		}

		graph[e.to].visited = true
		cost += e.cost

		for _, next := range graph[e.to].edges {
			if !graph[next.to].visited {
				heap.Push(&q, next)
			}
		}
	}

	return cost
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	graph := make([]node, n)
	for i := 0; i < m; i++ {
		var a, b, c int
		fmt.Scan(&a, &b, &c)
		graph[a].edges = append(graph[a].edges, edge{to: b, cost: c})
		graph[b].edges = append(graph[b].edges, edge{to: a, cost: c})
	}

	fmt.Println(prim(graph))
}
