package main

import (
	"fmt"
	"math"
)

func add(a, b []int32, p int) []int32 {
	var result []int32 = make([]int32, 0)
	j := 0
	mn := math.Min(float64(len(a)), float64(len(b)))
	mx := math.Max(float64(len(a)), float64(len(b)))
	for i := 0; float64(i) < mn; i++ {
		if int(a[i]+b[i]+int32(j)) >= p {
			if j == 1 {
				result = append(result, int32(int(a[i]+b[i]+1)%p))
			} else {
				result = append(result, int32(int(a[i]+b[i])%p))
			}
			j = 1
		} else {
			if j == 1 {
				result = append(result, int32(int(a[i]+b[i]+1)%p))
			} else {
				result = append(result, int32(int(a[i]+b[i])%p))
			}
			j = 0
		}
	}
	if len(b) > len(a) {
		for i := int(mn); float64(i) < mx; i++ {
			if j == 1 {
				result = append(result, int32(int(b[i]+1)%p))
				j = 1
			} else {
				result = append(result, int32(int(b[i])%p))
				j = 0
			}
		}
	} else if len(a) > len(b) {
		for i := int(mn); float64(i) < mx; i++ {
			if j == 1 {
				result = append(result, int32(int(a[i]+1)%p))
				j = 1
			} else {
				result = append(result, int32(int(a[i])%p))
				j = 0
			}
		}
	}
	return result
}
func main() {
	var a = []int32{1, 1, 0, 1}
	var b = []int32{0, 1, 0, 0, 0, 1}
	fmt.Println(add(a, b, 2))
}
