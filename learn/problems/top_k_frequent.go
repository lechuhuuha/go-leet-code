package problems

// Given an integer array nums and an integer k, return the k most frequent elements.
// Input: nums = [1,1,1,2,2,3], k = 2
// Output: [1,2]

type heapNode struct {
	value     int
	frequency int
}

type minHeapNode []*heapNode

func (a minHeapNode) Len() int           { return len(a) }
func (a minHeapNode) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a minHeapNode) Less(i, j int) bool { return a[i].frequency < a[j].frequency }
func (h *minHeapNode) Push(x heapNode) {
	*h = append(*h, &x)

	h.shiftUp(len(*h) - 1)
}
func (h *minHeapNode) Pop() heapNode {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return *item
}

func (h *minHeapNode) ReplaceRoot(node heapNode) {
	(*h)[0] = &node
	h.heapifyDown(0)
}

func (h *minHeapNode) heapifyDown(index int) {
	for {
		smallest := index
		leftChild := 2*index + 1
		rightChild := 2*index + 2
		if leftChild < len(*h) && (*h)[leftChild].frequency < (*h)[smallest].frequency {
			smallest = leftChild
		}

		if rightChild < len(*h) && (*h)[rightChild].frequency < (*h)[smallest].frequency {
			smallest = rightChild
		}

		if smallest == index {
			break
		}

		(*h)[index], (*h)[smallest] = (*h)[smallest], (*h)[index]

		index = smallest
	}
}

func (h minHeapNode) shiftUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2

		if h[index].frequency >= h[parent].frequency {
			break
		}

		h[index], h[parent] = h[parent], h[index]
		index = parent
	}
}

// Given an integer array nums and an integer k, return the k most frequent elements.
// Input: nums = [1,1,1,2,2,3], k = 2
// Output: [1,2]
func topKFrequent(nums []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	freqMap := map[int]int{}

	for _, v := range nums {
		freqMap[v]++
	}
	minHeap := minHeapNode{}

	for value, frequency := range freqMap {
		node := heapNode{
			value:     value,
			frequency: frequency,
		}

		if minHeap.Len() < k {
			minHeap.Push(node)
		} else if frequency > minHeap[0].frequency {
			minHeap.ReplaceRoot(node)
		}
	}

	result := []int{}

	for minHeap.Len() > 0 {
		node := minHeap.Pop()
		result = append(result, node.value)
	}

	return result
}
