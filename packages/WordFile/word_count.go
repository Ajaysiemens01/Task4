package WordFile

import (
	"os"
	"fmt"
	"errors"
	"sync"
	"log"
	"bufio"
	"sort"
	"strings"
)

// Extract text from the file, handle different file types (e.g., text, PDF, etc.)
func extractTextFromFile(fileName string) ([]string, error) {
	// Check if file exists and can be opened
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error opening file %s: %v", fileName, err)
	}
	defer file.Close()

	// Scan any type of files in text format
	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words = append(words, strings.Fields(strings.ToLower(line))...)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file %s: %v", fileName, err)
	}
    
	return words, nil
}

// Count words from a given file and update word count map
func countWordsFromFile(fileName string, wordCountMap map[string]int, mutex *sync.Mutex) error {
	words, err := extractTextFromFile(fileName)
	if err != nil {
		return err
	}

	if len(words) == 0 {
		log.Printf("Warning: %s is empty or contains no valid words.\n", fileName)
		return nil
	}

	// Lock map and update word counts
	mutex.Lock()
	defer mutex.Unlock()

	// Clean and count words
	for _, word := range words {
		word = strings.Trim(word, ".,!?\"'()")
		if word != "" {
			wordCountMap[word]++
		}
	}
	return nil
}

// Print the word count map in sorted order
func printSortedWordCount(wordCounts map[string]int) {
	words := make([]string, 0, len(wordCounts))
	for word := range wordCounts {
		words = append(words, word)
	}

	// Sort the words alphabetically
	sort.Strings(words)

	// Print the sorted word counts
	for _, word := range words {
		fmt.Printf("Word: %-10s | Count: %d\n", word, wordCounts[word])
	}
}

// Concurrently process files to count words
func ConcurrentCount(files []string) error {

	if len(files) == 0 {
		return errors.New("no files provided")
	}

	var wg sync.WaitGroup
	var mutex sync.Mutex
	wordCountMap := make(map[string]int)

	// Process each file concurrently
	for _, fileName := range files {


		wg.Add(1)
		go func(file string) {
			defer wg.Done()

			err := countWordsFromFile(file, wordCountMap, &mutex)
			if err != nil {
				log.Printf("Error processing file %s: %v\n", file, err)
			}
		}(fileName)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// If no words were found, return an error
	if len(wordCountMap) == 0 {
		return errors.New("no words found in the provided files")
	}

	// Print sorted word counts
	printSortedWordCount(wordCountMap)

	return nil
}