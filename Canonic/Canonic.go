package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
)

func dfs(state int, visited []bool, delta [][]int, transitions *[]int) {
    visited[state] = true
    *transitions = append(*transitions, state)
    for _, nextState := range delta[state] {
        if !visited[nextState] {
            dfs(nextState, visited, delta, transitions)
        }
    }
}

func readInt(scanner *bufio.Scanner) int {
    scanner.Scan()
    num, _ := strconv.Atoi(scanner.Text())
    return num
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)
    n := readInt(scanner)
    m := readInt(scanner)
    q0 := readInt(scanner)
    
    delta := make([][]int, n)
    for i := 0; i < n; i++ {
    	delta[i] = make([]int, m)
    	for j := 0; j < m; j++ {
    		delta[i][j] = readInt(scanner)
    	}
    }
    
    phi := make([][]string, n)
    for i := 0; i < n; i++ {
    	phi[i] = make([]string, m)
    	for j := 0; j < m; j++ {
    		scanner.Scan()
    		phi[i][j] = scanner.Text()
    	}
    }
    
    visited := make([]bool, n)
    transitions := make([]int, 0)
    dfs(q0, visited, delta, &transitions)
    
    canonicalNumbers := make(map[int]int)
    for i, state := range transitions {
    	canonicalNumbers[state] = i
    }
    
    builder := &strings.Builder{}
    builder.WriteString(fmt.Sprintf("%d\n%d\n%d\n", len(transitions), m, canonicalNumbers[q0]))
    for _, state := range transitions {
    	for j := 0; j < m; j++ {
    		nextState := delta[state][j]
    		builder.WriteString(fmt.Sprintf("%d ", canonicalNumbers[nextState]))
    	}
    	builder.WriteString("\n")
    }
    for _, state := range transitions {
    	for j := 0; j < m; j++ {
    		builder.WriteString(fmt.Sprintf("%s ", phi[state][j]))
    	}
    	builder.WriteString("\n")
    }
    fmt.Print(builder.String())
}
