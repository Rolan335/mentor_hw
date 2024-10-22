package queue

import "errors"

type Queue struct {
	queue []interface{}
}

func (q *Queue) Enqueue(elem interface{}) {
	q.queue = append(q.queue, elem)
}

func (q *Queue) Dequeue() interface{} {
	if len(q.queue) == 0 {
		return errors.New("error. Queue length is 0")
	}
	elem := q.queue[0]
	q.queue = q.queue[1:]
	return elem
}

func (q Queue) Peek() interface{} {
	if len(q.queue) == 0 {
		return errors.New("error. Queue length is 0")
	}
	return q.queue[0]
}

func (q Queue) Len() int{
	return len(q.queue)
}

