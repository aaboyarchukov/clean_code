package lesson1

import (
	"errors"
	_ "os"
	_ "reflect"
)

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

// task 5
// t = O(n), where n = len(list)
func (l *LinkedList) Count() int {
	// old name: count
	// new name: countAllNodesInLL
	var countAllNodes int

	// old name: tempNode
	// new name: currentNode
	currentNode := l.head

	for currentNode != nil {
		countAllNodes++
		currentNode = currentNode.next
	}
	return countAllNodes
}

func (l *LinkedList) Find(n int) (Node, error) {
	// old name: tempNode
	// new name: currentNode
	currentNode := l.head
	for currentNode != nil {
		if currentNode.value == n {
			return *currentNode, nil
		}
		currentNode = currentNode.next
	}
	return Node{value: -1, next: nil}, errors.New("node is not finding")
}

// task 4
// t = O(n), where n = len(list)
func (l *LinkedList) FindAll(n int) []Node {
	// old name: nodes
	// new name: listOfNodesWithValueN
	var listOfNodesWithValueN []Node

	// old name: tempNode
	// new name: currentNode
	currentNode := l.head
	for currentNode != nil {
		if currentNode.value == n {
			listOfNodesWithValueN = append(listOfNodesWithValueN, *currentNode)
		}
		currentNode = currentNode.next
	}
	return listOfNodesWithValueN
}

// task 1
// t = O(n), where n = len(list)
// task 2
// t = O(n), where n = len(list)
func (l *LinkedList) Delete(n int, all bool) {
	if l.head == nil {
		return
	}

	// old name: tempNode
	// new name: currentNode
	currentNode := l.head

	// old name: prev
	// new name: prevNode
	var prevNode *Node

	if l.Count() == 1 && currentNode.value == n {
		l.Clean()
		return
	}

	for currentNode != nil {
		// old name: deleted
		// new name: isNodeDeleted
		isNodeDeleted := false

		if currentNode.value == n && currentNode == l.head {
			l.head = currentNode.next
			isNodeDeleted = true
		} else if currentNode.value == n && currentNode == l.tail {
			prevNode.next = nil
			l.tail = prevNode
			isNodeDeleted = true
		} else if currentNode.value == n {
			prevNode.next = currentNode.next
			isNodeDeleted = true
		}
		if !all && isNodeDeleted {
			return
		}
		if !isNodeDeleted {
			prevNode = currentNode
		}
		currentNode = currentNode.next
	}
}

// task 6
// t = O(n), where n = len(list)
func (l *LinkedList) Insert(after *Node, add Node) {
	if l.head == nil {
		l.InsertFirst(add)
		return
	}

	// old name: tempNode
	// new name: currentNode
	currentNode := l.head

	// if node will not exists, then we have to finding it first
	for currentNode.value != after.value {
		currentNode = currentNode.next
	}
	if currentNode == l.tail {
		l.AddInTail(add)
	} else {
		nextNode := currentNode.next
		currentNode.next = &add
		add.next = nextNode
	}

}

func (l *LinkedList) InsertFirst(first Node) {
	if l.head == nil {
		l.tail = &first
	} else {
		first.next = l.head
	}
	l.head = &first

}

// task 3
// t = O(1)
func (l *LinkedList) Clean() {
	l.head = nil
	l.tail = nil
}

// func PrintLL(LL *LinkedList) {
// 	temp := LL.head
// 	for temp != nil {
// 		fmt.Printf("%d ", temp.value)
// 		temp = temp.next
// 	}
// 	fmt.Println()
// }

func GetLinkedList(values []int) *LinkedList {
	var resultLL LinkedList // resulting linked list
	for _, value := range values {
		resultLL.AddInTail(Node{
			value: value,
		})
	}
	return &resultLL
}

func EqualLists(l1 *LinkedList, l2 *LinkedList) bool {
	// equals len and elements
	if l1.head == nil &&
		l2.head == nil {
		return true
	}

	if l1.head.value != l2.head.value {
		return false
	}
	if l1.tail.value != l2.tail.value {
		return false
	}

	// old name: countL1, countL2
	// new name: countAllNodesOfL1, countAllNodesOfL2
	countAllNodesOfL1, countAllNodesOfL2 := l1.Count(), l2.Count()

	if countAllNodesOfL1 == countAllNodesOfL2 {
		// old name: tempL1, tempL2
		// new name: currentNodeOfL1, currentNodeOfL2
		currentNodeOfL1, currentNodeOfL2 := l1.head, l2.head
		for currentNodeOfL1 != nil && currentNodeOfL2 != nil {
			if currentNodeOfL1.value != currentNodeOfL2.value {
				return false
			}
			currentNodeOfL1 = currentNodeOfL1.next
			currentNodeOfL2 = currentNodeOfL2.next
		}

		return true
	}

	return false
}
