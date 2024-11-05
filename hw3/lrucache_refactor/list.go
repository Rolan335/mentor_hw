package lrucache_refactor

type List struct {
	Len int
	head *Node
}

type Node struct {
	data string
	next *Node
}

func NewList() *List {
	return &List{head: nil}
}

func (l *List) AddToStart(d string) {
	l.Len++
	if l.head == nil {
		l.head = &Node{data: d, next: nil}
		return
	}
	l.head = &Node{data: d, next: l.head}
}

func (l *List) GetEndVal() string {
	if l.head == nil {
		return ""
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	return current.data
}

func (l *List) Remove(d string) {
	if l.head == nil {
		return
	}
	current := l.head
	if l.head.data == d {
		l.head = l.head.next
		l.Len--
		return
	}
	for current.next != nil {
		if current.next.data == d {
			current.next = current.next.next
			l.Len--
			return
		}
		current = current.next
	}

}

func (l *List) ShowAsSlice() []string {
	slice := make([]string, 0)
	current := l.head
	for current != nil {
		slice = append(slice, current.data)
		current = current.next
	}
	return slice
}

// func (l *List) AddToEnd(d string) {
// 	newNode := &Node{data: d, next: nil}
// 	if l.head == nil {
// 		l.head = newNode
// 		return
// 	}
// 	current := l.head
// 	for current.next != nil {
// 		current = current.next
// 	}
// 	current.next = newNode
// }
