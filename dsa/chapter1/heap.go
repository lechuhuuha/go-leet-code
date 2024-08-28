package main

type IntegerHeap []int

func (iHeap IntegerHeap) Len() int           { return len(iHeap) }
func (iHeap IntegerHeap) Swap(i, j int)      { iHeap[i], iHeap[j] = iHeap[j], iHeap[i] }
func (iHeap IntegerHeap) Less(i, j int) bool { return iHeap[i] < iHeap[j] }
func (iHeap *IntegerHeap) Push(e interface{}) {
	*iHeap = append(*iHeap, e.(int))
}
func (Iheap *IntegerHeap) Pop() interface{} {
	var heapLen, lastInHeap int
	var prev IntegerHeap = *Iheap

	heapLen = len(prev)
	lastInHeap = prev[heapLen-1]

	*Iheap = prev[0 : heapLen-1]

	return lastInHeap
}
