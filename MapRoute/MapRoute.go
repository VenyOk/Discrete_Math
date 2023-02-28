package main

import (
	"fmt"
	"math"
)

func solve() {
	var n, k int
	fmt.Scan(&n)
	arr := make([][]int, n, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&k)
			arr[i] = append(arr[i], k)
		}
	}
	arr2 := make([][]int, n, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			arr2[i] = append(arr2[i], 0)
		}
	}
	arr2[0][0] = arr[0][0]
	for i := 1; i < n; i++ {
		arr2[0][i] += arr[0][i] + arr2[0][i-1]
	}
	for i := 1; i < n; i++ {
		arr2[i][0] += arr[i][0] + arr2[i-1][0]
	}
	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			arr2[i][j] = int(math.Min(float64(arr2[i-1][j]), float64(arr2[i][j-1]))) + arr[i][j]
		}
		for j := 1; j < n; j++ {
			arr2[j][i] = int(math.Min(float64(arr2[j-1][i]), float64(arr2[j][i-1]))) + arr[j][i]
		}
	}
	fmt.Println(arr2[n-1][n-1])
}

func main() {
	solve()
}
