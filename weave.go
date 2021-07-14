package main

import "fmt"

func main() {
	letters := []int{1, 2, 3, 4}
	// For each letter
	res := [][]int{{letters[0]}}
	for i := 1; i < len(letters); i++ {
		letter := letters[i]
		fmt.Printf("considering letter: %v\n", letter)
		// For each combination
		newRes := make([][]int, 0)
		for _, comb := range res {
			newRes = append(newRes, weave(comb, letter)...)
		}
		// fmt.Printf("newRes %v\n", newRes)
		// copy(res, newRes)
		res = newRes
	}
	fmt.Printf("res %v\n", res)
}

func weave(letters []int, letter int) [][]int {
	appLetters := make([][]int, 0)
	for i := 0; i <= len(letters); i++ {
		res := insert(letters, i, letter)
		appLetters = append(appLetters, res)
		fmt.Printf("appLetters %v\n", appLetters)
	}
	return appLetters
}

func insert(slice []int, index int, value int) []int {
	// fmt.Printf("val %v\n", value)
	tmp := make([]int, len(slice))
	copy(tmp, slice)
	// Grow the slice by one element.
	// fmt.Printf("ls: %v\n", len(tmp))
	tmp = append(tmp, -1)
	// fmt.Printf("ls: %v\n", len(tmp))
	tmp = tmp[:len(tmp)]
	// Use copy to move the upper part of the slice out of the way and open a hole.
	copy(tmp[index+1:], tmp[index:])
	// Store the new value.
	tmp[index] = value
	// fmt.Printf("slice %v\n", slice)
	// fmt.Printf("tmp %v\n", tmp)
	// Return the result.
	return tmp
}
