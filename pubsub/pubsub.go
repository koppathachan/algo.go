package pubsub

var ques map[string]chan int = make(map[string]chan int)

func Subscribe(qname string) chan int {
	c := make(chan int)
	q, ok := ques[qname]
	if !ok {
		q = make(chan int)
		ques[qname] = q
	}
	go func() {
		for {
			select {
			case x := <-q:
				c <- x
			}
		}
	}()
	return c
}

func Publish(k int, qname string) {
	q, ok := ques[qname]
	if !ok {
		q = make(chan int)
		ques[qname] = q
	}
	q <- k
}
