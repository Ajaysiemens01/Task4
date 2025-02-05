package Integer

import (
	"errors"
	"log"
	"sync"
)

// SumCalculator defines a function type for sum calculations - used to measure execution time
type SumCalculator func([]int) (int, error)

// ConcurrentSum computes sum concurrently with error handling
func ConcurrentSum(intSlice []int, splitSize int) (int, error) {
	if err := validateInput(intSlice, splitSize); err != nil {
		return 0, err
	}

	if splitSize > len(intSlice) {
		log.Printf("Split size is greater than length of slice")
		splitSize = len(intSlice) // Adjust splitSize if too large
	}

	ch := make(chan int)
	var wg sync.WaitGroup

	// Launch goroutines
	for start := 0; start < len(intSlice); start += splitSize {
		end := start + splitSize
		if end > len(intSlice) {
			end = len(intSlice)
		}
		wg.Add(1)
		go computeSubSliceSum(intSlice[start:end], ch, &wg)
	}

	// Close the channel when all goroutines are done
	go func() {
		wg.Wait()
		close(ch)
	}()

	return collectResults(ch), nil
}

// NormalSum computes sum sequentially
func NormalSum(intSlice []int) (int, error) {
	if len(intSlice) == 0 {
		return 0, errors.New("input slice is empty")
	}
	return computeSum(intSlice), nil
}

// validateInput checks if the input slice and split size are valid
func validateInput(intSlice []int, splitSize int) error {
	if len(intSlice) == 0 {
		return errors.New("input slice is empty")
	}
	if splitSize <= 0 {
		return errors.New("splitSize must be greater than zero")
	}
	return nil
}

// computeSubSliceSum computes the sum of a sub-slice and sends it to the channel
func computeSubSliceSum(subSlice []int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- computeSum(subSlice)
}

// computeSum calculates the sum of a slice
func computeSum(slice []int) int {
	sum := 0
	for _, num := range slice {
		sum += num
	}
	return sum
}

// collectResults collects results from the channel and returns the total sum
func collectResults(ch chan int) int {
	totalSum := 0
	for partialSum := range ch {
		totalSum += partialSum
	}
	return totalSum
}