# Two Sum Example

This Go program demonstrates the [Two Sum](https://leetcode.com/problems/two-sum/) leetcode problem solution.

## Problem Description

Given an array of integers `nums` and an integer `target`, return indices of the two numbers such that they add up
to `target`.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

## Usage

To use this program, follow these steps:

1. Ensure you have Go installed on your system
2. Clone this repository `git clone https://github.com/konstantin-evo/leetcode-problem-solving.git`
3. Navigate to the directory containing the `main.go` file `cd ./1-two-sum`
4. Run `go run main.go` to execute the program

## Algorithm Description

The program utilizes a hashmap to efficiently find the solution in a single pass through the array. Here's the
algorithm:

1. Initialize an empty hashmap to store the indices of elements.
2. Iterate through the array, for each element:
   - Calculate the complement (target - current element).
   - Check if the complement exists in the hashmap. If found, return the indices of the two numbers.
   - If not found, add the current element and its index to the hashmap.
3. If no solution is found after iterating through the array, return nil.

## Example

The program will run with a predefined example of `nums1 = [2, 7, 11, 15]` and `target1 = 9`. You can modify these
values in the `main` function of `main.go` to test the program with different inputs.

| Step | Description                          | Value                              |
|------|--------------------------------------|------------------------------------|
| 1    | Create an index map                  | `indexMap = {}`                    |
| 2    | Iterate through the array            |                                    |
|      | `num = 2`                            |                                    |
|      | `difference = 7`                     |                                    |
|      | `7` not found in `indexMap`          | Add `2` to `indexMap` with index 0 |
|      | `num = 7`                            |                                    |
|      | `difference = 2`                     |                                    |
|      | `2` found in `indexMap` with index 0 |                                    |
| 3    | Return result                        | `[0, 1]`                           |
