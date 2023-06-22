package main

import (
    "fmt"
    "bufio"
    "os"
)

func in_array(arr []string, s string) (result bool) {
	result = false
	for i := 0; i < len(arr); i++ {
		if s == arr[i] {
			result = true
			break
		}
	}
	return
}

func econom(s string) (result int) {
	var used_oper []string
	var braces []int
	result = 0
	for i, symb := range s {
		if string(symb) == "(" {
			braces = append(braces, i) 
		} else if string(symb) == ")" {
			if !in_array(used_oper, s[braces[len(braces)-1]:i]) {
				used_oper = append(used_oper, s[braces[len(braces)-1]:i])
			}
			braces = braces[:len(braces)-1]
		}
	}
	result = len(used_oper)
	return
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    s := scanner.Text()
    fmt.Println(econom(s))
}
