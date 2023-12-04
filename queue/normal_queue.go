package queue

type NormalQueue struct {
	list []int
}

func NewNormalQueue() NormalQueue {
	return NormalQueue{list: make([]int, 0)}
}

func (nq *NormalQueue) Add(x int) {
	nq.list = append(nq.list, x)
}

func (nq *NormalQueue) Remove() int {
	if len(nq.list) == 0 {
		return -1
	}

	val := nq.list[0]
	// remove 1st element
	nq.list = nq.list[1:]
	return val
}

func (nq *NormalQueue) Len() int {
	return len(nq.list)
}
