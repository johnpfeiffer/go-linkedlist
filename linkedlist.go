package linkedlist

import (
	"bytes"
	"fmt"
	"strings"
)

// Node contains data (and usually a value or a pointer to a value) and a pointer to the next node
type Node struct {
	next *Node
	Data int
}

// LinkedList with a single pointer, https://en.wikipedia.org/wiki/Linked_list
type LinkedList struct {
	Head *Node
}

// AppendValue is a helper function that can take a value and append it directly
func (list *LinkedList) AppendValue(n int) {
	list.Append(&Node{Data: n})
}

// Append adds a node to the end of the list
func (list *LinkedList) Append(n *Node) {
	if list.Head == nil {
		list.Head = n
		return
	}
	var current *Node
	for current = list.Head; current.next != nil; current = current.next {
	}
	current.next = n
}

// PrependValue is a helper function that can take a value and prepend it directly
func (list *LinkedList) PrependValue(n int) {
	list.Prepend(&Node{Data: n})
}

// Prepend adds a node to the beginning of the list
func (list *LinkedList) Prepend(n *Node) {
	if list.Head == nil {
		list.Head = n
		return
	}
	old := list.Head
	list.Head = n
	n.next = old
	return
}

// InsertValue is a helper function that can take a value and insert it directly
func (list *LinkedList) InsertValue(n, location int) {
	list.Insert(&Node{Data: n}, location)
}

// Insert adds a node to a specific (0 based index) location in the list
// a location beyond the list length is treated like an append
func (list *LinkedList) Insert(n *Node, location int) {
	if location == 0 {
		list.Prepend(n)
		return
	}
	if location >= list.Length() {
		fmt.Println("WARNING: treating insert beyond the end of the list as append")
		list.Append(n)
		return
	}
	count := 1
	previous := list.Head
	for current := list.Head.next; current != nil; current = current.next {
		if count == location {
			previous.next = n
			n.next = current
			return
		}
		previous = current
		count++
	}
}

// Length counts all the nodes in a linked list
func (list *LinkedList) Length() int {
	var current *Node
	count := 0
	for current = list.Head; current != nil; current = current.next {
		count++
	}
	return count
}

// Display returns all of the nodes in a linked list
func (list *LinkedList) Display() string {
	var current *Node
	var b bytes.Buffer
	for current = list.Head; current != nil; current = current.next {
		b.WriteString(fmt.Sprintf("  %v\n", current))
	}
	return b.String()
}

// Values returns all of the values in a linked list
func (list *LinkedList) Values() string {
	var current *Node
	var b bytes.Buffer
	for current = list.Head; current != nil; current = current.next {
		b.WriteString(fmt.Sprintf(" %d", current.Data))
	}
	return strings.TrimSpace(b.String())
}

// Find returns the first node that has a matching key
func (list *LinkedList) Find(target int) *Node {
	var current *Node
	for current = list.Head; current != nil; current = current.next {
		if current.Data == target {
			return current
		}
	}
	return nil
}

// Get the "Nth" node from the list using a "zero" index like an array
func (list *LinkedList) Get(index int) *Node {
	var current *Node
	var i int
	for current = list.Head; current != nil; current = current.next {
		if i == index {
			return current
		}
		i++
	}
	return nil
}

// Reduce removes 1 node from the end of the list
func (list *LinkedList) Reduce() {
	if list.Head == nil {
		return
	}
	if list.Head.next == nil {
		list.Head = nil
		return
	}
	var current *Node
	previous := list.Head
	for current = list.Head; current.next != nil; current = current.next {
		previous = current
	}
	// Note that other callers may continue to hold the reference
	previous.next = nil
}

// Delete removes a node given a provided index (using a zero index system)
func (list *LinkedList) Delete(index int) error {
	if index < 0 {
		return fmt.Errorf("Cannot remove an index that is less than zero")
	}
	if index > list.Length() {
		return fmt.Errorf("Index %d is out of range of the length of the list: %d", index, list.Length())
	}
	if index == 0 {
		if list.Head.next != nil {
			list.Head = list.Head.next
		} else {
			list.Head = nil
		}
		return nil
	}
	previous := list.Head
	current := list.Head.next
	count := 1
	for {
		if count == index {
			previous.next = current.next
			return nil
		}
		previous = current
		current = current.next
		count++
	}
}

// ReverseEasy changes the ordering of the nodes so the tail becomes the head and the head becomes the last item
func (list *LinkedList) ReverseEasy() {
	if list.Head == nil || list.Head.next == nil {
		return
	}
	a := make([]int, list.Length())
	current := list.Head
	for i := list.Length() - 1; i >= 0; i-- {
		a[i] = current.Data
		current = current.next
	}

	i := 0
	for current = list.Head; current != nil; current = current.next {
		current.Data = a[i]
		i++
	}
}

// Reverse changes the links of the nodes that eventually the head becomes the last item
// For each node we need a pointer to the previous node
// we need to preserve node.next to continue iterating
// then node.next can point to the previous
func (list *LinkedList) Reverse() {
	var previous *Node
	for current := list.Head; current != nil; {
		temp := current.next
		current.next = previous
		previous = current
		current = temp
	}
	list.Head = previous
}
