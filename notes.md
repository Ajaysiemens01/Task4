	content, err := os.ReadFile(fileName)
			if err != nil {
			log.Printf("Error reading file %s: %v\n", fileName, err)
			return
		}

		lines := strings.Split(string(content), "\n")
		for _, line := range lines {
			words := strings.Fields(strings.ToLower(line)) // Convert to lowercase for case insensitivity

			mutex.Lock()
			for _, word := range words {
				word = strings.Trim(word, ".,!?\"'()") // Remove punctuation
				wordCountMap[word]++
			}
			mutex.Unlock()
		}

Concurrency is useful when:

Processing large data chunks 
CPU-intensive operations like matrix multiplication, encryption, image processing.
I/O-bound tasks like downloading multiple files, database queries, API calls.
Utilizing multiple CPU cores effectively when the workload is heavy.
