// This package implements a priority queue on top of the container/heap package.
//
// In addition to the standard Push and Pop functionality, it also provides functions to update the priority of an item and check if an item exists in the queue.
//
// The main computation is done using the heap interface.
// Because a map of all items is maintained, the priority lookup is O(1) which is used for priority updates which is O(log n)
// (Source: [jupp0r/go-priority-queue](https://github.com/jupp0r/go-priority-queue/blob/master/priorty_queue.go#L5-L6))
//
// This package was inspired by [jupp0r/go-priority-queue](https://github.com/jupp0r/go-priority-queue)
// and the implementation example in the [container/heap](https://pkg.go.dev/container/heap) package.
// But both of them do not support generic types.

package pq

import (
	"container/heap"
)

type PriorityQueue[T comparable] struct {
	itemHeap *itemHeap[T]
	lookup   map[T]*item[T]
}

func New[T comparable]() PriorityQueue[T] {
	return PriorityQueue[T]{
		itemHeap: &itemHeap[T]{},
		lookup:   make(map[T]*item[T]),
	}
}

func (p *PriorityQueue[T]) Len() int {
	return p.itemHeap.Len()
}

func (p *PriorityQueue[T]) Push(value T, priority float64) error {
	_, ok := p.lookup[value]
	if ok {
		return ErrItemExists
	}

	newItem := &item[T]{
		value:    value,
		priority: priority,
	}
	heap.Push(p.itemHeap, newItem)
	p.lookup[value] = newItem

	return nil
}

func (p *PriorityQueue[T]) Pop() (T, error) {
	if p.itemHeap.Len() == 0 {
		var zero T
		return zero, ErrQueueEmpty
	}

	item := heap.Pop(p.itemHeap).(*item[T])
	delete(p.lookup, item.value)

	return item.value, nil
}

func (p *PriorityQueue[T]) Update(value T, newPriority float64) error {
	item, ok := p.lookup[value]
	if !ok {
		return ErrItemNotFound
	}

	item.priority = newPriority
	heap.Fix(p.itemHeap, item.index)
	return nil
}

func (p *PriorityQueue[T]) Contains(value T) bool {
	_, ok := p.lookup[value]
	return ok
}

type item[T any] struct {
	value    T
	priority float64
	index    int
}

type itemHeap[T any] []*item[T]

func (i itemHeap[T]) Len() int {
	return len(i)
}

func (i itemHeap[T]) Less(x, y int) bool {
	return i[x].priority < i[y].priority
}

func (i itemHeap[T]) Swap(x, y int) {
	i[x], i[y] = i[y], i[x]
	i[x].index = x
	i[y].index = y
}

func (i *itemHeap[T]) Push(x any) {
	item := x.(*item[T])
	item.index = len(*i)
	*i = append(*i, item)
}

func (i *itemHeap[T]) Pop() any {
	old := *i
	item := old[len(old)-1]
	*i = old[0 : len(old)-1]
	return item
}
