package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() int {
	item := q.items[0]
	q.items = q.items[1:]

	return item
}

func main() {
	myQueue := Queue{}

	myQueue.Enqueue(10)
	myQueue.Enqueue(2)
	myQueue.Enqueue(23)
	fmt.Println(myQueue)

	fmt.Println(myQueue.Dequeue())
	fmt.Println(myQueue)

	fmt.Println(myQueue.Dequeue())
	fmt.Println(myQueue)
}