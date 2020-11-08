package main

import "golang.org/x/tour/tree"
import "fmt"

// In Order Tree Traversal
func WalkInOrder(t *tree.Tree, ch chan int) {
	if t != nil {
		WalkInOrder(t.Left, ch)
		ch <- t.Value
		WalkInOrder(t.Right, ch)
	}
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	WalkInOrder(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	t1_contents, t2_contents := make([]int, 10), make([]int, 10)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		select {
		case t1_val, ok := <-ch1:
			if !ok {
				ch1 = nil
			} else {
				t1_contents = append(t1_contents, t1_val)
			}
		case t2_val, ok := <-ch2:
			if !ok {
				ch2 = nil
			} else {
				t2_contents = append(t2_contents, t2_val)
			}
		}
		if ch1 == nil && ch2 == nil {
			break
		}
	}
	for i, _ := range t1_contents {
		if t1_contents[i] != t2_contents[i] {
			return false
		}
	}
	return true
}

func main() {
	t1, t2 := tree.New(10), tree.New(1)
	fmt.Println(Same(t1, t2))
	/*
		// InOrder traversal test
		ch := make(chan int)
		t1 := tree.New(1)
		go Walk(t1, ch)
		for i := range ch {
			fmt.Println(i)
		}
	*/
}
