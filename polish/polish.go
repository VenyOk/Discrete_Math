package main

import "fmt"

func Reverse_with_delete_the_brackets(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	res := make([]rune, 0)
	n := 0
	for i := 0; i < len(r); i++ {
		if string(r[i]) != ")" && string(r[i]) != "(" {
			res = append(res, r[i])
			n++
		}
	}
	return string(res)
}

func polish(expression string) (result int) {
	result = 0
	s := Reverse_with_delete_the_brackets(expression)
	st := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if s[i] != 32 {
			if s[i] >= 48 && s[i] <= 57 {
				st = append(st, int(s[i])-'0')
			} else {
				op1 := int(st[len(st)-1])
				st = st[:len(st)-1]
				op2 := st[len(st)-1]
				st = st[:len(st)-1]
				if s[i] == 42 {
					st = append(st, op1*op2)
				} else if s[i] == 43 {
					st = append(st, op1+op2)
				} else if s[i] == 45 {
					st = append(st, op1-op2)
				}
			}
		}
	}
	result = st[0]
	return
}
func main() {
	s := "- 1 * 4 3"
	fmt.Println(polish(s))
}
