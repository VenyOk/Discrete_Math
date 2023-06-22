package main

import (
	"fmt"
	"sort"
)

type pair struct {
	first, second []int
}

func dfs(s int, state int, incomp [][]int, color []int, part *pair) bool {
	color[s] = state

	if state == 1 {
		part.first = append(part.first, s)
	} else {
		part.second = append(part.second, s)
	}

	for _, val := range incomp[s] {
		if color[val] == 0 {
			if state == 1 {
				if !dfs(val, 2, incomp, color, part) {
					return false
				}
			} else {
				if !dfs(val, 1, incomp, color, part) {
					return false
				}
			}
		} else if color[val] == state {
			return false
		}
	}

	return true
}

func concatVectors(src1, src2 []int) []int {
	result := make([]int, len(src1)+len(src2))
	copy(result, src1)
	copy(result[len(src1):], src2)
	return result
}

func solution(c int, parts []pair, result pair) pair {
	if c == len(parts) {
		return result
	}

	firstVariant := solution(c+1, parts, pair{concatVectors(result.first, parts[c].first), concatVectors(result.second, parts[c].second)})
	secondVariant := solution(c+1, parts, pair{concatVectors(result.first, parts[c].second), concatVectors(result.second, parts[c].first)})

	delta := abs(len(firstVariant.first)-len(firstVariant.second)) - abs(len(secondVariant.first)-len(secondVariant.second))

	if delta < 0 {
		return firstVariant
	} else if delta == 0 {
		if len(firstVariant.first) < len(firstVariant.second) {
			if ((len(secondVariant.first) < len(secondVariant.second) && less(firstVariant.first, secondVariant.first)) ||
				(len(secondVariant.first) >= len(secondVariant.second) && less(firstVariant.first, secondVariant.second))) {
				return firstVariant
			} else {
				return secondVariant
			}
		} else {
			if ((len(secondVariant.first) < len(secondVariant.second) && less(firstVariant.second, secondVariant.first)) ||
				(len(secondVariant.first) >= len(secondVariant.second) && less(firstVariant.second, secondVariant.second))) {
				return firstVariant
			} else {
				return secondVariant
			}
		}
	} else {
		return secondVariant
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func less(a, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return false
}

func main() {
	var n int
	fmt.Scan(&n)

	incomp := make([][]int, n)
	var ch string
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&ch)
			if ch == "+" {
				incomp[i] = append(incomp[i], j)
			}
		}
	}

	color := make([]int, n)
	var parts []pair

	for i := 0; i < n; i++ {
		if color[i] == 0 {
			part := &pair{nil, nil}
			if !dfs(i, 1, incomp, color, part) {
				fmt.Println("No solution")
				return
			}
			parts = append(parts, *part)
		}
	}

	optimal := solution(0, parts, pair{nil, nil})

	sort.Ints(optimal.first)
	sort.Ints(optimal.second)

	if len(optimal.first) < len(optimal.second) || (len(optimal.first) == len(optimal.second) && less(optimal.first, optimal.second)) {
		for _, val := range optimal.first {
			fmt.Print(val+1, " ")
		}
		fmt.Println()
	} else {
		for _, val := range optimal.second {
			fmt.Print(val+1, " ")
		}
		fmt.Println()
	}
}
