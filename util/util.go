package util

type IntArray []int

func (arr IntArray) Len() int {
	return len(arr)
}

func (arr IntArray) Map(fn func(int) int) IntArray {
	var res []int = make([]int, arr.Len())
	for i, a := range arr {
		res[i] = fn(a)
	}
	return res
}
