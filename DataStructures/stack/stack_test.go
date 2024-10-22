package stack

import (
	"errors"
	"reflect"
	"testing"
)

func TestStack_Push(t *testing.T) {
	type args struct {
		elem interface{}
	}
	tests := []struct {
		name string
		s    *Stack
		args args
	}{
		{
			name: "Push elem to stack",
			s:    &Stack{},
			args: args{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Push(tt.args.elem)
			if got := tt.s.stack[0].(int); got != tt.args.elem {
				t.Errorf("Stack.Pop() = %v, want %v", got, tt.args.elem)
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	tests := []struct {
		name string
		s    *Stack
		want interface{}
	}{
		{
			name: "Pop len 0",
			s:    &Stack{},
			want: errors.New("error. Stack len is 0"),
		},
		{
			name: "Pop len >0",
			s:    &Stack{[]any{1, 2, 3, 4, 5}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Peek(t *testing.T) {
	tests := []struct {
		name string
		s    Stack
		want interface{}
	}{
		{
			name: "Peek len 0",
			s:    Stack{},
			want: errors.New("error. Stack len is 0"),
		},
		{
			name: "Peek len >0",
			s:    Stack{[]any{1, 2, 3, 4, 5}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Peek(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Len(t *testing.T) {
	tests := []struct {
		name string
		s    Stack
		want int
	}{
		{
			name: "Len",
			s:    Stack{[]any{1, 2, 3, 4, 5, 6, 7}},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Len(); got != tt.want {
				t.Errorf("Stack.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}
