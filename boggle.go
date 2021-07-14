package main

import "fmt"

type node struct {
	complete bool
	letter   byte
	next     []*node
}

func main() {
	dictionary := []string{"ale", "all", "al", "ales", "a", "aliases", "bobward"}
	letters := []byte{'a', 'e', 'l', 's', 'i', 'a', 's', 's', 'b'}
	// correct answer is “aliases”
	head := &node{complete: false}
	for _, word := range dictionary {
		buildTree(head, word, 0)
	}
	letterHash := buildLetterHash(letters)
	// dict := walkTree(head)
	// fmt.Printf("dict len - %v\n", len(dict))
	// walkLetterHash(letterHash)
	longestWord := findLongestWord(head, letterHash)
	fmt.Printf("longestWord: %v\n", longestWord)
}

func findLongestWord(head *node, letterHash map[byte]int) string {
	var results []string
	f := func(s string) {
		results = append(results, s)
	}
	findWord(head, "", letterHash, f)
	longestWord := ""
	for _, word := range results {
		fmt.Printf("Considering word %v\n", word)
		if len(longestWord) < len(word) {
			longestWord = word
		}
	}
	return longestWord
}

func findWord(head *node, current string, letterHash map[byte]int, f func(string)) {
	// If we have a valid letter, concat to current word
	if head.letter != 0 {
		current = current + string(head.letter)
	}
	// Generate used hash
	usedHash := buildLetterHash([]byte(current))
	// If we do not have this letter available, return
	if usedHash[head.letter] > letterHash[head.letter] && head.letter != 0 {
		// walkLetterHash(letterHash)
		return
	}
	// If we are at end of a word, return it
	if head.complete {
		// fmt.Printf("longest: %v-%v, curr: %v-%v ", len(longest), longest, len(current), current)
		// fmt.Printf("new word: %v - len %v\n", current, len(words))
		f(current)
	}
	for _, val := range head.next {
		findWord(val, current, letterHash, f)
	}
}

func buildLetterHash(letters []byte) map[byte]int {
	letterHash := make(map[byte]int)
	for _, letter := range letters {
		if _, ok := letterHash[letter]; ok {
			letterHash[letter]++
		} else {
			letterHash[letter] = 1
		}
	}
	return letterHash
}

func buildTree(head *node, word string, index int) error {
	// If index is the length of the word -> base case return out
	if len(word) == index {
		return nil
	}
	// Determine if node contains next char
	foundHead := false
	for _, nextNode := range head.next {
		// If next node exists, set it to current head
		if nextNode.letter == word[index] {
			head = nextNode
			foundHead = true
		}
	}
	// If we did not find the letter, add new node to head
	if !foundHead {
		newNode := &node{letter: word[index], complete: false}
		head.next = append(head.next, newNode)
		// AND make that the new head
		head = newNode
	}
	index++
	// If we are at end of word update node for complete
	if len(word) == index {
		head.complete = true
	}
	// Move to the next letter in word
	buildTree(head, word, index)
	return nil
}

// Traverses entire tree, and prints all words, useful for debugging
func treeWalker(head *node, current string, ch chan string) {
	current = current + string(head.letter)
	if head.complete {
		ch <- current
		fmt.Printf("walking: %v\n", current)
	}
	for _, node := range head.next {
		treeWalker(node, current, ch)
	}
}

func walkTree(head *node) []string {
	c := make(chan string)
	go func() {
		treeWalker(head, "", c)
		close(c)
	}()
	dict := make([]string, 0)
	for {
		v, ok := <-c
		if ok {
			dict = append(dict, v)
		} else {
			break
		}
	}
	return dict
}

// Traverses a hash, printing letters & counts, useful for debugging
func walkLetterHash(letterHash map[byte]int) {
	for letter, count := range letterHash {
		fmt.Printf("Letter: %v - Count: %v\n", string(letter), count)
	}
}
