package main

func partition(l int, r int, less func(i, j int) bool, swap func(i, j int)) (i int) {
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

func qsortrec(l int, r int, less func(i, j int) bool, swap func(i, j int)) {
	if l < r {
		q := partition(l, r, less, swap)
		qsortrec(l, q-1, less, swap)
		qsortrec(q+1, r, less, swap)

	}
}

func qsort(n int, less func(i, j int) bool, swap func(i, j int)) {
	qsortrec(0, n-1, less, swap)
}

func main() {
}
