package main

import "fmt"

var id [10]int

func main() {
	for i := 0; i < 10; i++ {
		id[i] = i
	}

	fmt.Println(id)

	union(0, 5)
	union(5, 6)

	union(1, 2)
	union(1, 0)
	union(1, 7)

	union(8, 3)
	union(3, 4)
	union(8, 9)

	// union(8, 1)

	fmt.Println(id)
	fmt.Printf("Check for connectivity indices 3 & 8: %s\n", connected(3, 8))
	fmt.Printf("Check for connectivity indices 3 & 8: %s\n", connected(5, 9))
}

func root(a int) int {
	for {
		if id[a] == a {
			return a
		} else {
			a = id[a]
		}
	}
}

func connected(a int, b int) bool {
	return root(a) == root(b)
}

func union(a int, b int) {
	i := root(a)
	j := root(b)

	id[j] = i
}
