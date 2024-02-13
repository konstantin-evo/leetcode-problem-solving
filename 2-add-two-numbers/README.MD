# Add Two Numbers

This Go program demonstrates the [Add Two Numbers](https://leetcode.com/problems/add-two-numbers/) leetcode problem
solution.

## Problem Description

This Go program solves the problem of adding two numbers represented by linked lists. Each node in the linked list
represents a digit of a non-negative integer, with the least significant digit coming first.

The task is to add the two numbers and return the sum as a linked list.

## Usage

To use this program, follow these steps:

1. Ensure you have Go installed on your system
2. Clone this repository `git clone https://github.com/konstantin-evo/leetcode-problem-solving.git`
3. Navigate to the directory containing the `main.go` file `cd ./2-add-two-numbers`
4. Run `go run main.go` to execute the program

## Algorithm Description

The `addTwoNumbers` function takes two singly-linked lists representing non-negative integers and returns their sum as a
singly-linked list.

1. Initialize variables to track the carry and the head of the result list.
2. Iterate through the lists until both are exhausted and there's no carry left.
3. Get the current values from both lists, taking care of the possibility of one list being shorter than the other.
4. Calculate the sum and update the carry.
5. Create a new node for the sum and append it to the result list.
6. Return the head of the result list (excluding the dummy head node).

## Example with Debug Table

| Step | l1   | l2   | carry | sum | result                              | Description                                         |
|------|------|------|-------|-----|-------------------------------------|-----------------------------------------------------|
| Init | 6 -> | 5 -> | 0     | N/A | [Empty node]                        | Initialize lists and variables                      |
| 1    | 6 -> | 5 -> | 1     | 11  | [Empty node] -> 1 ->                | Initial sum of 6 + 5 + carry = 11, carry = 1        |
| 2    | 0 -> | 1 -> | 0     | 2   | [Empty node] -> 1 -> 2 ->           | Sum of 0 + 1 + carry = 2, carry = 0                 |
| 3    | 1    | 9    | 1     | 11  | [Empty node] -> 1 -> 2 -> 0 ->      | Sum of 1 + 9 + carry = 11, carry = 1                |
| 4    | N/A  | N/A  | 1     | 1   | [Empty node] -> 1 -> 2 -> 0 -> 1 -> | Sum of 0 + 0 + carry = 1, carry = 0                 |
| End  | N/A  | N/A  | N/A   | N/A | 1 -> 2 -> 0 -> 1 ->                 | End of iteration, drop first node and return result |

In this example, we have two input linked lists: `6 -> 0 -> 1` and `5 -> 1 -> 9`. The sum of these two lists results
in `1 -> 2 -> 0 -> 1`, which is returned as the output.
