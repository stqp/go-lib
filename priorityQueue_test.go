package lib

import (
	"testing"
)

func TestPush(t *testing.T) {
	pq := PriorityQueue{
		heapSize: 100,
	}
	in := []int{
		5, 9, 1, 10, 7, 6, 8,
	}
	expected := []int{
		1, 5, 6, 7, 8, 9, 10,
	}
	for _, v := range in {
		pq.Push(v)
	}
	for _, v := range expected {
		actual := pq.Pop()
		if actual != v {
			t.Error()
		}
	}
}

func TestPushWithReverseOption(t *testing.T) {
	pq := PriorityQueue{
		heapSize: 100,
		reverse:  true,
	}
	in := []int{
		5, 9, 1, 10, 7, 6, 8,
	}
	expected := []int{
		10, 9, 8, 7, 6, 5, 1,
	}
	for _, v := range in {
		pq.Push(v)
	}
	for _, v := range expected {
		actual := pq.Pop()
		if actual != v {
			t.Error()
		}
	}
}
