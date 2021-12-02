package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	first  *ListItem
	last   *ListItem
	length int
}

func (l list) Len() int {
	return l.length
}

func (l list) Front() *ListItem {
	return l.first
}

func (l list) Back() *ListItem {
	return l.last
}

func (l *list) PushFront(v interface{}) *ListItem {
	listItem := ListItem{Value: v}

	if l.first != nil {
		listItem.Next = l.first
		l.first.Prev = &listItem
	}

	if l.first == nil && l.last == nil {
		l.last = &listItem
	}

	l.first = &listItem
	l.length++

	return &listItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	listItem := ListItem{Value: v}

	if l.last != nil {
		listItem.Prev = l.last
		l.last.Next = &listItem
	}

	if l.first == nil && l.last == nil {
		l.first = &listItem
	}

	l.last = &listItem
	l.length++

	return &listItem
}

func (l *list) Remove(listItem *ListItem) {
	l.length--

	if l.first == listItem && l.last == listItem {
		l.first = nil
		l.last = nil
		return
	}

	if l.first == listItem {
		l.first = listItem.Next
	} else {
		listItem.Prev.Next = listItem.Next
	}
	if l.last == listItem {
		l.last = listItem.Prev
	} else {
		listItem.Next.Prev = listItem.Prev
	}

	listItem.Prev = nil
	listItem.Next = nil
}

func (l *list) MoveToFront(listItem *ListItem) {
	if l.first == listItem {
		return
	}
	if l.last == listItem {
		listItem.Prev.Next = nil
		l.last = listItem.Prev
	} else {
		listItem.Next.Prev = listItem.Prev
	}

	listItem.Prev.Next = listItem.Next
	listItem.Prev = nil
	listItem.Next = l.first
	l.first.Prev = listItem
	l.first = listItem
}

func NewList() List {
	return new(list)
}
