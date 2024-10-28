package queue

import (
	"errors"
	"reflect"
	"testing"
)

// How to test void functions
func TestQueue_Enqueue(t *testing.T) {
	type args struct {
		elem interface{}
	}
	tests := []struct {
		name string
		q    *Queue
		args args
	}{
		{
			name: "Enqueue string",
			q:    &Queue{},
			args: args{elem: "hello"},
		},
		{
			name: "Enqueue int",
			q:    &Queue{},
			args: args{elem: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.q.Enqueue(tt.args.elem)
		})
	}
}

func TestQueue_Dequeue(t *testing.T) {
	tests := []struct {
		name    string
		q       *Queue
		want    interface{}
		wantErr error
	}{
		{
			name: "Dequeue when nil",
			q:    &Queue{},
			wantErr: errors.New("error. Queue length is 0"),
		},
		{
			name: "Dequeue when not nil",
			q:    &Queue{queue: []interface{}{1, "hello"}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.q.Dequeue()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queue.Dequeue() = %v, want %v", got, tt.want)
			}
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("Queue.Dequeue() = %v, want %v", got, tt.wantErr)
			}
		})
	}
}

func TestQueue_Peek(t *testing.T) {
	tests := []struct {
		name string
		q    Queue
		want interface{}
		wantErr error
	}{
		{
			name: "Peek when nil",
			q:    Queue{},
			wantErr: errors.New("error. Queue length is 0"),
		},
		{
			name: "Peek when not nil",
			q:    Queue{[]interface{}{222, 5.542, "hello"}},
			want: 222,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.q.Peek()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queue.Peek() = %v, want %v", got, tt.want)
			}
			if err != nil && err.Error() != tt.wantErr.Error(){
				t.Errorf("Queue.Peek() = %v, want %v", got, tt.wantErr)
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
