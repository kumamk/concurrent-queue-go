package main

import (
	"concurrent-queue-go/queue"
	"fmt"
	"sync"
)

const (
	LIMIT = 100
)

func main() {
	var NQueueWg sync.WaitGroup
	var CQueueWg1 sync.WaitGroup
	var CQueueWg2 sync.WaitGroup

	fmt.Println("------Normal queue sequential flow-----")
	fmt.Println()

	nqueue := queue.NewNormalQueue()

	// Add 3 items
	nqueue.Add(10)
	nqueue.Add(20)
	nqueue.Add(30)
	fmt.Println("Add 3 items to normal queue")

	// remove all items
	nqueue.Remove()
	nqueue.Remove()
	nqueue.Remove()
	fmt.Println("Remove 3 items from normal queue")

	fmt.Printf("Normal queue size %d", nqueue.Len())
	fmt.Println()
	fmt.Println()

	fmt.Println("------Concurrent write to Normal queue-----")
	fmt.Println()

	fmt.Println("Adding 100 items conncurrently to normal queue")
	for i := 0; i < LIMIT; i++ {
		NQueueWg.Add(1)
		go func(i int) {
			nqueue.Add(i * 2)
			NQueueWg.Done()
		}(i)
	}
	fmt.Printf("Normal queue size %d, ideally should be %d", nqueue.Len(), LIMIT)
	fmt.Println()
	fmt.Println()

	fmt.Println("-----Concurrent queue flow-----")
	fmt.Println()

	// thread safe queue
	cque := queue.NewConcurrentQueue()
	fmt.Println("Adding 100 items to concurrent queue")
	for i := 0; i < LIMIT; i++ {
		CQueueWg1.Add(1)
		go func(i int) {
			cque.Add(i * 2)
			CQueueWg1.Done()
		}(i)
	}

	fmt.Println("Removing 100 items from concurrent queue")
	for i := 0; i < LIMIT; i++ {
		CQueueWg2.Add(1)
		go func(i int) {
			cque.Remove()
			CQueueWg2.Done()
		}(i)
	}

	NQueueWg.Wait()
	CQueueWg1.Wait()
	CQueueWg2.Wait()
	fmt.Printf("Concurrent queue size %d", cque.Len())
}
