package main

import (
	"fmt"
	"time"
	"Task4/packages/Integer"
	// "log"
    // "Task4/packages/WordFile" 
)

// measureExecutionTime runs a sum function and measures its execution time
func measureExecutionTime(f Integer.SumCalculator, data []int, label string) {
	start := time.Now()
	sum, err := f(data)
	duration := time.Since(start)

	if err != nil {
		fmt.Printf("%s Error: %v\n", label, err)
		return
	}
	fmt.Printf("%s Result: %d, Time: %v\n", label, sum, duration)
}


func main() {
	// Generate a large slice of intSlice
	
	intSlice := make([]int,1000000)
	for i := 1; i <= 1000000; i++ {
		intSlice = append(intSlice, i)
	}
	// Measure execution time of both functions
	measureExecutionTime(Integer.NormalSum, intSlice, "Normal Sum")
	measureExecutionTime(func(nums []int) (int, error) {
		return Integer.ConcurrentSum(nums, 4)
	}, intSlice, "Concurrent Sum")
}


// func main() {

// 	// Define file paths (assuming they are in the same folder as the executable)
// 	files := []string{"files/file1.txt", "files/file3.md"}

// 	// Call the concurrent word count function
// 	err := WordFile.ConcurrentCount(files)
// 	if err != nil {
// 		log.Fatalf("Error: %v", err)
// 	} else {
// 		fmt.Println("Word count completed successfully.")
// 	}
// }
