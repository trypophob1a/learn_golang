package singly_linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPushFront(t *testing.T) {

	linkedList := SinglyLinkedList{}
	linkedList.PushFront(1)
	linkedList.PushFront(2)

	//Test First
	assert.Equal(t, 2, linkedList.head.value)
	//Test Last
	assert.Equal(t, 1, linkedList.tail.value)
}

func TestPushBack(t *testing.T) {

	linkedList := SinglyLinkedList{}
	linkedList.PushBack(1)
	linkedList.PushBack(2)

	//Test First
	assert.Equal(t, 1, linkedList.head.value)
	//Test Last
	assert.Equal(t, 2, linkedList.tail.value)
}

func TestRemoveFront(t *testing.T) {

	linkedList := SinglyLinkedList{}
	linkedList.PushBack(1)
	linkedList.PushBack(2)
	linkedList.PushBack(3)
	_ = linkedList.RemoveFront()

	assert.Equal(t, 2, linkedList.head.value)
	assert.Equal(t, 3, linkedList.tail.value)
}

func TestRemoveBack(t *testing.T) {

	linkedList := SinglyLinkedList{}
	linkedList.PushBack(1)
	linkedList.PushBack(2)
	linkedList.PushBack(3)
	_ = linkedList.RemoveBack()
	assert.Equal(t, 1, linkedList.head.value)
	assert.Equal(t, 2, linkedList.tail.value)
	assert.Equal(t, "[1,2]", linkedList.ToString())
}

func TestPushBackAfterRemoveBack(t *testing.T) {

	linkedList := SinglyLinkedList{}
	linkedList.PushBack(1)
	linkedList.PushBack(2)
	linkedList.PushBack(3)
	_ = linkedList.RemoveBack()
	linkedList.PushBack(4)
	assert.Equal(t, 1, linkedList.head.value)
	assert.Equal(t, 4, linkedList.tail.value)
	assert.Equal(t, "[1,2,4]", linkedList.ToString())
}

func TestRemoveBackOneElement(t *testing.T) {

	linkedList := SinglyLinkedList{}
	linkedList.PushBack(1)
	_ = linkedList.RemoveBack()

	_, errF := linkedList.First()
	assert.EqualError(t, errF, "can't get first list is empty! ")
	_, errL := linkedList.Last()
	assert.EqualError(t, errL, "can't get last list is empty! ")

	assert.Equal(t, "[]", linkedList.ToString())
}
