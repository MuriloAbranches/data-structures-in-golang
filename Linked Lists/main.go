package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head   *Node
	Length int
}

func (l *LinkedList) Prepend(node *Node) {
	second := l.Head
	l.Head = node
	l.Head.Next = second
	l.Length++
}

func (l LinkedList) PrintData() {
	toPrint := l.Head

	for l.Length != 0 {
		fmt.Printf("%d ", toPrint.Data)

		toPrint = toPrint.Next
		l.Length--
	}

	fmt.Printf("\n")
}

func (l *LinkedList) DeleteWithValue(value int) {
	if l.Length == 0 {
		return
	}

	if l.Head.Data == value {
		l.Head = l.Head.Next
		l.Length--
		return
	}

	previousToDelete := l.Head

	for previousToDelete.Next.Data != value {
		if previousToDelete.Next.Next == nil {
			return
		}

		previousToDelete = previousToDelete.Next
	}

	previousToDelete.Next = previousToDelete.Next.Next
	l.Length--
}

func main() {
	myLinkedList := LinkedList{}
	node1 := &Node{Data: 38}
	node2 := &Node{Data: 13}
	node3 := &Node{Data: 150}
	node4 := &Node{Data: 10}

	myLinkedList.Prepend(node1)
	myLinkedList.Prepend(node2)
	myLinkedList.Prepend(node3)
	myLinkedList.Prepend(node4)

	fmt.Println(myLinkedList)
	myLinkedList.PrintData()

	myLinkedList.DeleteWithValue(13)
	myLinkedList.PrintData()

	myLinkedList.DeleteWithValue(100)
	myLinkedList.PrintData()

	myLinkedList.DeleteWithValue(10)
	myLinkedList.PrintData()

	emptyList := LinkedList{}
	emptyList.DeleteWithValue(15)
	emptyList.PrintData()
}
