package lesson2

type Node struct {
	next  *Node
	value int
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (l *LinkedList) AddInTail(item Node) {
	if l.head == nil {
		l.head = &item
	} else {
		l.tail.next = &item
	}
	l.tail = &item
}

func GetLinkedList(values []int) *LinkedList {
	// 6.2
	// old name: resultLL
	// new name: linkedList
	var linkedList LinkedList // resulting linked list
	for _, value := range values {
		linkedList.AddInTail(Node{
			value: value,
		})
	}
	return &linkedList
}
