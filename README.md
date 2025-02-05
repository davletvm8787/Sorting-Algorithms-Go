# Sorting-Algorithms-Go
Bubble Sort in Go (with Explanation) Bubble Sort is one of the simplest sorting algorithms. It repeatedly steps through the list, compares adjacent elements, and swaps them if they are in the wrong order. The process continues until the list is sorted.
Bubble Sort in Go (with Explanation)
Bubble Sort is one of the simplest sorting algorithms. It repeatedly steps through the list, compares adjacent elements, and swaps them if they are in the wrong order. The process continues until the list is sorted.

# Algorithm Steps
Start from the first element of the array.
Compare it with the next element. If they are in the wrong order, swap them.
Move to the next pair and repeat the process.
After one full pass, the largest element will be at the last index.
Repeat the process for the remaining elements until the entire list is sorted.
Time Complexity
Best case (already sorted list): 
𝑂
(
𝑛
)
O(n)
Worst and average case: 
𝑂
(
𝑛
2
)
O(n 
2
 )

# Explanation of Code
Function BubbleSort

It takes a slice of integers (arr) as input.
The outer loop iterates through the array (n-1 times).
The inner loop compares adjacent elements and swaps them if necessary.
A swapped flag is used to optimize the algorithm by stopping early if no swaps occur in a full pass.
Function main

Creates an unsorted array.
Calls the BubbleSort function to sort it.
Prints the sorted array.
############### #################### ################### ###################### #################
## Selection Sort in Go (with Explanation)
Selection Sort is a simple sorting algorithm that repeatedly finds the minimum element from the unsorted part of the array and places it at the beginning.

Algorithm Steps
Start with the first element as the minimum.
Compare it with the rest of the array to find the smallest element.
Swap the smallest element with the first element.
Move to the next index and repeat the process for the remaining elements.
Continue until the entire array is sorted.
Time Complexity
Best case: 
𝑂
(
𝑛
2
)
O(n 
2
 )
Worst case: 
𝑂
(
𝑛
2
)
O(n 
2
 )
Average case: 
𝑂
(
𝑛
2
)
O(n 
2
 )
Selection Sort is not efficient for large datasets but works well for small lists.