package linkedlist

import (
	"bytes"
	"fmt"
)

// Node contains data (and usually a value or a pointer to a value) and a pointer to the next node
type Node struct {
	next *Node
	data int
}

// LinkedList with a single pointer, https://en.wikipedia.org/wiki/Linked_list
type LinkedList struct {
	head *Node
}

// AppendValue is a helper function that can take a value and append it directly
func (list *LinkedList) AppendValue(n int) {
	temp := Node{data: n}
	list.Append(&temp)
}

// Append adds a node to the end of the list
func (list *LinkedList) Append(n *Node) {
	if list.head == nil {
		list.head = n
	} else {
		var current *Node
		for current = list.head; current.next != nil; current = current.next {
		}
		current.next = n
	}
}

// Length counts all the nodes in a linked list
func (list *LinkedList) Length() int {
	var current *Node
	count := 0
	for current = list.head; current != nil; current = current.next {
		count++
	}
	return count
}

// Display returns all of the nodes in a linked list
func (list *LinkedList) Display() string {
	var current *Node
	var b bytes.Buffer
	for current = list.head; current != nil; current = current.next {
		b.WriteString(fmt.Sprintf("  %v\n", current))
	}
	return b.String()
}

// Find returns the first node that has a matching key
func (list *LinkedList) Find(target int) *Node {
	var current *Node
	for current = list.head; current != nil; current = current.next {
		if current.data == target {
			return current
		}
	}
	return nil
}

// Get the "Nth" node from the list using a "zero" index like an array
func (list *LinkedList) Get(index int) *Node {
	var current *Node
	var i int
	for current = list.head; current != nil; current = current.next {
		if i == index {
			return current
		}
		i++
	}
	return nil
}
