package pq

type PriorityQueueError string

func (e PriorityQueueError) Error() string {
	return string(e)
}

const (
	ErrQueueEmpty   = PriorityQueueError("queue is empty")
	ErrItemNotFound = PriorityQueueError("item not found in the queue")
	ErrItemExists   = PriorityQueueError("item already exists in the queue")
)
