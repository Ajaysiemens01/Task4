package WordFile

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// TestConcurrentCount_Valid checks WordCount with valid input files
func TestConcurrentCount_Valid(t *testing.T) {
	// Assuming file1.txt and file2.txt contain words for the test
	files := []string{"files/file1.txt", "files/file2.txt"}
	err := ConcurrentCount(files)
	assert.NoError(t, err, "Expected no error for valid input")
}

// TestConcurrentCount_UnexistFile checks WordCount with nonexistent files
func TestConcurrentCount_UnexistFile(t *testing.T) {
	// file3.txt is assumed to be nonexistent
	files := []string{"files/file1.txt", "files/file3.txt"}
	err := ConcurrentCount(files)
	// We expect no error, but logs should take care of missing files
	assert.NoError(t, err, "Expected no error, missing file should be handled by logs")
}

// TestConcurrentCount_FileWithNoWords checks WordCount with a file that contains no words
func TestConcurrentCount_FileWithNoWords(t *testing.T) {
	// file4.txt is assumed to contain no valid words (could be empty or just punctuation)
	files := []string{"files/file4.txt"}
	err := ConcurrentCount(files)
	expectedError := "no words found in the provided files"
	assert.EqualError(t, err, expectedError, "Expected error because no words were found in the file")
}

// TestConcurrentCount_AnyFileType checks WordCount with a file of any type
func TestConcurrentCount_AnyFileType(t *testing.T) {
	// Assuming file5.pdf is a PDF, this test is to check if the file type is handled properly
	files := []string{"files/file3.md"}
	err := ConcurrentCount(files)
	// If you don't support PDF, this will log and continue without error.
	assert.NoError(t, err, "Expected no error, logs should handle non-text files")
}

// TestConcurrentCount_MixedFiles checks WordCount with a mix of files
//Word docx gies you  wierd words because scanner doesnot extract inner text in word file to do that 
//we use unoffice library
func TestConcurrentCount_MixedFiles(t *testing.T) {
	// files contains a mix of valid and invalid files (e.g., .txt and .md)
	files := []string{"files/file1.txt", "files/file3.md", "files/file3.docx"}
	err := ConcurrentCount(files)
	assert.NoError(t, err, "Expected no error, invalid files should be handled by logs")
}

// TestConcurrentCount_EmptyFile checks WordCount with an empty file
func TestConcurrentCount_EmptyFile(t *testing.T) {
	// Assuming file6.txt is an empty file
	files := []string{"files/empty.txt"}
	err := ConcurrentCount(files)
	expectedError := "no words found in the provided files"
	assert.EqualError(t, err, expectedError, "Expected error because the file is empty")
}