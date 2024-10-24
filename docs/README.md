# Big o notation

**Big O notation is a mathematical concept used to describe the performance or complexity of an algorithm in terms of time and space**. It provides a high-level understanding of how the runtime or memory requirements of an algorithm grow as the size of the input increases. Here’s a brief overview:

### Key Concepts of Big O Notation

1. **Worst-Case Analysis**:
   - Big O notation primarily focuses on the worst-case scenario, which represents the maximum amount of time or space required by an algorithm for any input of size \( n \).

2. **Asymptotic Behavior**:
   - It describes how the runtime or space grows relative to the input size, ignoring constant factors and lower-order terms. This helps in simplifying the analysis and focusing on the most significant factors.

3. **Notation**:
   - Big O notation is expressed as \( O(f(n)) \), where \( f(n) \) is a function that describes the growth rate of the algorithm’s resource consumption (time or space) as a function of the input size \( n \).

### Common Big O Complexities

Here are some common complexities represented in Big O notation:

1. **Constant Time: \( O(1) \)**
   - The algorithm’s execution time is constant and does not change with the input size.
   - **Example**: Accessing an element in an array by index.

2. **Logarithmic Time: \( O(\log n) \)**
   - The execution time grows logarithmically as the input size increases, which is typically seen in algorithms that divide the problem size in each step.
   - **Example**: Binary search in a sorted array.

3. **Linear Time: \( O(n) \)**
   - The execution time grows linearly with the input size. If the input size doubles, the execution time also doubles.
   - **Example**: Finding an element in an unsorted array.

4. **Linearithmic Time: \( O(n \log n) \)**
   - This complexity arises in algorithms that perform a linear scan on the input and perform logarithmic operations (often seen in efficient sorting algorithms).
   - **Example**: Merge sort and quicksort (average case).

5. **Quadratic Time: \( O(n^2) \)**
   - The execution time grows quadratically with the input size. This often occurs in algorithms that involve nested loops over the input.
   - **Example**: Bubble sort, selection sort.

6. **Cubic Time: \( O(n^3) \)**
   - The execution time grows cubically with the input size, often seen in algorithms with three nested loops.
   - **Example**: Some dynamic programming algorithms.

7. **Exponential Time: \( O(2^n) \)**
   - The execution time doubles with each additional element in the input size. This is generally impractical for larger inputs.
   - **Example**: Solving the Traveling Salesman Problem using brute force.

8. **Factorial Time: \( O(n!) \)**
   - The execution time grows factorially with the input size. This is the worst-case scenario and is impractical for even moderately sized inputs.
   - **Example**: Generating all permutations of a set.

### Example Analysis

Here’s a quick example of analyzing the time complexity of a simple algorithm:

```java
void exampleMethod(int[] arr) {
    for (int i = 0; i < arr.length; i++) {           // O(n)
        for (int j = 0; j < arr.length; j++) {       // O(n)
            // Some constant time operation O(1)
        }
    }
}
```

- The outer loop runs \( n \) times.
- The inner loop also runs \( n \) times for each iteration of the outer loop.
- Therefore, the overall time complexity is \( O(n^2) \).

### Summary

- Big O notation provides a framework for evaluating and comparing the efficiency of algorithms.
- It focuses on the upper bound of the runtime or space used, emphasizing the most significant terms while ignoring constant factors and lower-order terms.
- Understanding Big O notation is crucial for optimizing algorithms and ensuring efficient software development.

---
# ***Examples***

Each example will demonstrate how the time complexity of an algorithm can be analyzed using Big O notation.

### 1. **Constant Time: \( O(1) \)**

**Description**: The **execution** time **remains constant regardless of the "input" size**.

**Example**:
```java
public int getFirstElement(int[] arr) {
    return arr[0]; // Always takes the same time
}
```
**Analysis**: No matter how large the array is, accessing the first element takes the same amount of time, so the time complexity is \( O(1) \).

### 2. **Logarithmic Time: \( O(\log n) \)**

**Description**: The execution time **grows "logarithmically" as the input size increases**, typically seen in algorithms that divide the input size in each step.

**Example**: Binary Search
```java
public int binarySearch(int[] arr, int target) {
    int left = 0;
    int right = arr.length - 1;
    
    while (left <= right) {
        int mid = left + (right - left) / 2;
        
        if (arr[mid] == target) {
            return mid; // Target found
        } else if (arr[mid] < target) {
            left = mid + 1; // Search in the right half
        } else {
            right = mid - 1; // Search in the left half
        }
    }
    
    return -1; // Target not found
}
```
**Analysis**: Each iteration cuts the array size in half, leading to a time complexity of \( O(\log n) \).

### 3. **Linear Time: \( O(n) \)**

**Description**: The execution time **grows "linearly" with the input size**.

**Example**: Finding an Element in an Unsorted Array
```java
public boolean containsElement(int[] arr, int target) {
    for (int num : arr) {
        if (num == target) {
            return true; // Element found
        }
    }
    return false; // Element not found
}
```
**Analysis**: In the worst case, the algorithm has to check each element, leading to a time complexity of \( O(n) \).

### 4. **Linearithmic Time: \( O(n \log n) \)**

**Description**: This **complexity arises in algorithms that perform a linear scan while also involving logarithmic operations**.

**Example**: Merge Sort
```java
public void mergeSort(int[] arr) {
    if (arr.length < 2) return; // Base case

    int mid = arr.length / 2;
    int[] left = Arrays.copyOfRange(arr, 0, mid);
    int[] right = Arrays.copyOfRange(arr, mid, arr.length);
    
    mergeSort(left);  // Sort left half
    mergeSort(right); // Sort right half
    merge(arr, left, right); // Merge sorted halves
}

private void merge(int[] arr, int[] left, int[] right) {
    int i = 0, j = 0, k = 0;
    while (i < left.length && j < right.length) {
        if (left[i] <= right[j]) {
            arr[k++] = left[i++];
        } else {
            arr[k++] = right[j++];
        }
    }
    while (i < left.length) arr[k++] = left[i++];
    while (j < right.length) arr[k++] = right[j++];
}
```
**Analysis**: The merge sort divides **the array recursively (logarithmic) and then merges the sorted halves (linear), leading to a time complexity of \( O(n \log n) \)**.

### 5. **Quadratic Time: \( O(n^2) \)**

**Description**: The execution time grows quadratically with the input size, often seen in algorithms with nested loops.

**Example**: Bubble Sort
```java
public void bubbleSort(int[] arr) {
    int n = arr.length;
    for (int i = 0; i < n - 1; i++) {
        for (int j = 0; j < n - 1 - i; j++) {
            if (arr[j] > arr[j + 1]) {
                // Swap arr[j] and arr[j + 1]
                int temp = arr[j];
                arr[j] = arr[j + 1];
                arr[j + 1] = temp;
            }
        }
    }
}
```
**Analysis**: There are two nested loops, each iterating through the array, resulting in a time complexity of \( O(n^2) \).

### 6. **Exponential Time: \( O(2^n) \)**

**Description**: The **execution time doubles with each additional element in the input size**.

**Example**: Fibonacci Sequence (Naive Recursive Implementation)
```java
public int fibonacci(int n) {
    if (n <= 1) {
        return n;
    }
    return fibonacci(n - 1) + fibonacci(n - 2); // Two recursive calls
}
```
**Analysis**: This algorithm has a time complexity of \( O(2^n) \) due to the two recursive calls made for each element, resulting in an exponential growth.

### Summary

- **Constant Time \( O(1) \)**: Execution time does not change with input size.
- **Logarithmic Time \( O(\log n) \)**: Execution time increases logarithmically.
- **Linear Time \( O(n) \)**: Execution time increases linearly.
- **Linearithmic Time \( O(n \log n) \)**: Execution time is linear multiplied by a logarithmic factor.
- **Quadratic Time \( O(n^2) \)**: Execution **time increases quadratically, common in nested loops**.
- **Exponential Time \( O(2^n) \)**: Execution time doubles with each additional input element.

Understanding these complexities helps in analyzing the efficiency of algorithms and making informed choices about which algorithms to use in software development.