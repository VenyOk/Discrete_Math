package main

import (
	"fmt"
	"math"
	"sort"
)

func isprime(n int) bool {
	count := 2
	for i := 2; i < int(math.Sqrt(float64(n)))+1; i++ {
		if n%i == 0 {
			count += 1
			if count > 2 {
				break
			}
		}
	}

	return count == 2
}
func solve(n int) {
	arr := make([]int, 0)
	for i := 1; i < int(math.Sqrt(float64(n)))+1; i++ {
		if n%i == 0 {
			if i*i == n {
				arr = append(arr, i)
			} else {
				arr = append(arr, i)
				arr = append(arr, n/i)
			}
		}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	for i := 0; i < len(arr); i++ {
		fmt.Println("\t", arr[i])
	}
	for i := len(arr) - 1; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if arr[j]%arr[i] == 0 && isprime(arr[j]/arr[i]) {
				fmt.Printf("\t%d -- %d\n", arr[j], arr[i])
			}
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println("graph {")
	solve(n)
	fmt.Print("}")
}
