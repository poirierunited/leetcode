# LeetCode Solutions

This repository contains solutions to various LeetCode problems implemented in Go.

## Project Structure

```
.
├── cmd/                    # Command-line applications
│   └── problems/          # Individual problem solutions
├── pkg/                   # Shared packages and utilities
│   └── utils/            # Common utilities used across problems
├── testdata/             # Test data for problems
└── Makefile              # Build and run commands
```

## Adding a New Problem

1. Create a new directory under `cmd/problems/` with your problem name (e.g., `cmd/problems/twosum/`)
2. Create a `main.go` file in the problem directory
3. Implement your solution
4. Add test cases in the same file

## Running Problems

Use the Makefile commands:

```bash
# Run a specific problem
make run PROBLEM=twosum

# Run all problems
make run-all

# Run tests for a specific problem
make test PROBLEM=twosum

# Run all tests
make test-all
```

## Problem Template

Each problem should follow this structure:

```go
package main

import (
    "fmt"
)

func solveProblem(input ...interface{}) interface{} {
    // Your solution here
    return nil
}

func main() {
    // Test cases
    testCases := []struct {
        input    []interface{}
        expected interface{}
    }{
        // Add your test cases here
    }

    // Run test cases
    for i, tc := range testCases {
        result := solveProblem(tc.input...)
        if result != tc.expected {
            fmt.Printf("Test case %d failed: expected %v, got %v\n", i+1, tc.expected, result)
        } else {
            fmt.Printf("Test case %d passed!\n", i+1)
        }
    }
}
``` 