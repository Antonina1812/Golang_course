package algorithms

import "fmt"

type Queue struct {
	Queue []int
}

func (q *Queue) Enqueue(val int) {
	q.Queue = append(q.Queue, val)
}

func (q *Queue) Pop() {
	if !q.IsEmpty() {
		q.Queue = q.Queue[1:]
	}
}

func (q *Queue) IsEmpty() bool {
	return len(q.Queue) == 0
}

func (q *Queue) Peek() int {
	if !q.IsEmpty() {
		return q.Queue[0]
	}
	return -1
}

func QueuePrint() {
	var queue Queue

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	queue.Pop()
	first := queue.Peek()
	fmt.Println(queue, first)
}
