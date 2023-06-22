package main

func add(a, b []int32, p int) []int32 {
	result := make([]int32, 0)
	var carry int32
        carry = 0
	for i := 0; i < len(a) || i < len(b) || carry > 0; i++ {
		sum := carry
		if i < len(a) {
			sum += a[i]
		}
		if i < len(b) {
			sum += b[i]
		}
		digit := sum % int32(p)
		carry = sum / int32(p)
		result = append(result, digit)
	}

	return result
}

func main() {
}
