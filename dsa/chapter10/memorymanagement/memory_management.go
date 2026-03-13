package memorymanagement

import "sync"

// 1. Combining Small Objects into Larger Objects
type CombinedStruct struct {
	a int8
	b int8
	c int8
	d int8
}

func NewCombinedStruct(a, b, c, d int8) CombinedStruct {
	return CombinedStruct{a: a, b: b, c: c, d: d}
}

// 2. Promoting Local Variables to Heap
func CreateHeapVariable() *int {
	val := 42   // Escape analysis determines that `val` will be promoted to the heap.
	return &val // `val` is returned as a pointer, requiring it to escape the function scope.
}

// 3. Pre-Allocating Slice Capacity
func PreallocateSlice(size int) []int {
	slice := make([]int, 0, size) // Allocates a slice with a specified capacity.
	for i := 0; i < size; i++ {
		slice = append(slice, i)
	}
	return slice
}

// 4. Using int8 Instead of int
type SmallIntContainer struct {
	a int8
	b int8
}

func NewSmallIntContainer(a, b int8) SmallIntContainer {
	return SmallIntContainer{a: a, b: b}
}

// 5. Skipping GC Scans with Pointer-Free Objects
type NonPointerObject struct {
	a int64
	b float64
}

func NewNonPointerObject(a int64, b float64) NonPointerObject {
	return NonPointerObject{a: a, b: b}
}

// 6. Using FreeLists for Object Reuse

type PoolObject struct {
	Value int
}

var pool = sync.Pool{
	New: func() interface{} {
		return &PoolObject{}
	},
}

func UsePoolObject() {
	obj := pool.Get().(*PoolObject)
	obj.Value = 42
	// Use obj as needed
	pool.Put(obj) // Return obj to the pool for reuse
}
