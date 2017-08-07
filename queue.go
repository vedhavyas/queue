/*
package queue implements the Queue FIFO with O(1) for
Enqueue and Dequeue operations.
*/

package queue

import (
	"bytes"
	"fmt"
	"sync"
)

var iobError = fmt.Errorf("Queue: Index out of bounds")

type (
	// node is a single item in the queue
	// holds next, previous item and the its value
	node struct {
		prev  *node
		next  *node
		value interface{}
	}

	// Queue represents the FIFO list of items
	Queue struct {
		count   int
		first   *node
		last    *node
		current *node
		mu      sync.RWMutex // protects the items above
	}
)

// Len returns the items count in the queue
func (q *Queue) Len() int {
	q.mu.RLock()
	defer q.mu.RUnlock()

	return q.count
}

func (q *Queue) enqueue(item interface{}) {
	n := &node{value: item}
	q.count++
	if q.first == nil {
		q.first = n
		q.last = n
		return
	}

	q.last.next = n
	n.prev = q.last
	q.last = n
}

// Enqueue add the item to the queue: O(1)
func (q *Queue) Enqueue(item interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.enqueue(item)
}

func (q *Queue) dequeue() (interface{}, error) {
	if q.count == 0 {
		return nil, iobError
	}

	q.count--
	n := q.first
	nn := n.next
	if nn != nil {
		nn.prev = n.prev
	} else {
		q.last = nil
	}

	q.first = nn
	return n.value, nil
}

// Dequeue returns the first item from queue: O(1).
// if queue is empty, returns an error
func (q *Queue) Dequeue() (interface{}, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	return q.dequeue()
}

// Get returns the item at the given index from the list: O(n)
// returns an error if queue empty/out of bounds
func (q *Queue) Get(i int) (interface{}, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if i >= q.count {
		return nil, iobError
	}

	if i == 0 {
		return q.dequeue()
	}

	var n *node
	for x := 0; x <= i; x++ {
		if n == nil {
			n = q.first
			continue
		}

		n = n.next
	}

	q.count--
	p := n.prev
	nn := n.next
	p.next = nn
	if nn != nil {
		nn.prev = p
	}

	return n.value, nil
}

func (q *Queue) peakAt(i int) (interface{}, error) {
	if i >= q.count {
		return nil, iobError
	}
	var n *node
	for x := 0; x <= i; x++ {
		if n == nil {
			n = q.first
			continue
		}

		n = n.next
	}

	return n.value, nil
}

// Peak returns the next value in queue but does not remove from queue: O(1)
// returns error if queue empty
func (q *Queue) Peak() (interface{}, error) {
	q.mu.RLock()
	defer q.mu.RUnlock()

	return q.peakAt(0)
}

// PeakAt returns the item at the index i from the queue: O(n)
// returns error if queue empty/out of bounds
func (q *Queue) PeakAt(i int) (interface{}, error) {
	q.mu.RLock()
	defer q.mu.RUnlock()

	return q.peakAt(i)
}

// String dumps the queue in human readable format: O(n)
func (q *Queue) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("[")
	n := q.first
	for n != nil {
		if n.prev != nil {
			buffer.WriteString(" ")
		}
		buffer.WriteString(fmt.Sprint(n.value))
		n = n.next
	}
	buffer.WriteString("]")
	return buffer.String()
}

// ResetRange will set the seek to start: O(1)
func (q *Queue) ResetRange() {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.current = nil
}

// Next will return the next item in queue
// Returns nil if the range reaches the end of the queue: O(n)
func (q *Queue) Next() interface{} {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.count == 0 {
		return nil
	}

	if q.current == nil {
		q.current = q.first
	} else {
		q.current = q.current.next
	}

	if q.current == nil {
		return nil
	}

	return q.current.value
}

// CutRangeItem removes the current item in the range: O(1)
func (q *Queue) CutRangeItem() {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.count == 0 || q.current == nil {
		return
	}

	defer func() { q.count-- }()

	if q.count == 1 {
		q.current = nil
		q.first = nil
		q.last = nil
		return
	}

	p := q.current.prev
	nn := q.current.next

	if nn != nil {
		nn.prev = p
		if p == nil {
			q.first = nn
		}
	}

	if p != nil {
		p.next = nn
		if nn == nil {
			q.last = p
		}
	}

	q.current = q.current.prev
}
