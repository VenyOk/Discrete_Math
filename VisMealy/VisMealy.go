package main

import "fmt"

const N = 1000
const M = 1000

func main() {
	var n, m, q0 int
	fmt.Scan(&n)
	fmt.Scan(&m)
	fmt.Scan(&q0)
	arr := [N][M]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&arr[i][j])
		}
	}
	arr2 := [N][M]string{}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&arr2[i][j])
		}
	}
	fmt.Println("digraph {\n \trankdir = LR")
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("\t%d -> %d [label = \"%c(%s)\"]\n", i, arr[i][j], 97+j, arr2[i][j])
		}
	}
	fmt.Println("}")
}

