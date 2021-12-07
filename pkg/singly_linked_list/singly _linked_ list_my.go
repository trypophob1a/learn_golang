package singly_linked_list

import (
	"fmt"
	"strings"
)

type SinglyLinkedList struct {
	size int
	head *Node
	tail *Node
}

func (linkedList *SinglyLinkedList) PushFront(v interface{}) {
	node := &Node{v, nil}
	if linkedList.IsEmpty() {
		linkedList.tail = node
	} else {
		current := linkedList.head
		node.next = current
	}

	linkedList.head = node
	linkedList.size++
}

func (linkedList *SinglyLinkedList) PushBack(v interface{}) {
	node := &Node{v, nil}
	if linkedList.IsEmpty() {
		linkedList.head = node
	} else {
		linkedList.tail.next = node
	}
	linkedList.tail = node
	linkedList.size++
}

func (linkedList *SinglyLinkedList) RemoveFront() error {

	if linkedList.IsEmpty() {
		return fmt.Errorf("can't Remove Front list is empty! ")
	}

	linkedList.head = linkedList.head.next
	linkedList.size--
	return nil
}

func (linkedList *SinglyLinkedList) RemoveBack() error {

	if linkedList.IsEmpty() {
		return fmt.Errorf("can't Remove Back list is empty! ")
	}

	if linkedList.size == 1 {
		linkedList.tail = nil
		linkedList.head = nil
		linkedList.size--
		return nil
	}
	current := linkedList.head
	for current.next != nil {
		linkedList.tail = current
		current = current.next
	}

	linkedList.tail.next = nil
	linkedList.size--
	return nil
}

func (linkedList *SinglyLinkedList) IsEmpty() bool {
	return linkedList.size == 0
}

func (linkedList *SinglyLinkedList) Len() int {
	return linkedList.size
}

func (linkedList *SinglyLinkedList) First() (interface{}, error) {
	if linkedList.IsEmpty() {
		return nil, fmt.Errorf("can't get first list is empty! ")
	}
	return linkedList.head.value, nil
}

func (linkedList *SinglyLinkedList) Last() (interface{}, error) {
	if linkedList.IsEmpty() {
		return nil, fmt.Errorf("can't get last list is empty! ")
	}
	return linkedList.head.value, nil
}

func (linkedList *SinglyLinkedList) ToString() string {
	sb := strings.Builder{}
	sb.WriteString("[")
	data := linkedList.head
	for data != nil {

		sb.WriteString(fmt.Sprintf("%v,", data.value))
		data = data.next
	}
	list := strings.Trim(sb.String(), ",") + "]"
	return list
}
