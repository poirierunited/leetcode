package utils

import (
	"fmt"
	"reflect"
)

// AssertEqual compares two values and returns true if they are equal
func AssertEqual(expected, actual interface{}) bool {
	return reflect.DeepEqual(expected, actual)
}

// PrintTestResult prints the result of a test case
func PrintTestResult(testNum int, expected, actual interface{}, passed bool) {
	if passed {
		fmt.Printf("Test case %d passed!\n", testNum)
	} else {
		fmt.Printf("Test case %d failed: expected %v, got %v\n", testNum, expected, actual)
	}
}

// RunTestCases is a generic function to run test cases
func RunTestCases[T any, U any](testCases []struct {
	Input    T
	Expected U
}, testFunc func(T) U) {
	for i, tc := range testCases {
		result := testFunc(tc.Input)
		passed := AssertEqual(tc.Expected, result)
		PrintTestResult(i+1, tc.Expected, result, passed)
	}
}
