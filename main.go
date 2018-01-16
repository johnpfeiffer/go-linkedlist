package linkedlist

/*
package main

import (
	"fmt"

	"bitbucket.org/johnpfeiffer/linkedlist"
)

func main() {
	a := linkedlist.Node{Data: 1}
	list := linkedlist.LinkedList{}
	displayListInfo(list)

	list.Append(&a)
	displayListInfo(list)
	list.AppendValue(2)
	displayListInfo(list)
	list.AppendValue(3)
	displayListInfo(list)

	fmt.Printf("\ngetting the first item (index 0): %v \n", list.Get(0))
	fmt.Printf("getting second item: %v \n", list.Get(1))
	fmt.Printf("getting third item: %v \n", list.Get(2))
	fmt.Printf("getting a node using an out of bounds index: %v \n", list.Get(3))

	fmt.Printf("\nfinding value 1: %v \n", list.Find(1))
	fmt.Printf("finding value 3: %v \n", list.Find(3))
	fmt.Printf("finding value 9 (does not exist): %v \n", list.Find(9))

	fmt.Printf("\nremoving nodes 3 and 2")
	list.Reduce()
	list.Reduce()
	displayListInfo(list)

	fmt.Printf("\nremoving the last node")
	list.Reduce()
	displayListInfo(list)

	fmt.Printf("\nremoving a non-existent node \n")
	list.Reduce()
	displayListInfo(list)

	fmt.Printf("\nInserting values \n")
	list.InsertValue(101, 0)
	list.InsertValue(104, 9)
	list.InsertValue(100, 0)
	list.InsertValue(102, 2)
	list.InsertValue(103, 3)
	displayListInfo(list)
	fmt.Println(list.Values())

	fmt.Println("Deleting odd values")
	list.Delete(1)
	list.Delete(2)
	fmt.Println(list.Values())

	fmt.Println("done")
}

func displayListInfo(list linkedlist.LinkedList) {
	fmt.Printf("\n%d nodes, head: %v\n", list.Length(), list.Head)
	fmt.Printf(list.Display())
}
*/
