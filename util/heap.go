package util

type Heap[T any] struct {
	data        []T
	compareFunc func(a, b T) bool
}

// Returns a new max or min heap
//   - Max heap will be returned when compareFunc returns true in case of a < b
//   - Min heap will be returned when compareFunc returns true in case of a > b
func GetNewHeap[T any](compareFunc func(a, b T) bool) *Heap[T] {
	return &Heap[T]{
		data:        make([]T, 0),
		compareFunc: compareFunc,
	}
}

func (h *Heap[T]) Push(item T) {
	h.data = append(h.data, item)
	h.heapifyUp(h.Size() - 1)
}

func (h *Heap[T]) Pop() (item T, ok bool) {
	if h.Size() == 0 {
		return
	}

	item = h.data[0]
	h.swap(0, h.Size()-1)
	h.data = h.data[:h.Size()-1]
	h.heapifyDown(0)

	return item, true
}

func (h *Heap[T]) Peek() (item T, ok bool) {
	if h.Size() == 0 {
		return
	}

	return h.data[0], true
}

func (h *Heap[T]) Size() int {
	return len(h.data)
}

func (h *Heap[T]) heapifyUp(idx int) {
	for h.compareFunc(h.data[getParentIndex(idx)], h.data[idx]) {
		h.swap(idx, getParentIndex(idx))
		idx = getParentIndex(idx)
	}
}

func (h *Heap[T]) heapifyDown(idx int) {
	leftChildIndex, rightChildIndex := getLeftChildIndex(idx), getRightChildIndex(idx)
	parentIdx := idx

	if leftChildIndex < h.Size() && h.compareFunc(h.data[parentIdx], h.data[leftChildIndex]) {
		parentIdx = leftChildIndex
	}

	if rightChildIndex < h.Size() && h.compareFunc(h.data[parentIdx], h.data[rightChildIndex]) {
		parentIdx = rightChildIndex
	}

	if parentIdx != idx {
		h.swap(idx, parentIdx)
		h.heapifyDown(parentIdx)
	}
}

func (h *Heap[T]) swap(idx1, idx2 int) {
	h.data[idx1], h.data[idx2] = h.data[idx2], h.data[idx1]
}

func getLeftChildIndex(i int) int {
	return 2*i + 1
}

func getRightChildIndex(i int) int {
	return 2*i + 2
}

func getParentIndex(i int) int {
	return (i - 1) / 2
}
