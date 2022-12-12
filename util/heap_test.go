package util

import (
	"reflect"
	"testing"
)

type CustomInt int

func TestMaxHeapInt(t *testing.T) {
	heap := GetNewHeap(func(a, b CustomInt) bool { return a < b })
	values := []CustomInt{6, 5, 4, 8, 9, 10, 13, 12, 11, 7}

	// Push
	for _, v := range values {
		heap.Push(v)
	}
	expected := []CustomInt{13, 12, 10, 11, 7, 4, 9, 5, 8, 6}
	if !reflect.DeepEqual(heap.data, expected) {
		t.Errorf("Expected %v, got %v", expected, heap.data)
	}

	// Pop
	expected = []CustomInt{13, 12, 11, 10, 9, 8, 7, 6, 5, 4}
	for _, v := range expected {
		if val, ok := heap.Pop(); val != v && ok {
			t.Errorf("Expected %v, got %v", val, v)
		}
	}
}

func TestMinHeapInt(t *testing.T) {
	heap := GetNewHeap(func(a, b CustomInt) bool { return a > b })
	values := []CustomInt{6, 5, 4, 8, 9, 10, 13, 12, 11, 7}

	// Push
	for _, v := range values {
		heap.Push(v)
	}
	expected := []CustomInt{4, 6, 5, 8, 7, 10, 13, 12, 11, 9}
	if !reflect.DeepEqual(heap.data, expected) {
		t.Errorf("Expected %v, got %v", expected, heap.data)
	}

	// Pop
	expected = []CustomInt{4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	for _, v := range expected {
		if val, ok := heap.Pop(); val != v && ok {
			t.Errorf("Expected %v, got %v", val, v)
		}
	}
}
