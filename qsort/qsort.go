package main

import "fmt"

var arr = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

// Код полностью взят из презентации 2 модуля Алгоритмов и структур данных

func swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func less(i, j int) (result bool) {
	result = arr[i] < arr[j]
	return
}
func partition(l, r int) (i int) {
	i = l
	j := l
	for j < r {
		if less(j, r) {
			swap(i, j)
			i++
		}
		j++
	}
	swap(i, r)
	return
}
func qsortrec(l, r int) {
	if l < r {
		q := partition(l, r)
		qsortrec(l, q-1)
		qsortrec(q+1, r)

	}
}
func qsort(n int, less func(i, j int) bool, swap func(i, j int)) {
	qsortrec(0, n-1)
}
func main() {
	qsort(10, less, swap)
	fmt.Println(arr)
}
