package main

import "fmt"

type queue []int

func (s queue) Push(value int) queue {
	return append(s, value)
}

func (s queue) Pop() (queue, int) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s queue) Front() int {
	l := len(s)
	return s[l-1]
}

func main() {
	q := make(queue, 0)
	q = q.Push(1)
	q = q.Push(2)
	q = q.Push(3)
	fmt.Printf("Head %d\n", q.Front())
	q, _ = q.Pop()
	fmt.Printf("Head %d\n", q.Front())
	q, _ = q.Pop()
	fmt.Printf("Head %d\n", q.Front())
	q, _ = q.Pop()
}
