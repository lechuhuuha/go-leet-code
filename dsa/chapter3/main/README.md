1. Where can you use double linked list? Please provide an example.

- Undo/Redo functionality: Text editors or other software where the ability to undo/redo an action is required, as it allows easy traversal back and forth between different states.

2. Which method on linked list can be used for printing out node values?

- IterateList (Its loop through each of the node and print out property)

3. Which queue was shown with channels from the Go language?

- Synchronized queue
  - A synchronized queue consists of elements that need to be processed in a particular
    sequence. Passenger queue and ticket processing queues are types of synchronized queues,
    as follows:

4. Write a method that returns multiple values. What data structure can be used for
   returning multiple values?

- Tuples
  - Tuples are finite ordered sequences of objects

5. Can set have duplicate elements?

- A Set is a linear data structure that has a collection of values that are not repeated.
- A set can store unique values without any particular order.

6. Write a code sample showing the union and intersection of two sets.

```go
func (set *Set) Intersect(anotherSet *Set) *Set {
	var intersectSet = &Set{}
	intersectSet.New()
	var value int
	for value = range set.integerMap {
		if anotherSet.ContainsElement(value) {
			intersectSet.AddElement(value)
		}
	}
	return intersectSet
}

func (set *Set) Union(anotherSet *Set) *Set {
	var unionSet = &Set{}
	unionSet.New()
	var value int
	for value = range set.integerMap {
		unionSet.AddElement(value)
	}
	for value = range anotherSet.integerMap {
		unionSet.AddElement(value)
	}
	return unionSet
}
```

7. In a linked list, which method is used to find the node between two values?

```go

func (l *LinkedList) NodeWithValue(property int) *Node {
	var (
		nodeWith *Node
	)
	for node := l.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			nodeWith = node
			break
		}
	}
	return nodeWith
}
```

8. We have elements that are not repeated and unique. What is the correct data
   structure that represents the collection?

- A Set is a linear data structure that has a collection of values that are not repeated.
- A set can store unique values without any particular order.

9. In Go, how do you generate a random integer between the values 3 and 5?

- math/rand

10. Which method is called to check if an element of value 5 exists in the Set?

```go
func (set *Set) ContainsElement(element int) bool {
	_, exists := set.integerMap[element]
	return exists
}
```
