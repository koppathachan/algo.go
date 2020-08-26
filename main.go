package main

import (
	"fmt"
	"github.com/sasidakh/algo.go/bst"
)

type Anode int

func (a Anode) IsLess(node bst.Node) bool {
	n, ok := node.(Anode)
	if !ok {
		return false
	}
	return a < n
}

func main() {
	var arr = []Anode{6, 4, 3, 5, 8, 9, 1}
	root := bst.NewRoot()
	for _, a := range arr {
		fmt.Println("Inserting..", a)
		root.Insert(a)
	}
	fmt.Println(root.InOrder())
	fmt.Println(root.PostOrder())
	fmt.Println(root.PreOrder())
}
