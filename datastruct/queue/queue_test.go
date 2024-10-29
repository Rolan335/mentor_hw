package queue

import (
	"reflect"
	"testing"
)

// Make normal test for append
func TestQueue_Enqueue(t *testing.T) {
	queue := Queue{}
	wantElem, wantBool := "enqueued 1 elem", true
	queue.Enqueue(wantElem)
	queue.Enqueue(11)
	res, ok := queue.Peek()
	if res != wantElem || ok != wantBool {
		t.Errorf("Queue.Peek() after Queue.Enqueue() res, ok = %v, %v, want %v, %v", res, ok, wantElem, wantBool)
	}
}

func TestQueue_Dequeue(t *testing.T) {
	tests := []struct {
		name     string
		q        *Queue
		want     interface{}
		wantBool bool
	}{
		{
			name:     "Dequeue when nil",
			q:        &Queue{},
			wantBool: false,
			want:     nil,
		},
		{
			name:     "Dequeue when not nil",
			q:        &Queue{queue: []interface{}{1, "hello"}},
			want:     1,
			wantBool: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := tt.q.Dequeue()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queue.Dequeue() = %v, want %v", got, tt.want)
			}
			if tt.wantBool != ok {
				t.Errorf("Queue.Dequeue() = %v, want %v", got, tt.wantBool)
			}
		})
	}
}

func TestQueue_Peek(t *testing.T) {
	tests := []struct {
		name     string
		q        Queue
		want     interface{}
		wantbool bool
	}{
		{
			name:     "Peek when nil",
			q:        Queue{},
			wantbool: false,
			want:     nil,
		},
		{
			name:     "Peek when not nil",
			q:        Queue{[]interface{}{222, 5.542, "hello"}},
			want:     222,
			wantbool: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := tt.q.Peek()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queue.Peek() = %v, want %v", got, tt.want)
			}
			if tt.wantbool != ok {
				t.Errorf("Queue.Peek() = %v, want %v", got, tt.wantbool)
			}
		})
	}
}

func TestQueue_Len(t *testing.T) {
	tests := []struct {
		name string
		q    Queue
		want int
	}{
		{
			name: "Len",
			q:    Queue{[]interface{}{223, 5, 35, 1}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.Len(); got != tt.want {
				t.Errorf("Queue.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}
