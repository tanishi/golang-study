package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func w(t *tree.Tree, c chan int) {
	if t != nil {
		w(t.Left, c)
		c <- t.Value
		w(t.Right, c)
	}
}
func Walk(t *tree.Tree, ch chan int) {

	w(t, ch)

	close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)

	go Walk(t1, c1)
	go Walk(t2, c2)

	for v1 := range c1 {
		v2 := <-c2
		if v1 != v2 {
			return false
		}
	}

	_, ok := <-c2
	if ok {
		return false
	}

	return true
}

func main() {
	c := make(chan int, 10)
	go Walk(tree.New(1), c)

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
