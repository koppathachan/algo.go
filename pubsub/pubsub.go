package pubsub

var ques map[string]*[]int = make(map[string]*[]int)

func Subscribe(qname string) chan int {
	c := make(chan int)
	q, ok := ques[qname]
	if !ok {
		var nq []int
		q = &nq
		ques[qname] = q
	}
	go func() {
		for {
			if len(*q) != 0 {
				c <- (*q)[0]
				*q = (*q)[1:]
			}
		}
	}()
	return c
}

func Publish(k int, qname string) {
	q, ok := ques[qname]
	if !ok {
		var nq []int
		q = &nq
		ques[qname] = q
	}
	*q = append(*q, k)
}
