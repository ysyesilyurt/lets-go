package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	i, j := 0, 1
	fun := func() int {
		next := i + j
		i, j = j, next
		return i
	}
	return fun
}

func main() {
	n := 70
	f := fibonacci()
	for i := 0; i < n; i++ {
		fmt.Println(f())
	}
}
