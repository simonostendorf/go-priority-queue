package pq

import (
	"testing"

	"github.com/simonostendorf/go-priority-queue/internal/helpers/ptr"
)

func TestLen(t *testing.T) {
	tests := []struct {
		name     string
		queue    PriorityQueue[int]
		expected int
	}{
		{
			name:     "empty queue",
			queue:    New[int](),
			expected: 0,
		},
		{
			name: "non-empty queue",
			queue: func() PriorityQueue[int] {
				q := New[int]()
				q.Push(1, 1.0) // nolint:errcheck
				q.Push(2, 2.0) // nolint:errcheck
				return q
			}(),
			expected: 2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.queue.Len() != test.expected {
				t.Errorf("expected %d, got %d", test.expected, test.queue.Len())
			}
		})
	}
}

func TestPush(t *testing.T) {
	tests := []struct {
		name             string
		queue            PriorityQueue[int]
		value            int
		priority         float64
		expectedLength   int
		expectedPosition *int
		expectedError    error
	}{
		{
			name:             "push to empty queue",
			queue:            New[int](),
			value:            1,
			priority:         1.0,
			expectedLength:   1,
			expectedPosition: ptr.Ptr(0),
			expectedError:    nil,
		},
		{
			name: "push to non-empty queue",
			queue: func() PriorityQueue[int] {
				q := New[int]()
				q.Push(1, 1.0) // nolint:errcheck
				return q
			}(),
			value:            2,
			priority:         2.0,
			expectedLength:   2,
			expectedPosition: ptr.Ptr(1),
			expectedError:    nil,
		},
		{
			name: "push existing item",
			queue: func() PriorityQueue[int] {
				q := New[int]()
				q.Push(1, 1.0) // nolint:errcheck
				return q
			}(),
			value:            1,
			priority:         2.0,
			expectedLength:   1,
			expectedPosition: nil,
			expectedError:    ErrItemExists,
		},
		{
			name: "push with zero priority",
			queue: func() PriorityQueue[int] {
				q := New[int]()
				q.Push(1, 1.0) // nolint:errcheck
				return q
			}(),
			value:            3,
			priority:         0.0,
			expectedLength:   2,
			expectedPosition: ptr.Ptr(0),
			expectedError:    nil,
		},
		{
			name: "push with negative priority",
			queue: func() PriorityQueue[int] {
				q := New[int]()
				q.Push(1, 1.0) // nolint:errcheck
				return q
			}(),
			value:            4,
			priority:         -1.0,
			expectedLength:   2,
			expectedPosition: ptr.Ptr(0),
			expectedError:    nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.queue.Push(test.value, test.priority)
			if err != test.expectedError {
				t.Errorf("expected error %v, got %v", test.expectedError, err)
			}
			if test.queue.Len() != test.expectedLength {
				t.Errorf("expected length %d, got %d", test.expectedLength, test.queue.Len())
			}
			if test.expectedPosition != nil {
				for i := range test.queue.Len() {
					currentItem, err := test.queue.Pop()
					if err != nil {
						t.Errorf("unexpected error: %v", err)
					}
					if i == *test.expectedPosition {
						if currentItem != test.value {
							t.Errorf("expected value %d at position %d, got %d", test.value, *test.expectedPosition, currentItem)
						} else {
							break
						}
					} else {
						if currentItem == test.value {
							t.Errorf("expected value %d not to be at position %d", test.value, *test.expectedPosition)
						}
					}
				}
			}
		})
	}
}

func TestPop(t *testing.T) {
	tests := []struct {
		name          string
		queue         PriorityQueue[int]
		expected      int
		expectedError error
	}{
		{
			name:          "pop from empty queue",
			queue:         New[int](),
			expected:      0,
			expectedError: ErrQueueEmpty,
		},
		{
			name: "pop from non-empty queue",
			queue: func() PriorityQueue[int] {
				q := New[int]()
				q.Push(1, 1.0) // nolint:errcheck
				return q
			}(),
			expected:      1,
			expectedError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			value, err := test.queue.Pop()
			if err != nil && test.expected != 0 {
				t.Errorf("unexpected error: %v", err)
			}
			if value != test.expected {
				t.Errorf("expected %d, got %d", test.expected, value)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	tests := []struct {
		name             string
		queue            PriorityQueue[int]
		value            int
		newPriority      float64
		expectedPosition *int
		expectedError    error
	}{
		{
			name:             "update non-existing item",
			queue:            New[int](),
			value:            1,
			newPriority:      2.0,
			expectedPosition: nil,
			expectedError:    ErrItemNotFound,
		},
		{
			name: "update existing item",
			queue: func() PriorityQueue[int] {
				q := New[int]()
				q.Push(1, 1.0) // nolint:errcheck
				return q
			}(),
			value:            1,
			newPriority:      2.0,
			expectedPosition: ptr.Ptr(0),
			expectedError:    nil,
		},
		{
			name: "update with multiple items",
			queue: func() PriorityQueue[int] {
				q := New[int]()
				q.Push(1, 1.0) // nolint:errcheck
				q.Push(2, 2.0) // nolint:errcheck
				q.Push(3, 3.0) // nolint:errcheck
				return q
			}(),
			value:            2,
			newPriority:      0.5,
			expectedPosition: ptr.Ptr(0),
			expectedError:    nil,
		},
		{
			name: "update with zero priority",
			queue: func() PriorityQueue[int] {
				q := New[int]()
				q.Push(1, 1.0) // nolint:errcheck
				return q
			}(),
			value:            1,
			newPriority:      0.0,
			expectedPosition: ptr.Ptr(0),
			expectedError:    nil,
		},
		{
			name: "update with negative priority",
			queue: func() PriorityQueue[int] {
				q := New[int]()
				q.Push(1, 1.0) // nolint:errcheck
				return q
			}(),
			value:            1,
			newPriority:      -1.0,
			expectedPosition: ptr.Ptr(0),
			expectedError:    nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.queue.Update(test.value, test.newPriority)
			if err != test.expectedError {
				t.Errorf("expected error %v, got %v", test.expectedError, err)
			}
			if test.expectedPosition != nil {
				for i := range test.queue.Len() {
					currentItem, err := test.queue.Pop()
					if err != nil {
						t.Errorf("unexpected error: %v", err)
					}
					if i == *test.expectedPosition {
						if currentItem != test.value {
							t.Errorf("expected value %d at position %d, got %d", test.value, *test.expectedPosition, currentItem)
						} else {
							break
						}
					} else {
						if currentItem == test.value {
							t.Errorf("expected value %d not to be at position %d", test.value, *test.expectedPosition)
						}
					}
				}
			}
		})
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		name          string
		queue         PriorityQueue[int]
		value         int
		expected      bool
		expectedError error
	}{
		{
			name:          "contains in empty queue",
			queue:         New[int](),
			value:         1,
			expected:      false,
			expectedError: nil,
		},
		{
			name: "contains in non-empty queue",
			queue: func() PriorityQueue[int] {
				q := New[int]()
				q.Push(1, 1.0) // nolint:errcheck
				return q
			}(),
			value:         1,
			expected:      true,
			expectedError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.queue.Contains(test.value)
			if result != test.expected {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}
