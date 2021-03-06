package main

import "fmt"

func main() {
	// Array declaration
	nums := [2]int {1,2}
	// Array without specifying lenght
	nums2 := [...]int {3,4}
	
	// Slice declaration: slice of string
	// We can define values in {}
	cards := []string{ newCard(), "Ace of Spades"}

	// Adding element. append() does not modify existing slice
	cards = append(cards, "Six of Club")

	// Print cards
	fmt.Println(cards, nums, nums2)

	// Iterate on slice
	for index, card := range cards {
		fmt.Println(index, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
