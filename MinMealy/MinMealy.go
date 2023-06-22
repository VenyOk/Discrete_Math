package main

import "fmt"

var parent, rank_ []int

func find(x int) int {
	if parent[x] == x {
		return x
	}
	parent[x] = find(parent[x])
	return parent[x]
}

func merge(x, y int) {
	x = find(x)
	y = find(y)
	if x != y {
		if rank_[x] < rank_[y] {
			t := x
			x = y
			y = t
		}
		parent[y] = x
		if rank_[x] == rank_[y] {
			rank_[x]++
		}
	}
}

type Automata struct {
	states      int
	alph        int
	init        int
	transitions [][]int
	output      [][]string
}

func CreateNewAutomata(states, alph, initialState int) *Automata {
	a := &Automata{
		states,
		alph,
		initialState,
		make([][]int, states),
		make([][]string, states),
	}

	for i := 0; i < states; i++ {
		a.transitions[i] = make([]int, alph)
		a.output[i] = make([]string, alph)
	}

	return a
}

func (automata *Automata) Can_Num() ([]int, int, []bool) {
	visited := make([]bool, automata.states)
	numbering := make([]int, automata.states)
	for i := range numbering {
		numbering[i] = -1
	}

	count := 0

	var dfs func(int)
	dfs = func(index int) {
		numbering[index] = count
		count++
		visited[index] = true
		for i := 0; i < automata.alph; i++ {
			if !visited[automata.transitions[index][i]] {
				dfs(automata.transitions[index][i])
			}
		}
	}
	dfs(automata.init)
	return numbering, count, visited
}

func (automata *Automata) Make_Canonic() *Automata {
	canonized := CreateNewAutomata(automata.states, automata.alph, 0)
	numbering, lastIndex, visited := automata.Can_Num()
	canonized.states = lastIndex
	for i := 0; i < automata.states; i++ {
		if visited[i] && numbering[i] != -1 {
			canonized.output[numbering[i]] = automata.output[i]
			for j := 0; j < automata.alph; j++ {
				canonized.transitions[numbering[i]][j] = numbering[automata.transitions[i][j]]
			}
		}
	}
	return canonized
}

func (automata *Automata) Split1(m int, roots []int) (int, []int) {
	m = automata.states
	parent = make([]int, automata.states)
	rank_ = make([]int, automata.states)
	for i := 0; i < automata.states; i++ {
		parent[i] = i
	}
	for i := 0; i < automata.states; i++ {
		for j := i + 1; j < automata.states; j++ {
			if find(i) != find(j) {
				equal := true
				for k := 0; k < automata.alph; k++ {
					if automata.output[i][k] != automata.output[j][k] {
						equal = false
						break
					}
				}

				if equal {
					merge(i, j)
					m--
				}
			}
		}
	}
	for i := 0; i < automata.states; i++ {
		roots[i] = find(i)
	}

	return m, roots
}

func (automata *Automata) Split(m int, roots []int) (int, []int) {
	m = automata.states
	parent = make([]int, automata.states)
	rank_ = make([]int, automata.states)
	for i := 0; i < automata.states; i++ {
		parent[i] = i
	}
	for i := 0; i < automata.states; i++ {
		for j := i + 1; j < automata.states; j++ {
			if roots[i] == roots[j] && find(i) != find(j) {
				equal := true
				for k := 0; k < automata.alph; k++ {
					if roots[automata.transitions[i][k]] != roots[automata.transitions[j][k]] {
						equal = false
						break
					}
				}
				if equal {
					merge(i, j)
					m--
				}
			}
		}
	}
	for i := 0; i < automata.states; i++ {
		roots[i] = find(i)
	}

	return m, roots
}

func (automata *Automata) Minimization() *Automata {
	roots := make([]int, automata.states)
	var m, m_ int
	m, roots = automata.Split1(m, roots)
	for {
		m_, roots = automata.Split(m_, roots)
		if m == m_ {
			break
		}
		m = m_
	}
	a := make([]int, automata.states)
	b := make([]int, automata.states)
	counter := 0
	for i := 0; i < automata.states; i++ {
		if roots[i] == i {
			a[counter] = i
			b[i] = counter
			counter++
		}
	}
	minimized := CreateNewAutomata(m, automata.alph, b[roots[automata.init]])

	for i := 0; i < minimized.states; i++ {
		for j := 0; j < minimized.alph; j++ {
			minimized.transitions[i][j] = b[roots[automata.transitions[a[i]][j]]]
			minimized.output[i][j] = automata.output[a[i]][j]
		}
	}
	return minimized
}

func get_symbol(j int) rune{
    return rune(97 + j)
}

func answer(automata *Automata) {
    fmt.Println("digraph {")
    fmt.Println("\trankdir = LR")
    for i := 0; i < automata.states; i++ {
        for j := 0; j < automata.alph; j++ {
            fmt.Printf("\t%d -> %d [label = \"%c(%s)\"]\n", i, automata.transitions[i][j], get_symbol(j), automata.output[i][j])
        }
    }
    fmt.Println("}")
}

func main() {
	var states, alph, initialState int
	fmt.Scanf("%d\n%d\n%d", &states, &alph, &initialState)
	automata := CreateNewAutomata(states, alph, initialState)

	for i := 0; i < states; i++ {
		for j := 0; j < alph; j++ {
			fmt.Scan(&automata.transitions[i][j])
		}
	}

	for i := 0; i < states; i++ {
		for j := 0; j < alph; j++ {
			fmt.Scan(&automata.output[i][j])
		}
	}
	answer(automata.Minimization().Make_Canonic())
}
