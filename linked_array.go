package algo_structures

import (
	"fmt"
	"strings"
)

type linkedListItem struct {
	value int
	next  *linkedListItem
}

func (i *linkedListItem) Value() int {
	return i.value
}

type linkedList struct {
	firstItem *linkedListItem
	size      int
}

func NewLinkedList() *linkedList {
	return &linkedList{}
}

func (ll *linkedList) Len() int {
	return ll.size
}

func (ll *linkedList) InsertByIndex(index, val int) {
	if ll.firstItem == nil {
		ll.firstItem = &linkedListItem{}
		ll.size++
	}

	currentItem := ll.firstItem

	for i := 1; i < index; i++ {
		if currentItem.next == nil {
			currentItem.next = &linkedListItem{}
			ll.size++
		}
		currentItem = currentItem.next
	}

	newItem := &linkedListItem{value: val}
	newItem.next = currentItem.next
	currentItem.next = newItem
	ll.size++
}

func (ll *linkedList) RemoveByIndex(index int) {
	if ll.size-1 < index {
		panic("index out of range")
	}

	if index == 0 && ll.size > 1 {
		ll.firstItem = ll.getElemByIndex(1)
	} else {
		elemBeforeIndex := ll.getElemByIndex(index - 1)
		indexElem := elemBeforeIndex.next
		elemAfterIndex := indexElem.next

		elemBeforeIndex.next = elemAfterIndex
	}

	ll.size--
}

func (ll *linkedList) AddToTheEnd(val int) {
	newItem := &linkedListItem{value: val}
	if ll.size > 0 {
		ll.getLastElem().next = newItem
	} else {
		ll.firstItem = newItem
	}

	ll.size++
}

func (ll *linkedList) RemoveLast() {
	ll.getElemByIndex(ll.size - 2).next = nil
	ll.size--
}

func (ll *linkedList) GetByIndex(index int) int {
	if ll.size-1 < index {
		panic("index out of range")
	}

	item := ll.firstItem
	for i := 1; i <= index; i++ {
		item = item.next
	}

	return item.Value()
}

func (ll *linkedList) getLastElem() *linkedListItem {
	return ll.getElemByIndex(ll.size - 1)
}

func (ll *linkedList) getElemByIndex(index int) *linkedListItem {
	if ll.size-1 < index {
		panic("index out of range")
	}

	item := ll.firstItem
	for i := 1; i <= index; i++ {
		item = item.next
	}

	return item
}

func (ll *linkedList) Print() {
	strBuilder := strings.Builder{}
	item := ll.firstItem

	strBuilder.WriteString(fmt.Sprintf("%d(%p)", item.Value(), item))
	for i := 1; i < ll.size; i++ {
		if item.next != nil {
			item = item.next
			strBuilder.WriteString(fmt.Sprintf("->%d(%p)", item.Value(), item))
		}
	}

	fmt.Println(strBuilder.String())
}
