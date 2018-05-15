package main

import "strings"

// Node is a tree node object
type Node struct {
	PhraseID int
	Children map[string]*Node
}

// NewNode returns a pointer to a new Node object
func NewNode(id int) *Node {
	return &Node{
		PhraseID: id,
		Children: make(map[string]*Node),
	}
}

func addPhrase(root *Node, phrase string, phraseID int) {
	// a pointer to traverse the trie without damaging
	// the original reference
	node := root
	words := strings.Split(phrase, " ")

	for i := 0; i < len(words); i++ {
		word := words[i]
		// if the current word does not exist as a child
		// to current node, add it
		if _, found := node.Children[word]; found == false {
			node.Children[word] = NewNode(-1)
		}
		// move traversal pointer to current word
		node = node.Children[word]

		// if current word is the last one, mark it with
		// phrase Id
		if i == len(words)-1 {
			node.PhraseID = phraseID
		}
	}
}

// FindPhrases is an implementation of the Trie search
// algorithm which looks for matches and returns them as
// an indexed list of ints
func FindPhrases(root *Node, textBody string) []int {
	// a pointer to traverse the trie without damaging
	// the original reference
	node := root
	var foundPhrases []int
	words := strings.Split(textBody, " ")

	// starting traversal at trie root and first
	// word in text body
	for i := 0; i < len(words); {
		word := words[i]
		// if current node has current word as a child
		// move both node and words pointer forward
		if _, found := node.Children[word]; found == true {
			// move trie pointer forward
			node = node.Children[word]
			// move words pointer forward
			i++
		} else {
			// current node does not have current
			// word in its children

			// if there is a phrase Id, then the previous
			// sequence of words matched a phrase, add Id to
			// found list
			if node.PhraseID != -1 {
				foundPhrases = append(foundPhrases, node.PhraseID)
			}

			if node == root {
				// if trie pointer is already at root, increment
				// words pointer
				i++
			} else {
				// if not, leave words pointer at current word
				// and return trie pointer to root
				node = root
			}
		}
	}

	// one case remains, word pointer as reached the end
	// and the loop is over but the trie pointer is pointing to
	// a phrase Id
	if node.PhraseID != -1 {
		foundPhrases = append(foundPhrases, node.PhraseID)
	}
	return foundPhrases
}

func main() {

}
