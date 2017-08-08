package queue

import (
	"testing"
)

func TestQueue_Enqueue(t *testing.T) {
	q := &Queue{}

	tests := []struct {
		value int
		len   int
	}{
		{value: 01, len: 1},
		{value: 21, len: 2},
		{value: 31, len: 3},
		{value: 43, len: 4},
		{value: 50, len: 5},
		{value: 12, len: 6},
	}

	for _, c := range tests {
		q.Enqueue(c.value)
		if c.len != q.Len() {
			t.Fatalf("Expected queue length %d but got %d\n", c.len, q.Len())
		}
	}
}

func TestQueue_Dequeue(t *testing.T) {
	q := &Queue{}
	for i := 0; i < 5; i++ {
		q.Enqueue(i)
	}

	tests := []struct {
		value int
		next  int
		len   int
		err   bool
	}{
		{value: 0, next: 0, len: 4},
		{value: 1, next: 0, len: 3},
		{value: 2, next: 0, len: 2},
		{value: 3, next: 0, len: 1},
		{value: 4, next: 10, len: 0},
		{value: 10, next: 100, len: 0},
		{value: 100, next: 0, len: 0},
		{err: true},
	}

	for _, c := range tests {
		v, err := q.Dequeue()
		if err != nil {
			if c.err {
				continue
			}

			t.Fatalf("Expected %d but got an error: %v\n", c.value, err)
		}

		if c.value != v.(int) {
			t.Fatalf("Expected %d but got %v\n", c.value, v)
		}

		if c.len != q.Len() {
			t.Fatalf("Expected queue length %d but got %d\n", c.len, q.Len())
		}

		if c.next != 0 {
			q.Enqueue(c.next)
		}
	}
}

func TestQueue_Get(t *testing.T) {
	q := &Queue{}
	for i := 0; i < 5; i++ {
		q.Enqueue(i)
	}

	tests := []struct {
		i     int
		value int
		len   int
		err   bool
	}{
		{i: 1, value: 1, len: 4},
		{i: 5, err: true},
		{i: 3, value: 4, len: 3},
		{i: 0, value: 0, len: 2},
		{i: 1, value: 3, len: 1},
	}

	for _, c := range tests {
		v, err := q.Get(c.i)
		if err != nil {
			if c.err {
				continue
			}

			t.Fatalf("Expected %d but got an error: %v\n", c.value, err)
		}

		if c.value != v.(int) {
			t.Fatalf("Expected %d but got %v\n", c.value, v)
		}

		if c.len != q.Len() {
			t.Fatalf("Expected queue length %d but got %d\n", c.len, q.Len())
		}

	}
}

func TestQueue_Peak(t *testing.T) {
	q := &Queue{}
	for i := 0; i < 5; i++ {
		q.Enqueue(i)
	}

	tests := []struct {
		value int
		len   int
		err   bool
	}{
		{value: 0, len: 5},
		{value: 1, len: 4},
		{value: 2, len: 3},
		{value: 3, len: 2},
		{value: 4, len: 1},
		{err: true},
	}

	for _, c := range tests {
		v, err := q.Peak()
		if err != nil {
			if c.err {
				continue
			}

			t.Fatalf("Expected %d but got an error: %v\n", c.value, err)
		}

		if c.value != v.(int) {
			t.Fatalf("Expected %d but got %v\n", c.value, v)
		}

		if c.len != q.Len() {
			t.Fatalf("Expected queue length %d but got %d\n", c.len, q.Len())
		}

		q.Dequeue()
	}
}

func TestQueue_PeakAt(t *testing.T) {
	q := &Queue{}
	for i := 0; i < 5; i++ {
		q.Enqueue(i)
	}

	tests := []struct {
		i     int
		value int
		len   int
		err   bool
	}{
		{i: 1, value: 1, len: 5},
		{i: 5, err: true},
		{i: 3, value: 3, len: 5},
		{i: 0, value: 0, len: 5},
		{i: 2, value: 2, len: 5},
	}

	for _, c := range tests {
		v, err := q.PeakAt(c.i)
		if err != nil {
			if c.err {
				continue
			}

			t.Fatalf("Expected %d but got an error: %v\n", c.value, err)
		}

		if c.value != v.(int) {
			t.Fatalf("Expected %d but got %v\n", c.value, v)
		}

		if c.len != q.Len() {
			t.Fatalf("Expected queue length %d but got %d\n", c.len, q.Len())
		}

	}
}

func TestQueue_Next(t *testing.T) {
	q := &Queue{}
	for i := 0; i < 5; i++ {
		q.Enqueue(i)
	}

	tests := []struct {
		value int
		len   int
		err   bool
	}{
		{value: 0, len: 4},
		{value: 1, len: 3},
		{value: 2, len: 2},
		{value: 3, len: 1},
		{value: 4, len: 0},
		{err: true},
	}

	q.ResetRange()
	for _, c := range tests {
		v := q.Next()
		if v == nil {
			if c.err {
				continue
			}

			t.Fatalf("Expected value %v but got nil\n", c.value)
		}

		if c.value != v.(int) {
			t.Fatalf("Expected %d but got %v\n", c.value, v)
		}
	}
}

func TestQueue_CutRangeItem_one(t *testing.T) {
	q := &Queue{}
	q.Enqueue(1)

	q.ResetRange()
	q.Next()
	q.CutRangeItem()

	if q.Len() != 0 {
		t.Fatalf("Expected 0 but got %d\n", q.Len())
	}

	v, err := q.Dequeue()
	if err == nil {
		t.Fatalf("Expected error but got: %v", v)
	}
}

func TestQueue_CutRangeItem_two_start(t *testing.T) {
	q := &Queue{}
	q.Enqueue(0)
	q.Enqueue(1)

	q.ResetRange()
	q.Next()
	q.CutRangeItem()

	if q.Len() != 1 {
		t.Fatalf("Expected 1 but got %d\n", q.Len())
	}

	v, err := q.Dequeue()
	if err != nil {
		t.Fatalf("Expected value but got: %v\n", err)
	}

	if v.(int) != 1 {
		t.Fatalf("Expected %d but got %v\n", 1, v)
	}
}

func TestQueue_CutRangeItem_two_end(t *testing.T) {
	q := &Queue{}
	q.Enqueue(0)
	q.Enqueue(1)

	q.ResetRange()
	q.Next()
	q.Next()
	q.CutRangeItem()

	if q.Len() != 1 {
		t.Fatalf("Expected 1 but got %d\n", q.Len())
	}

	v, err := q.Dequeue()
	if err != nil {
		t.Fatalf("Expected value but got: %v\n", err)
	}

	if v.(int) != 0 {
		t.Fatalf("Expected %d but got %v\n", 0, v)
	}
}

func TestQueue_CutRangeItem_middle(t *testing.T) {
	q := &Queue{}
	q.Enqueue(0)
	q.Enqueue(1)
	q.Enqueue(2)

	q.ResetRange()
	q.Next()
	q.Next()
	q.CutRangeItem()

	if q.Len() != 2 {
		t.Fatalf("Expected 1 but got %d\n", q.Len())
	}

	v, err := q.Dequeue()
	if err != nil {
		t.Fatalf("Expected value but got: %v\n", err)
	}

	if v.(int) != 0 {
		t.Fatalf("Expected %d but got %v\n", 0, v)
	}

	v, err = q.Dequeue()
	if err != nil {
		t.Fatalf("Expected value but got: %v\n", err)
	}

	if v.(int) != 2 {
		t.Fatalf("Expected %d but got %v\n", 0, v)
	}
}

func TestQueue_CutRangeItem(t *testing.T) {
	q := &Queue{}
	for i := 0; i < 5; i++ {
		q.Enqueue(i)
	}

	tests := []struct {
		dequeue bool
		value   int
		length  int
		error   bool
	}{
		{dequeue: true, value: 0, length: 4},
		{dequeue: false, value: 1, length: 4},
		{dequeue: true, value: 2, length: 3},
		{dequeue: true, value: 3, length: 2},
		{dequeue: false, value: 4, length: 2},
		{error: true},
	}

	q.ResetRange()
	for _, c := range tests {
		v := q.Next()
		if v == nil {
			if c.error {
				continue
			}

			t.Fatalf("Expected %d but got nil\n", c.value)
		}

		if c.value != v.(int) {
			t.Fatalf("Expected %d but got %v\n", c.value, v)
		}

		if c.dequeue {
			q.CutRangeItem()
		}

		if c.length != q.Len() {
			t.Fatalf("Expected length %d but got %v\n", c.length, q.Len())
		}
	}
}
