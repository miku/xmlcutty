package xmlcutty

import "strings"

type StringStack struct {
	queue []string
}

func (q *StringStack) Push(s string) {
	q.queue = append(q.queue, s)
}

func (q *StringStack) Top() string {
	if len(q.queue) == 0 {
		panic("Top from empty queue")
	}
	return q.queue[len(q.queue)-1]
}

func (q *StringStack) Pop() string {
	if len(q.queue) == 0 {
		panic("Pop from empty queue")
	}
	r := q.Top()
	q.queue = q.queue[:len(q.queue)-1]
	return r
}

func (q *StringStack) String() string {
	return "/" + strings.Join(q.queue, "/")
}
