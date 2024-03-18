# Problem
A **word machine** is a system that performs a sequence of simple operations on a stack of integers given. Initially the stack is empty. The sequence of operations is given as a string. Operations are separated by single spaces. The following operations may be specified:

- an integer X (from 0 to 2^20 - 1): the machine pushes X onto the stack;
- "POP": the machine removes the topmost number from the stack;
- "DUP": the machine pushes a duplicate of the topmost number onto the stack;
- "+": the machine pops the two topmost elements from the stack, adds them together and pushes the sum onto the stack;
- "-": the machine pops the two topmost elements from the stack, subtracts the second one from the first (topmost) one and pushes the difference onto the stack.

After processing all the operations, the machine returns the topmost value from the stack.

The machine processes 20-bit unsigned integers (numbers from 0 to 2^20 - 1). An overflow in addition or underflow in subtraction causes an error. The machine also reports an error when it tries to perform an operation that expects more numbers on the stack than the stack actually contains. Also, if, after performing all the operations, the stack is empty, the stack is empty, the machine returns the error.

For example, given a string "13 DUP 4 POP 5 DUP + DUP + -", the machine performs the following operations:

| Operation | Comment | Stack |
|-----------|---------|-------|
| (initial) | [empty] |
| "13"      | push 13 | 13    |
| "DUP"     | duplicate 13 | 13, 13 |
| "4"       | push 4 | 13, 13, 4 |
| "POP"     | pop 4 | 13, 13 |
| "5"       | push 5 | 13, 13, 5 |
| "DUP"     | duplicate 5 | 13, 13, 5, 5 |
| "+"       | add 5 and 5 | 13, 13, 10 |
| "DUP"     | duplicate 10 | 13, 13, 10, 10 |
| "+"       | add 10 and 10 | 13, 13, 20 |
| "-"       | subtract 13 from 20 | 13, 7 |

Finally, the machine will return 7.

Given a string "5 6 + -", the machine reports an error, since, after the addition, there is only one number on the stack, but the subtraction operation expects two.

Given a string "3 DUP 5 - -", the machine reports an error, since the second subtraction yields a negative result.

Write a function: `def solution(S)`, that, given a string S containing a sequence of operations for the word machine, returns the result the machine would return after processing the operations. The function should return -1 if the machine would report an error.

For example, given string S = "13 DUP 4 POP 5 DUP + DUP + -" the function should return 7, as explained in the example above. Given string S = "5 6 + -" or S = "3 DUP 5 - -" the function should return -1.

Assume that:
- the length of S is within the range (0..2,000);
- S is a valid sequence of word machine operations.

Write a function that, given a string S consisting of N characters containing input for the stack machine, returns the result the machine would return if given this string. The function should return -1 if the machine would report an error when processing the string.

In your solution, focus on correctness. The performance of your solution will not be the focus of the assessment.
# Solution
To solve this problem, you can follow these high-level steps:

1. Initialize an empty stack to simulate the stack machine's behavior.
2. Split the input string S into individual words or tokens. You can use the split() method to split the string based on whitespace.
3. Iterate through each token in the split string.
4. For each token, check if it represents a valid operation for the stack machine. Valid operations include:
     - Numeric values: Push the value onto the stack.
     - "DUP": Duplicate the top value on the stack and push the duplicate onto the stack.
     - "POP": Remove the top value from the stack.
     - "+": Pop the top two values from the stack, add them, and push the result onto the stack.
     - "-": Pop the top two values from the stack, subtract the second value from the first, and push the result onto the stack.
5. After processing all the tokens, check the state of the stack. If the stack is empty, return -1 to indicate an error. Otherwise, return the top value on the stack as the result of the stack machine's computation.

Here's a high-level example of how this solution would work:

Input: "13 DUP 4 POP 5 DUP + DUP + -"

- Initialize an empty stack: []
- Split the input string into tokens: ["13", "DUP", "4", "POP", "5", "DUP", "+", "DUP", "+", "-"]
- Iterate through each token:
    - "13": Push 13 onto the stack: [13]
    - "DUP": Duplicate the top value on the stack (13) and push the duplicate onto the stack: [13, 13]
    - "4": Push 4 onto the stack: [13, 13, 4]
    - "POP": Remove the top value from the stack: [13, 13]
    - "5": Push 5 onto the stack: [13, 13, 5]
    - "DUP": Duplicate the top value on the stack (5) and push the duplicate onto the stack: [13, 13, 5, 5]
    - "+": Pop the top two values from the stack (5, 5), add them (10), and push the result onto the stack: [13, 13, 10]
    - "DUP": Duplicate the top value on the stack (10) and push the duplicate onto the stack: [13, 13, 10, 10]
    - "+": Pop the top two values from the stack (10, 10), add them (20), and push the result onto the stack: [13, 13, 20]
    - "-": Pop the top two values from the stack (20, 13), subtract the second value from the first (7), and push the result onto the stack: [13, 7]
- After processing all the tokens, check the state of the stack. The top value is 7, so the result of the stack machine's computation is 7.
