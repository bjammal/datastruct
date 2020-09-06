// Package llist provides a type and helper functions to work with single linked list.
// This is not an exhaustive list of helpers.
// We stick to the element type and don't wrap it with a List type.
package llist

import "fmt"

//Node in a list
type Node struct {
	Val  int
	Next *Node
}

//PrintList prints the list of elements in a list
func PrintList(head *Node) {
	runner := head
	for runner != nil {
		fmt.Printf("%d\t", runner.Val)
		runner = runner.Next
	}
	fmt.Printf("\n")
}

//ListLength returns the length of a list
func ListLength(head *Node) int {
	runner := head
	length := 0
	for runner != nil {
		length++
		runner = runner.Next
	}
	return length
}
