package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
	finalIndex := len(s.items) - 1

	item := s.items[finalIndex]
	s.items = s.items[:finalIndex]

	return item
}

func main() {
	myStack := Stack{}

	myStack.Push(10)
	myStack.Push(2)
	myStack.Push(23)
	fmt.Println(myStack)

	fmt.Println(myStack.Pop())
	fmt.Println(myStack)

	fmt.Println(myStack.Pop())
	fmt.Println(myStack)
}
