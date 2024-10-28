package queue

import "errors"

type Queue struct {
	queue []interface{}
}

func (q *Queue) Enqueue(elem interface{}) {
	q.queue = append(q.queue, elem)
}

func (q *Queue) Dequeue() (interface{}, error) {
	if len(q.queue) == 0 {
		return nil, errors.New("error. Queue length is 0")
	}
	elem := q.queue[0]
	q.queue = q.queue[1:]
	return elem, nil
}

func (q Queue) Peek() (interface{}, error) {
	if len(q.queue) == 0 {
		return nil, errors.New("error. Queue length is 0")
	}
	return q.queue[0], nil
}

func (q Queue) Len() int {
	return len(q.queue)
}
