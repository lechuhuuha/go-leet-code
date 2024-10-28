1. How do you ensure a BinarySearchTree is synchronized?

- Synchronization in TreeSets is achieved using a
  sync.RWMutex lock

2. Which method is called to postpone the invocation of a function?

- Using defer keyword

3. How do you define dictionary keys and values with custom types?

- The dictionary has the value DictVal of type string mapped to DictKey

4. How do you find the length of a map?

- Using func len()

5. What keyword is used to traverse a list of treeNodes in a tree?

- inOrderTraverseTree

6. In a Farey sequence, what are the real numbers in the series called?

- the real numbers in the series are called Farey fractions

7. What is a Fibonacci number?

- Fibonacci consists of a list of numbers in which every number is the sum of
  the two preceding numbers.

8. How do you convert an integer into a string?

- strconv.Itoa and fmt.Srpintf
- strconv.Itoa faster

9. What method is used to convert a byte into a string?

- string()

10. What method is called to add elements to a dictionary?

- A Put method
  - takes the key and value
    parameters of the DictKey and DictVal types respectively.
  - The Lock method of the
    dictionary's lock instance is invoked, and the Unlock method is deferred.
  - If there are empty map elements in the dictionary, elements are initialized using make.
    The map elements are set with a key and a value if they are not empty
