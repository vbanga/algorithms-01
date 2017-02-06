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
	union(2, 7)

	union(3, 4)
	union(4, 9)
	union(3, 8)

	fmt.Println(id)
	fmt.Printf("Check for connectivity indices 3 & 8: %s", connected(3, 8))
}

func union(a int, b int) {
	val1 := id[a]
	val2 := id[b]

	if val1 == val2 {
		return
	}

	for i := 0; i < 10; i++ {
		if id[i] == val2 {
			id[i] = val1
		}
	}
}

func connected(a int, b int) bool {
	return id[a] == id[b]
}
