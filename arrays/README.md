
# DSA Tutorial: Arrays in Go

This repository is a step-by-step tutorial to learn arrays in Go, designed for interview preparation. It includes basic, intermediate, and advanced problems with solutions and explanations.

---

## Key Concepts of Arrays

### 1. **Definition and Basics**
- An array is a collection of elements, all of the same type, stored in contiguous memory locations.
- Arrays in Go have a fixed size, and the size is part of the array's type.

#### Declaration and Initialization
```go
var arr [5]int // Declare an array of size 5
data := [5]int{1, 2, 3, 4, 5} // Initialize an array with values
```

#### Accessing Elements
```go
arr[0] = 10 // Update the first element
fmt.Println(arr[0]) // Access the first element
```

### 2. **Key Properties of Arrays**
- **Length**: Use `len()` to get the number of elements in an array.
- **Indexing**: Arrays are zero-indexed.
- **Type-Safety**: All elements must be of the same type.

### 3. **Traversing an Array**
- Use a `for` loop or `range` to iterate through elements.
```go
for i := 0; i < len(arr); i++ {
    fmt.Println(arr[i])
}

for idx, val := range arr {
    fmt.Printf("Index: %d, Value: %d\n", idx, val)
}
```

### 4. **Arrays vs. Slices**
- Arrays have a fixed size, while slices are dynamic and more commonly used in Go.
- A slice is a view over an array and supports resizing and other operations.

---

## Problem Categories

### Basic (Fundamental Concepts and Operations)
1. Reverse an Array
2. Find Maximum and Minimum Element
3. Sum and Average of Array Elements
4. Check if Array is Sorted
5. Merge Two Sorted Arrays
6. Count the Number of Even and Odd Elements
7. Find the Index of a Target Element
8. Find the Frequency of Each Element
9. Remove an Element from an Array
10. Create a Copy of an Array

### Intermediate (More Complex Logic and Patterns)
1. Rotate an Array
2. Find Second Largest Element
3. Find All Pairs with a Given Sum
4. Remove Duplicates from Sorted Array
5. Longest Subarray with Equal 0s and 1s
6. Find Missing Number in a Series
7. Sort an Array of 0s, 1s, and 2s
8. Find the Intersection of Two Arrays
9. Find the Union of Two Arrays
10. Find the Majority Element (Appears More than n/2 Times)

### Advanced (Challenging and Optimized Problems)
1. Maximum Product Subarray
2. Subarray with Given Sum
3. Count Distinct Elements in Every Window of Size `k`
4. Find Missing and Repeating Numbers
5. Trapping Rainwater Problem
6. Longest Consecutive Sequence
7. Find the Smallest Subarray with a Given Sum
8. Maximum Sum Subarray of Size `k`
9. Merge Overlapping Intervals
10. Median of Two Sorted Arrays

---

## How to Use

1. Clone the repository.
2. Navigate to the relevant problem folder.
3. Run the Go file or tests using:
```bash
go run <filename>.go
```