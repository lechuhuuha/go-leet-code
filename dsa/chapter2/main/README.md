1. What is the name of the method to get the size of an array?

- len

2. How do you find the capacity of the slice?

- cap

3. How do you initialize the 2D slice of a string type?

```go
	test := make([][]int, 0, 3)
    test := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
```

4. How do you add an element to the slice?

```go
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	matrix = append(matrix, []int{})
```

5. Using code, can you demonstrate how to create a map of key strings and value strings?
   Initialize the map with keys and values in the code, iterate them in a
   loop, and print the keys and values in the code.

```go
	stringMap := make(map[string]string, 0)
	stringMap["test"] = "test"
	stringMap["test1"] = "test1"
	stringMap["test2"] = "test2"
	stringMap["test3"] = "test3"

	for k, v := range stringMap {
		fmt.Println(k, v)
	}
```

6. How do you delete a value in a map?

```go
	delete(stringMap, "test1")
```

7. What parameters are required for getting a database connection?

- URL, username, password, database

8. Which sql.Rows class method makes it possible to read the attributes of the
   entity in a table?

- method next() to advance from row to row.

9. What does defer do when a database connection is closed?
- defer run when the return of the func is called, mutiple defer will run last-in-first-out order. 

10. Which method allows the sql.DB class to create a prepared statement?

- PrepareContext()
