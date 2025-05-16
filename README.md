# Priority Queue for Go

This package implements a priority queue on top of the container/heap package.

In addition to the standard Push and Pop functionality, it also provides functions to update the priority of an item and check if an item exists in the queue.

The main computation is done using the heap interface.
Because a map of all items is maintained, the priority lookup is O(1) which is used for priority updates which is O(log n)
(Source: [jupp0r/go-priority-queue](https://github.com/jupp0r/go-priority-queue/blob/master/priorty_queue.go#L5-L6))

This package was inspired by [jupp0r/go-priority-queue](https://github.com/jupp0r/go-priority-queue)
and the implementation example in the [container/heap](https://pkg.go.dev/container/heap) package.
But both of them do not support generic types.
