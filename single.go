package main

import (
	"fmt"
)

type LinkedList struct {
	Head *Node
	Size int
}

type Node struct {
	value string
	Next  *Node
}

// Insert - adds a new element to the start of the linked list
func (l *LinkedList) Insert(elem string) {
	node := Node{
		value: elem,
		Next:  l.Head,
	}
	l.Head = &node
	l.Size++
}

// DeleteFirst - removes the first item in the linked list
func (l *LinkedList) DeleteFirst(elem string) {
	l.Head = l.Head.Next
	l.Size--
}

// Search - searchs throught every item in the linked list
// returns the first element that matches the string otherwise
// it returns nil
func (l *LinkedList) Search(elem string) *Node {
	current := l.Head
	for current != nil {
		if current.value == elem {
			return current
		} else {
			current = current.Next
		}
	}
	return nil
}

// Delete - removes an item that matches the string
func (l *LinkedList) Delete(elem string) {
	current := l.Head
	previous := l.Head
	for current != nil {
		if current.value == elem {
			previous.Next = current.Next
			l.Size--
		}
		previous = current
		current = current.Next
	}
}

// List - iterate trhought all items in the linked list and prints
func (l *LinkedList) List() {
	current := l.Head
	for current != nil {
		fmt.Println("%s", current)
		current = current.Next
	}
}

func main() {
	fmt.Println("Single linked list")
	var ll LinkedList
	ll.Insert("one")
	ll.Insert("two")
	ll.Insert("three")
	ll.Delete("two")
	ll.List()
	fmt.Println(ll.Size)
}
