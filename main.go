package main

import (
	"fmt"
	"github.com/sasidakh/algo.go/util"
)

func main() {
	var list = util.IntArray{1, 2, 3, 4, 5, 6, 7}
	var fn = func(a int) int {
		return a + a
	}

	fmt.Println(list)
	fmt.Println(list.Len())
	fmt.Println(list.Map(fn))
}
