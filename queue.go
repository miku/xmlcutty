package xmlcutty

import (
	"strings"
	"sync"
)

type StringFifo struct {
	mu    sync.RWMutex
	queue []string
}

func (q *StringFifo) Push(s string) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.queue = append(q.queue, s)
}

func (q *StringFifo) Top() string {
	if len(q.queue) == 0 {
		panic("pop from empty queue")
	}
	return q.queue[len(q.queue)-1]
}

func (q *StringFifo) Pop() string {
	if len(q.queue) == 0 {
		panic("pop from empty queue")
	}
	var u []string
	for _, s := range q.queue[:len(q.queue)-1] {
		u = append(u, s)
	}
	r := q.Top()
	q.mu.Lock()
	defer q.mu.Unlock()
	q.queue = u
	return r
}

func (q *StringFifo) String() string {
	return "/" + strings.Join(q.queue, "/")
}
