package Integer

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// TestConcurrentSum_Valid checks concurrent sum with valid input
func TestConcurrentSum_Valid(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expectedSum := 55
	result, err := ConcurrentSum(numbers, 1)

	assert.NoError(t, err, "Expected no error for valid input")
	assert.Equal(t, expectedSum, result, "Sum does not match expected result")
}

// TestConcurrentSum_EmptySlice checks for error when input slice is empty
func TestConcurrentSum_EmptySlice(t *testing.T) {
	numbers := []int{}
	result, err := ConcurrentSum(numbers, 2)
	expectedError := "input slice is empty"
	assert.EqualError(t, err,expectedError, "Expected an error for empty slice")
	assert.Equal(t, 0, result, "Expected sum to be 0 for empty slice")
}

// TestConcurrentSum_NegativeSplitSize checks for error when splitSize is negative
func TestConcurrentSum_NegativeSplitSize(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	result, err := ConcurrentSum(numbers, -1)
	expectedError := "splitSize must be greater than zero"
	assert.EqualError(t, err, expectedError, "Expected Error for Negitive Split size")
	assert.Equal(t, 0, result, "Expected sum to be 0 for invalid splitSize")
}

// TestConcurrentSum_LargeSplitSize checks behavior when splitSize is larger than the slice length
func TestConcurrentSum_LargeSplitSize(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	expectedSum := 15
	result, err := ConcurrentSum(numbers, 10)

	assert.NoError(t, err, "Expected no error for valid input")
	assert.Equal(t, expectedSum, result, "Sum does not match expected result")
}

// TestConcurrentSum_SplitSizeOne checks if splitSize=1 works correctly (acts like sequential sum)
func TestConcurrentSum_SplitSizeOne(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	expectedSum := 15
	result, err := ConcurrentSum(numbers, 1)

	assert.NoError(t, err,"Expected no error for valid input")
	assert.Equal(t, expectedSum, result, "Sum does not match expected result")
}
