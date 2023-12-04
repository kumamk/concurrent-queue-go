package queue

import "sync"

// thread safe queue via mutex
type ConcurrentQueue struct {
	list []int
	mu   sync.Mutex
}

func NewConcurrentQueue() ConcurrentQueue {
	return ConcurrentQueue{list: make([]int, 0)}
}

func (cq *ConcurrentQueue) Add(x int) {
	cq.mu.Lock()
	defer cq.mu.Unlock()

	cq.list = append(cq.list, x)
}

func (cq *ConcurrentQueue) Remove() int {
	cq.mu.Lock()
	defer cq.mu.Unlock()

	if len(cq.list) == 0 {
		return -1
	}

	val := cq.list[0]
	// remove 1st element
	cq.list = cq.list[1:]
	return val
}

func (cq *ConcurrentQueue) Len() int {
	cq.mu.Lock()
	defer cq.mu.Unlock()

	return len(cq.list)
}
