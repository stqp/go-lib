package lib

// PriorityQueue is something.
type PriorityQueue struct {
	heapSize int
	heap     []int
	size     int
	reverse  bool
}

// Push pushs a number to queue
func (p *PriorityQueue) Push(x int) {

	// initialization
	if p.heapSize == 0 {
		p.heapSize = 100
	}
	if len(p.heap) == 0 && p.heapSize > 0 {
		p.heap = make([]int, p.heapSize)
	}
	i := p.size
	p.size++
	for i > 0 {
		n := (i - 1) / 2
		if p.reverse {
			if p.heap[n] > x {
				break
			}
		} else {
			if p.heap[n] <= x {
				break
			}
		}

		p.heap[i] = p.heap[n]
		i = n
	}
	p.heap[i] = x
}

// Pop pops a number
func (p *PriorityQueue) Pop() int {
	ret := p.heap[0]
	p.size--
	x := p.heap[p.size]
	i := 0
	for (i*2 + 1) < p.size {
		a := i*2 + 1
		b := i*2 + 2
		if p.reverse {
			if b < p.size && p.heap[b] >= p.heap[a] {
				a = b
			}
			if p.heap[a] < x {
				break
			}
		} else {
			if b < p.size && p.heap[b] < p.heap[a] {
				a = b
			}
			if p.heap[a] >= x {
				break
			}
		}

		p.heap[i] = p.heap[a]
		i = a
	}
	p.heap[i] = x
	return ret
}
