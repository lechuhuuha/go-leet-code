package problems

// Given an integer array nums and an integer k, return the k most frequent elements.
// Input: nums = [1,1,1,2,2,3], k = 2
// Output: [1,2]

type heapNode struct {
	value     int
	frequency int
	Root      *heapNode
	parent    *heapNode
}

func (a heapNode) Len() int           { return len(a) }
func (a heapNode) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a heapNode) Less(i, j int) bool { return a[i] < a[j] }
func (h *heapNode) Push(x heapNode)
func (h *heapNode) Pop() heapNode

func topKFrequent(nums []int, k int) []int {
	freqMap := map[int]int{}

	for _, v := range nums {
		freqMap[v]++
	}
	minHeap := heapNode{}

	for value, frequency := range freqMap {
		node := heapNode{
			value:     value,
			frequency: frequency,
		}

		if minHeap.Len() < k {
			minHeap.Push(node)
		} else if frequency < minHeap.Root.frequency {
			minHeap.Root.Pop()
			minHeap.Push(node)
		}
	}

	result := []int{}

	for minHeap.Len() > 0 {
		node := minHeap.Pop()
		result = append(result, node.value)
	}

	return result
}
