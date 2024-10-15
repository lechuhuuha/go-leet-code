1. Can you give an example where you can use a binary search tree?

- Binary Search trees (and other kinds of balanced trees, like B-trees) are a very good general-purpose data structure for storing database tables and indexes.
- Or website like akinator

2. Which method is used to search for an element in a binary search tree?

- Using search func, if the input is smaller then the root then search from the left and vice versa

3. Which techniques are used to adjust the balance in an AVL tree?

- The balance factor is obtained
  by finding the difference between the heights of the left and right sub-trees.
- Balancing is done using rotation techniques.
- If the balance factor is greater than one, rotation shifts the
  nodes to the opposite of the left or right sub-trees.

4. What is a symbol table?

- A symbol table is present in memory during the program translation process.
- It can be present in program binaries.
- A symbol table contains the symbol's name, location, and
  address.

5. Which class and method are called to generate a binary marshaled hash on the hash class?

- Marshaling (changing the string to an encoded
  form) saves the internal state, which is used for other purposes later.

6. Which container in Go is used to model a circular linked list?

- container/ring

7. How do you create a JSON (indented) from a tree structure? Which class and method are used?

- The tree data
  structure is converted in to JSON in bytes. The JSON bytes are printed after being changed
  to a string
  ```go
  avlTree, _ = json.MarshalIndent(treeNode, "", " ")
  ```

8. How do you compare the sum of hashes?

-     fmt.Println(bytes.Equal(firstHash.Sum(nil), secondHash.Sum(nil)))

9. What is the balance factor in an AVL tree?

- The balance factor is obtained
  by finding the difference between the heights of the left and right sub-trees.

10. How do you identify a row and column in a table?

- In trees

  - row(level)
  - column(position)

- In Graphs

  - row(node level in BFS/DFS)
  - column(position in adjacency matrix)

- In Heaps
  - row(level in tree)
  - column(position in array representation)

Conclusion
Rows: Levels of the structure (e.g., tree depth, BFS levels in graphs).
Columns: Relative positions within that level (e.g., left-to-right positions in trees, connections in adjacency matrices for graphs)
