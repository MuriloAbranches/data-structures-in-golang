package main

import "fmt"

const ARRAY_SIZE = 7

type HashTable struct {
	Array [ARRAY_SIZE]*Bucket
}

type Bucket struct {
	Head *BucketNode
}

type BucketNode struct {
	Key  string
	Next *BucketNode
}

func Init() *HashTable {
	result := &HashTable{}

	for i := range result.Array {
		result.Array[i] = &Bucket{}
	}

	return result
}

func Hash(key string) int {
	sum := 0

	for _, v := range key {
		sum += int(v)
	}

	return sum % ARRAY_SIZE
}

func (b *Bucket) Insert(key string) {
	if !b.Search(key) {
		newNode := &BucketNode{Key: key}
		newNode.Next = b.Head
		b.Head = newNode
	} else {
		fmt.Println(key, "already exists")
	}
}

func (b *Bucket) Search(key string) bool {
	currentNode := b.Head

	for currentNode != nil {
		if currentNode.Key == key {
			return true
		}
		currentNode = currentNode.Next
	}

	return false
}

func (b *Bucket) Delete(key string) {
	if b.Head.Key == key {
		b.Head = b.Head.Next
		return
	}

	previousNode := b.Head

	for previousNode.Next != nil {
		if previousNode.Next.Key == key {
			previousNode.Next = previousNode.Next.Next
			return
		}

		previousNode = previousNode.Next
	}
}

func (h *HashTable) Insert(key string) {
	index := Hash(key)
	h.Array[index].Insert(key)
}

func (h *HashTable) Search(key string) bool {
	index := Hash(key)
	return h.Array[index].Search(key)
}

func (h *HashTable) Delete(key string) {
	index := Hash(key)
	h.Array[index].Delete(key)
}

func main() {
	myHashTable := Init()
	list := []string{
		"ERIC",
		"KENNY",
		"KYLE",
		"STAN",
		"RANDY",
		"BUTTERS",
		"TOKEN",
	}

	for _, v := range list {
		myHashTable.Insert(v)
	}

	fmt.Println(myHashTable.Search("KENNY"))
	myHashTable.Insert("KENNY")
	myHashTable.Delete("STAN")
	fmt.Println(myHashTable.Search("STAN"))
}
