package xmlcutty

import "strings"

type StringFifo struct {
	queue []string
}

func (q *StringFifo) Push(s string) {
	q.queue = append(q.queue, s)
}

func (q *StringFifo) Top() string {
	if len(q.queue) == 0 {
		panic("Top from empty queue")
	}
	return q.queue[len(q.queue)-1]
}

func (q *StringFifo) Pop() string {
	if len(q.queue) == 0 {
		panic("Pop from empty queue")
	}
	r := q.Top()
	q.queue = q.queue[:len(q.queue)-1]
	return r
}

func (q *StringFifo) String() string {
	return "/" + strings.Join(q.queue, "/")
}
