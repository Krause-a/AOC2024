package datastructures

type LinkedList struct {
	Size int
	Head *Node
}

type Node struct {
	Value int // int for now. Generics when needed
	Next *Node
}

func MakeLinkedList(inputs []int) LinkedList {
	list := LinkedList {
		Size: len(inputs),
	}

	var previous *Node
	for i, v := range inputs {
		n := Node {
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
