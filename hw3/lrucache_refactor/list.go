package lrucache_refactor

type list struct {
	len  int
	head *node
}

type node struct {
	data string
	next *node
}

func newList() *list {
	return &list{head: nil}
}

func (l *list) addToStart(d string) {
	l.len++
	if l.head == nil {
		l.head = &node{data: d, next: nil}
		return
	}
	l.head = &node{data: d, next: l.head}
}

func (l *list) moveToStart(d string) {
	if l.head == nil || l.head.data == d {
		return
	}
	var prev *node
	current := l.head
	for current != nil && current.data != d {
		prev = current
		current = current.next
	}
	if current == nil {
		return
	}
	if prev != nil {
		prev.next = current.next
	}
	current.next = l.head
	l.head = current
}

func (l *list) getEndVal() string {
	if l.head == nil {
		return ""
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	return current.data
}

func (l *list) remove(d string) {
	if l.head == nil {
		return
	}
	current := l.head
	if l.head.data == d {
		l.head = l.head.next
		l.len--
		return
	}
	for current.next != nil {
		if current.next.data == d {
			current.next = current.next.next
			l.len--
			return
		}
		current = current.next
	}
}
