1. Which method of the sort.Sort interface returns the size of the elements to be
   sorted?

- len

```go
func (a SortByAge) Len() int           { return len(a) }
```

2. Which function needs to be passed to the sort.Slice method to sort a slice?

- less - its combine two element to each other

```go
func (a SortByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
```

3. What does the swap method do to the elements at the i and j indices?

- its swap the position of i to j

```go
func (a SortByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
```

4. What is the default order for sorting elements using sort.Sort?

- Sort sorts data in ascending order as determined by the Less method.
- It makes one call to data.Len to determine n and O(n\*log(n)) calls to data.Less and data.Swap.
- The sort is not guaranteed to be stable.

5. How do you implement ascending and descending sorting with sort.Slice?

- Ascending Sort:
  - To sort in ascending order, the comparison function should return true if the element at index i is less than the element at index j.
- Descending Sort:
  - To sort in descending order, the comparison function should return true if the element at index i is greater than the element at index j.

6. How do you sort an array and keep the original order of the elements?

- SliceStable sorts the slice x using the provided less function, keeping equal elements in their original order. It panics if x is not a slice.

```go
	people := []struct {
		Name string
		Age  int
	}{
		{"Alice", 25},
		{"Elizabeth", 75},
		{"Alice", 75},
		{"Bob", 75},
		{"Alice", 75},
		{"Bob", 25},
		{"Colin", 25},
		{"Elizabeth", 25},
	}

	// Sort by name, preserving original order
	sort.SliceStable(people, func(i, j int) bool { return people[i].Name < people[j].Name })
	fmt.Println("By name:", people)
```

7. Which interface is used to reverse the order of the data?

- Use less in sort.Reverse

8. Show an example of sorting a slice.

```go
type SortBy []int

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i] < a[j] }

func test() {
	sortBy := SortBy{
		1, 2, 3,
	}

	sort.Sort(sortBy)
}
```

9. Which method is called to add elements to an unordered list?

- The AddtoHead method adds the node to the head of the unordered list.
- It will make the headNode point to a new node created with property, and the nextNode points
  to the current headNode of the unordered list

10. Write a code example of an unordered list of floats.

```go
type UnorderedList struct {
	headNode *NodeFloat
}

func (u *UnorderedList) AddToHead(p float32) {
	node := NodeFloat{
		property: p,
	}
	if u.headNode != nil {
		node.nextNode = u.headNode
	}
	u.headNode = &node
}

func (u *UnorderedList) IterateList() {
	for node := u.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.property)
	}
}

type NodeFloat struct {
	nextNode *NodeFloat
	property float32
}

```
