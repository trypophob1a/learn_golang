package doubly_linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOneElementInLinkedList(t *testing.T) {

	linkedList := DoublyLinkedList{}
	linkedList.PushFront(5)
	first, _ := linkedList.First()
	last, _ := linkedList.Last()
	//Test First
	assert.Equal(t, 5, first)
	//Test Last
	assert.Equal(t, 5, last)
}

func TestOneElementInLinkedListError(t *testing.T) {

	linkedList := DoublyLinkedList{}

	_, errorF := linkedList.First()
	_, errorL := linkedList.Last()

	if errorF != nil {
		assert.EqualError(t, errorF, "can't remove front element because LinkedList is empty! ")
	}
	if errorL != nil {
		assert.EqualError(t, errorL, "can't get last element because LinkedList is empty! ")
	}
}

func TestPushFront(t *testing.T) {
	linkedList := DoublyLinkedList{}
	linkedList.PushFront(1)
	linkedList.PushFront(2)
	linkedList.PushFront(3)
	linkedList.PushFront(4)
	linkedList.PushFront(5)
	linkedList.PushFront(6)
	first, _ := linkedList.First()
	last, _ := linkedList.Last()
	assert.Equal(t, 6, first)
	assert.Equal(t, 1, last)
	assert.Equal(t, 3, linkedList.tail.prev.prev.prev.next.value)

}
func TestPushBack(t *testing.T) {
	linkedList := DoublyLinkedList{}
	linkedList.PushBack(1)
	linkedList.PushBack(2)
	linkedList.PushBack(3)
	linkedList.PushBack(4)
	linkedList.PushBack(5)
	linkedList.PushBack(6)
	first, _ := linkedList.First()
	last, _ := linkedList.Last()
	assert.Equal(t, 1, first)
	assert.Equal(t, 6, last)
	assert.Equal(t, 2, linkedList.head.next.next.prev.value)

}

func TestToString(t *testing.T) {
	linkedList := DoublyLinkedList{}
	linkedList.PushFront(1)
	linkedList.PushFront(2)
	linkedList.PushFront(3)

	assert.Equal(t, "[3,2,1]", linkedList.ToString())
}

func TestReverse(t *testing.T) {
	linkedList := DoublyLinkedList{}
	linkedList.PushFront(1)
	linkedList.PushFront(2)
	linkedList.PushFront(3)
	linkedList.PushFront(4)
	linkedList.PushFront(5)
	linkedList.PushFront(6)
	reverse := linkedList.Reverse()
	first, _ := reverse.First()
	last, _ := reverse.Last()
	assert.Equal(t, "[1,2,3,4,5,6]", reverse.ToString())
	assert.Equal(t, 1, first)
	assert.Equal(t, 6, last)

}

func TestRemoveFront(t *testing.T) {
	linkedList := DoublyLinkedList{}
	linkedList.PushFront(1)
	linkedList.PushFront(2)
	linkedList.PushFront(3)
	err := linkedList.RemoveFront()
	if err != nil {
		println(err.Error())
	}
	first, _ := linkedList.First()
	last, _ := linkedList.Last()

	assert.Equal(t, 2, first)
	assert.Equal(t, 1, last)
}

func TestRemoveFrontTwoElement(t *testing.T) {
	linkedList := DoublyLinkedList{}
	linkedList.PushFront(1)
	linkedList.PushFront(2)

	err := linkedList.RemoveFront()
	if err != nil {
		println(err.Error())
	}
	first, _ := linkedList.First()
	last, _ := linkedList.Last()

	assert.Equal(t, 1, first)
	assert.Equal(t, 1, last)
}

func TestRemoveFrontOneElement(t *testing.T) {
	linkedList := DoublyLinkedList{}
	linkedList.PushBack(1)
	_ = linkedList.RemoveFront()
	assert.Equal(t, "[]", linkedList.ToString())
}

func TestRemoveFrontErrorEmpty(t *testing.T) {
	linkedList := DoublyLinkedList{}
	err := linkedList.RemoveFront()
	if err != nil {
		assert.EqualError(t, err, "can't remove front element because LinkedList is empty! ")
	}
}

func TestRemoveBack(t *testing.T) {
	linkedList := DoublyLinkedList{}
	linkedList.PushBack(1)
	linkedList.PushBack(2)
	linkedList.PushBack(3)

	_ = linkedList.RemoveBack()
	first, _ := linkedList.First()
	last, _ := linkedList.Last()

	assert.Equal(t, 1, first)
	assert.Equal(t, 2, last)
}

func TestRemoveBackTwoElement(t *testing.T) {
	linkedList := DoublyLinkedList{}
	linkedList.PushBack(1)
	linkedList.PushBack(2)

	_ = linkedList.RemoveBack()
	first, _ := linkedList.First()
	last, _ := linkedList.Last()

	assert.Equal(t, 1, first)
	assert.Equal(t, 1, last)
}
func TestRemoveBackOneElement(t *testing.T) {
	linkedList := DoublyLinkedList{}
	linkedList.PushBack(1)

	_ = linkedList.RemoveBack()

	assert.Equal(t, "[]", linkedList.ToString())

}

func TestRemoveByValue(t *testing.T) {
	linkedList := DoublyLinkedList{}
	linkedList.PushBack(1)
	linkedList.PushBack(2)
	linkedList.PushBack(3)
	_ = linkedList.RemoveByValue(2)
	first, _ := linkedList.First()
	last, _ := linkedList.Last()
	assert.Equal(t, "[1,3]", linkedList.ToString())
	assert.Equal(t, 1, first)
	assert.Equal(t, 3, last)
}

func TestPushPushByIndexBefore(t *testing.T) {
	linkedList := DoublyLinkedList{}
	linkedList.PushBack(1)
	linkedList.PushBack(2)
	linkedList.PushBack(3)
	linkedList.PushBack(4)
	_ = linkedList.PushByIndexBefore(2, 5)

	assert.Equal(t, "[1,2,5,3,4]", linkedList.ToString())

}

func TestPushByIndexBeforePreviousElement(t *testing.T) {
	linkedList := DoublyLinkedList{}
	linkedList.PushBack(1)
	linkedList.PushBack(2)
	linkedList.PushBack(3)
	linkedList.PushBack(4)
	_ = linkedList.PushByIndexBefore(2, 5)

	assert.Equal(t, 3, linkedList.tail.prev.value)
	assert.Equal(t, 5, linkedList.tail.prev.prev.value)
	assert.Equal(t, 2, linkedList.tail.prev.prev.prev.value)

}
func TestPushByIndexBeforeNextElement(t *testing.T) {
	linkedList := DoublyLinkedList{}
	linkedList.PushBack(1)
	linkedList.PushBack(2)
	linkedList.PushBack(3)
	linkedList.PushBack(4)
	_ = linkedList.PushByIndexBefore(2, 5)

	assert.Equal(t, 2, linkedList.head.next.value)
	assert.Equal(t, 5, linkedList.head.next.next.value)
	assert.Equal(t, 3, linkedList.head.next.next.next.value)
	assert.Equal(t, 4, linkedList.head.next.next.next.next.value)

}

func TestPushByIndexAfter(t *testing.T) {
	linkedList := DoublyLinkedList{}
	linkedList.PushBack(1)
	linkedList.PushBack(2)
	linkedList.PushBack(3)
	linkedList.PushBack(4)
	_ = linkedList.PushByIndexAfter(2, 5)

	assert.Equal(t, "[1,2,3,5,4]", linkedList.ToString())

}

func TestPushByIndexAfterPreviousElement(t *testing.T) {
	linkedList := DoublyLinkedList{}
	linkedList.PushBack(1)
	linkedList.PushBack(2)
	linkedList.PushBack(3)
	linkedList.PushBack(4)
	_ = linkedList.PushByIndexAfter(2, 5)

	assert.Equal(t, 5, linkedList.tail.prev.value)
	assert.Equal(t, 3, linkedList.tail.prev.prev.value)
	assert.Equal(t, 2, linkedList.tail.prev.prev.prev.value)

}

func TestPushByIndexAfterNextElement(t *testing.T) {
	linkedList := DoublyLinkedList{}
	linkedList.PushBack(1)
	linkedList.PushBack(2)
	linkedList.PushBack(3)
	linkedList.PushBack(4)
	_ = linkedList.PushByIndexAfter(2, 5)

	assert.Equal(t, 2, linkedList.head.next.value)
	assert.Equal(t, 3, linkedList.head.next.next.value)
	assert.Equal(t, 5, linkedList.head.next.next.next.value)
	assert.Equal(t, 4, linkedList.head.next.next.next.next.value)

}
