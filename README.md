# queue
--
    import "github.com/vedhavyas/queue"


## Usage

#### type Queue

```go
type Queue struct {}
```

Queue represents the FIFO list of items

#### func (*Queue) Dequeue

```go
func (q *Queue) Dequeue() (interface{}, error)
```
Dequeue returns the first item from queue: O(1)

Returns an error if queue empty/out of bounds

#### func (*Queue) Enqueue

```go
func (q *Queue) Enqueue(item interface{})
```
Enqueue add the item to the queue: O(1)

#### func (*Queue) Get

```go
func (q *Queue) Get(i int) (interface{}, error)
```
Get returns the item at the given index from the list: O(n)

Returns an error if queue empty/out of bounds

#### func (*Queue) Len

```go
func (q *Queue) Len() int
```
Len returns the items count in the queue

#### func (*Queue) Peak

```go
func (q *Queue) Peak() (interface{}, error)
```
Peak returns the next value in queue but does not remove from queue: O(1)

Returns error if queue empty

#### func (*Queue) PeakAt

```go
func (q *Queue) PeakAt(i int) (interface{}, error)
```
PeakAt returns the item at the index i from the queue: O(n)

Returns error if queue empty/out of bounds

#### func (*Queue) String

```go
func (q *Queue) String() string
```
String dumps the queue in human readable format: O(n)
