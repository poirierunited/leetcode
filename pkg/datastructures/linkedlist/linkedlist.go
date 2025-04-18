package linkedlist

// Node represents a node in the linked list
type Node struct {
	Value int
	Next  *Node
}

// LinkedList represents a singly linked list
type LinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

// NewLinkedList creates a new empty linked list
func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head:   nil,
		Tail:   nil,
		Length: 0,
	}
}

// Add to the end of the linked list
func (ll *LinkedList) Add(value int) {
	newNode := &Node{Value: value}

	if ll.Head == nil {
		ll.Head = newNode
		ll.Tail = newNode
	} else {
		ll.Tail.Next = newNode
		ll.Tail = newNode
	}

	ll.Length++
}

// Delete last element
func (ll *LinkedList) Pop() {
	if ll.Tail == nil {
		return
	}

	if ll.Length == 1 {
		ll.Head = nil
		ll.Tail = nil
		ll.Length--
		return
	}

	// get preultimate Node
	prt := ll.Head
	for prt.Next != ll.Tail {
		prt = prt.Next
	}

	ll.Tail = prt
	ll.Tail.Next = nil
	ll.Length--
}
