package queue

import (
	"fmt"
	"hw5/list"
)

// Queue interface
type Queue interface {
	Enqueue(item interface{})
	Dequeue() (interface{}, error)
	Peek() (interface{}, error)
	IsEmpty() bool
}

// queue struct - simple queue implementation
type queue struct {
	first *list.Node
	last  *list.Node
}

// New creates new queue
func New() Queue {
	return &queue{}
}

// Enqueue adds item to the queue
func (q *queue) Enqueue(item interface{}) {
	n := &list.Node{Data: item}
	if q.first == nil {
		q.first = n
		q.last = n
	} else {
		q.last.Next = n
		q.last = n
	}
}

// Dequeue removes item from the queue
func (q *queue) Dequeue() (interface{}, error) {
	if q.first == nil {
		return nil, fmt.Errorf("The queue is empty")
	}
	r := q.first.Data
	q.first = q.first.Next
	if q.first == nil {
		q.last = nil
	}
	return r, nil
}

// Peek returns current item in the queue
func (q *queue) Peek() (interface{}, error) {
	if q.first == nil {
		return nil, fmt.Errorf("The queue is empty")
	}
	return q.first.Data, nil
}

// IsEmpty checks if queue is empty
func (q *queue) IsEmpty() bool {
	return q.first == nil
}
