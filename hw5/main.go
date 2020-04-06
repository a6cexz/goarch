package main

import (
	"fmt"
	"hw5/queue"
	"hw5/stack"
)

func main() {
	fmt.Println("Test queue")
	q := queue.New()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	for !q.IsEmpty() {
		item, err := q.Dequeue()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(item)
		}
	}

	fmt.Println("Test stack")
	s := stack.New()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	for !s.IsEmpty() {
		item, err := s.Pop()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(item)
		}
	}
}
