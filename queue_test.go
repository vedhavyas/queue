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