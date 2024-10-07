1. Give an example where you can use a composite pattern.

- Search operation for files and folders. For a file, it will just look into the contents of the file. For a folder, it will go through all files of that folder to find that keyword

2. For an array of 10 elements with a random set of integers, identify the maximum2.
   and minimum. Calculate the complexity of the algorithm.

- Using linear search. Get the start of the slice as min and max then traverse its through slice to find min and max. Complexity of the algorithm is O(n)

3. To manage the state of an object, which structural pattern is relevant?

- Flyweight to manage the state of an object

4. A window is sub-classed to add a scroll bar to make it a scrollable window.
   Which pattern is applied in this scenario?

- Decorator allows adding new behaviors to objects dynamically by placing them inside special wrapper objects.
- Using decorators you can wrap objects countless number of times since both target objects and decorators follow the same interface. The resulting object will get a stacking behavior of all wrappers.

5. Find the complexity of a binary tree search algorithm.

- It has logarithmic complexity O(log n)

6. Identify the submatrices of 2x2 in a 3x3 matrix. What is the complexity of the6.
   algorithm that you have used?

- Quadratic complexity

7. Explain with a scenario the difference between brute force and backtracking
   algorithms.

- backtracking is able to find and search through result based on decision and backtracking when neccessary. Making the complexity of time is smaller and the input get divided each time
- brute force is able to search through all possible result until a valid solution is found

8. A rules engine uses backtracking to identify the rules affected by the change.
   Show an example where backtracking identifies the affected rules.

9. Draw a flow chart for the algorithm of the calculation of profit-loss given the cost price, selling price, and quantity.

          +-----------------------------+
          |            Start             |
          +-----------------------------+
                       |
          +-----------------------------+
          |     Input Cost Price (CP)    |
          +-----------------------------+
                       |
          +-----------------------------+
          |   Input Selling Price (SP)   |
          +-----------------------------+
                       |
          +-----------------------------+
          |        Input Quantity (Q)    |
          +-----------------------------+
                       |
          +-----------------------------+
          | Is Selling Price (SP) > CP?  |
          +-----------------------------+
                /           \
        Yes  /                 \  No
           /                     \
+------------------------+  +------------------------+
| Profit = (SP - CP) * Q  |  | Is Selling Price (SP)  |
|                        |  |          < CP?         |
+------------------------+  +------------------------+
         |                        /         \
         |                  Yes  /           \  No
         |                     /               \
+------------------------+  +------------------------+
|    Output Profit        |  | Loss = (CP - SP) * Q   |
+------------------------+  +------------------------+
         |                        |
         |              +------------------------+
         |              |    Output Loss         |
         |              +------------------------+
         |                        |
  +-----------------------------+ |
  | No Profit, No Loss (SP = CP) | 
  +-----------------------------+
                       |
          +-----------------------------+
          |              End             |
          +-----------------------------+


10. Write the pseudo code for an algorithm that compares the strings and identifies.
    the substring within a string.

FUNCTION isSubstring(mainString, subString):
    // Get the lengths of the main string and the substring
    SET mainLength = LENGTH(mainString)
    SET subLength = LENGTH(subString)
    
    // Check if the substring is longer than the main string
    IF subLength > mainLength THEN
        RETURN "Substring not found"
    
    // Loop through the main string
    FOR i FROM 0 TO (mainLength - subLength) DO
        // Initialize a flag to track if we found the substring
        SET found = TRUE
        
        // Compare the substring with the corresponding part of the main string
        FOR j FROM 0 TO (subLength - 1) DO
            IF mainString[i + j] != subString[j] THEN
                SET found = FALSE
                BREAK // Exit the inner loop if characters don't match
            END IF
        END FOR
        
        // If the substring was found
        IF found THEN
            RETURN "Substring found at index " + i
        END IF
    END FOR
    
    RETURN "Substring not found"

END FUNCTION

// Example usage
SET mainString = "Hello, World!"
SET subString = "World"
OUTPUT isSubstring(mainString, subString) // Output: Substring found at index 7
