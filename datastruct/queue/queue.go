package queue

type Queue struct {
	queue []interface{}
}

func (q *Queue) Enqueue(elem interface{}) {
	q.queue = append(q.queue, elem)
}

func (q *Queue) Dequeue() (interface{}, bool) {
	if len(q.queue) == 0 {
		return nil, false
	}
	elem := q.queue[0]
	q.queue = q.queue[1:]
	return elem, true
}

func (q *Queue) Peek() (interface{}, bool) {
	if len(q.queue) == 0 {
		return nil, false
	}
	return q.queue[0], true
}

func (q *Queue) Len() int {
	return len(q.queue)
}
