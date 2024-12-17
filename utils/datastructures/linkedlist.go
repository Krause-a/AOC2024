package datastructures

import "fmt"

type LinkedList[T any] struct {
	Size int
	Head *Node[T]
}

type Node[T any] struct {
	Value T
	Next *Node[T]
}

func MakeLinkedList[T any](inputs []T) LinkedList[T] {
	list := LinkedList[T] {
		Size: len(inputs),
	}

	var previous *Node[T]
	for i, v := range inputs {
		n := Node[T] {
			Value: v,
		}
		if i != 0 {
			previous.Next = &n
		}
		previous = &n
		if i == 0 {
			list.Head = &n
		}
	}

	return list
}

func (list *LinkedList[T]) PrintList() {
	n := list.Head
	for n != nil {
		fmt.Printf("%v", n.Value)
		n = n.Next
		if n != nil {
			fmt.Printf(" -> ")
		}
	}
	fmt.Println()
}
