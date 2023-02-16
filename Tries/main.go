package main

import "fmt"

const ALPHABET_SIZE = 26

type Node struct {
	Children [ALPHABET_SIZE]*Node
	IsEnd    bool
}

type Trie struct {
	Root *Node
}

func InitTrie() *Trie {
	return &Trie{Root: &Node{}}
}

func (t *Trie) Insert(word string) {
	wordLength := len(word)
	currentNode := t.Root

	for i := 0; i < wordLength; i++ {
		charIndex := word[i] - 'a'

		if currentNode.Children[charIndex] == nil {
			currentNode.Children[charIndex] = &Node{}
		}

		currentNode = currentNode.Children[charIndex]
	}

	currentNode.IsEnd = true
}

func (t *Trie) Search(word string) bool {
	wordLength := len(word)
	currentNode := t.Root

	for i := 0; i < wordLength; i++ {
		charIndex := word[i] - 'a'

		if currentNode.Children[charIndex] == nil {
			return false
		}

		currentNode = currentNode.Children[charIndex]
	}

	if currentNode.IsEnd == true {
		return true
	}

	return false
}

func main() {
	myTrie := InitTrie()

	toInsert := []string{
		"aragorn",
		"aragon",
		"argon",
		"eragon",
		"oregon",
		"oregano",
		"oreo",
	}

	for _, v := range toInsert {
		myTrie.Insert(v)
	}

	fmt.Println(myTrie.Search("orc"))

	myTrie.Insert("orc")
	fmt.Println(myTrie.Search("orc"))
}
