package main

import (
	"fmt"
)

var count int

type TYPE interface {
	int | int8 | int16 | int32 | int64
}

func min[T TYPE](a T, b T) T {
	if a > b {
		return b
	} else {
		return a
	}
}
func solve() {
	count = 0
	var n, m int
	fmt.Scan(&n)
	fmt.Scan(&m)
	graph := make([][]int, n, n)
	used := make([]bool, n)
	tin := make([]int, n)
	fup := make([]int, n)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Scan(&u, &v)
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}
	for i := 0; i < n; i++ {
		used[i] = false
	}
	for i := 0; i < n; i++ {
		if !used[i] {
			dfs(i, -1, used, tin, fup, 0, graph)
		}
	}
	fmt.Println(count)
}

func dfs(v int, p int, used []bool, tin []int, fup []int, timer int, graph [][]int) {
	used[v] = true
	timer++
	tin[v] = timer
	fup[v] = timer
	for i := 0; i < len(graph[v]); i++ {
		if graph[v][i] == p {
			continue
		}
		if used[graph[v][i]] {
			fup[v] = min(fup[v], fup[graph[v][i]])
		} else {
			dfs(graph[v][i], v, used, tin, fup, timer, graph)
			fup[v] = min(fup[v], fup[graph[v][i]])
			if fup[graph[v][i]] > tin[v] {
				count++
			}
		}
	}
}
func main() {
	solve()
}
