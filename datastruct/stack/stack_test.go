package stack

import (
	"reflect"
	"testing"
)

func TestStack_Push(t *testing.T) {
	stack := Stack{}
	wantElem, wantBool := "pushed 1 elem", true
	stack.Push(11)
	stack.Push(wantElem)
	res, ok := stack.Peek()
	if res != wantElem || ok != wantBool{
		t.Errorf("Stack.Peek() after Stack.Push() res, ok = %v, %v, want %v, %v", res, ok, wantElem, wantBool)
	}
}

func TestStack_Pop(t *testing.T) {
	tests := []struct {
		name     string
		s        *Stack
		want     interface{}
		wantBool bool
	}{
		{
			name:     "Pop len 0",
			s:        &Stack{},
			wantBool: false,
			want:     nil,
		},
		{
			name:     "Pop len >0",
			s:        &Stack{[]any{1, 2, 3, 4, 5}},
			want:     5,
			wantBool: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := tt.s.Pop()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.Pop() = %v, want %v", got, tt.want)
			}
			if ok != tt.wantBool {
				t.Errorf("Pack() error = %v, wantBool %v", ok, tt.wantBool)
				return
			}
		})
	}
}

func TestStack_Peek(t *testing.T) {
	tests := []struct {
		name     string
		s        Stack
		want     interface{}
		wantBool bool
	}{
		{
			name:     "Peek len 0",
			s:        Stack{},
			wantBool: false,
			want:     nil,
		},
		{
			name:     "Peek len >0",
			s:        Stack{[]any{1, 2, 3, 4, 5}},
			want:     5,
			wantBool: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := tt.s.Peek()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.Peek() = %v, want %v", got, tt.want)
			}
			if ok != tt.wantBool {
				t.Errorf("Pack() error = %v, wantBool %v", ok, tt.wantBool)
				return
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
