package doubly_linked_list

import (
	"errors"
	"fmt"
	"strings"
)

type DoublyLinkedList struct {
	size int
	head *Node
	tail *Node
}

func (linkedList *DoublyLinkedList) PushFront(v interface{}) {
	current := linkedList.head

	node := &Node{v, nil, nil}
	if linkedList.IsEmpty() {
		linkedList.tail = node
	} else {
		current.prev = node
		node.next = current
	}
	linkedList.head = node
	linkedList.size++

}
func (linkedList *DoublyLinkedList) RemoveFront() error {
	if linkedList.IsEmpty() {
		return errors.New("can't remove front element because LinkedList is empty! ")
	}
	if linkedList.size > 1 {
		linkedList.head.next.prev = nil
	} else {
		linkedList.tail = linkedList.head.next
	}

	linkedList.head = linkedList.head.next

	linkedList.size--
	return nil
}

func (linkedList *DoublyLinkedList) PushBack(v interface{}) {
	node := &Node{v, nil, nil}

	if linkedList.IsEmpty() {
		linkedList.head = node

	} else {
		linkedList.tail.next = node
		node.prev = linkedList.tail
	}
	linkedList.tail = node
	linkedList.size++
}

func (linkedList *DoublyLinkedList) PushByIndexBefore(index int, v interface{}) error {
	if index == 0 {
		linkedList.PushFront(v)
		return nil
	}
	err := linkedList.PushByIndex(index, func(current *Node) {
		node := &Node{v, current.prev, current}
		current.prev.next = node
		current.prev = node
	})
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
	return nil
}

func (linkedList *DoublyLinkedList) PushByIndexAfter(index int, v interface{}) error {
	if index == linkedList.size-1 {
		linkedList.PushBack(v)
		return nil
	}
	err := linkedList.PushByIndex(index, func(current *Node) {
		node := &Node{v, current, current.next}
		current.next.prev = node
		current.next = node
	})
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
	return nil
}

func (linkedList *DoublyLinkedList) PushByIndex(index int, callBack func(current *Node)) error {

	length := linkedList.size

	if index > length-1 {
		if linkedList.IsEmpty() {
			return fmt.Errorf("this Index: %d out of range linst(0 - %d)", index, length)
		}
		return fmt.Errorf("this Index: %d out of range linst(0 - %d)", index, length-1)
	}

	current := linkedList.head
	for i := 0; i < length; i++ {
		if i == index {
			// 0 <=> 1 <=> (2) <=> 3       - indexes in list
			// 1 <=> 2 <=>  3  <=> 4 	   - elements in list
			// 1 <=> 2 <=> *5  <=> 3 <=> 4 - add new element in list
			callBack(current)
			linkedList.size++
			break
		}
		current = current.next
	}

	return nil
}

func (linkedList *DoublyLinkedList) RemoveBack() error {
	if linkedList.IsEmpty() {
		return errors.New("can't remove back element because LinkedList is empty! ")
	}
	if linkedList.size > 1 {
		linkedList.tail.prev.next = nil
	} else {
		linkedList.head = linkedList.tail.prev
	}
	linkedList.tail = linkedList.tail.prev

	linkedList.size--
	return nil
}

func (linkedList *DoublyLinkedList) RemoveByValue(value interface{}) error {
	if linkedList.IsEmpty() {
		return errors.New("can't remove by value element because LinkedList is empty! ")
	}

	if linkedList.head.value == value {
		err := linkedList.RemoveFront()
		if err != nil {
			return err
		}
		return nil
	}
	if linkedList.tail.value == value {
		err := linkedList.RemoveBack()
		if err != nil {
			return err
		}
		return nil
	}
	current := linkedList.head

	for current != nil {
		if current.value == value {
			//nil <=> 1 <=> 2 <=> (3) <=> 4 <=> nil
			current.next.prev = current.prev
			current.prev.next = current.next
			linkedList.size--
			break
		}
		current = current.next
	}

	return nil
}

func (linkedList *DoublyLinkedList) Reverse() DoublyLinkedList {
	doublyLinkedList := DoublyLinkedList{}
	current := linkedList.tail
	for current != nil {
		doublyLinkedList.PushBack(current.value)
		current = current.prev
	}
	return doublyLinkedList
}
func (linkedList *DoublyLinkedList) Len() int {
	return linkedList.size
}

func (linkedList *DoublyLinkedList) IsEmpty() bool {
	return linkedList.size < 1
}

func (linkedList *DoublyLinkedList) First() (interface{}, error) {
	if linkedList.IsEmpty() {
		return errors.New("can't get first element because LinkedList is empty! "), nil
	}
	return linkedList.head.value, nil
}

func (linkedList *DoublyLinkedList) Last() (interface{}, error) {
	if linkedList.IsEmpty() {
		return errors.New("can't get last element because LinkedList is empty! "), nil
	}
	return linkedList.tail.value, nil
}

func (linkedList *DoublyLinkedList) ToString() string {
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

func (linkedList *DoublyLinkedList) Print() {
	fmt.Println(linkedList.ToString())
}
