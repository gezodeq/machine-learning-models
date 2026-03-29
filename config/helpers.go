package machine-learning-models

import (
	"errors"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/golang-collections/collection"
)

// getFilesFromFolder returns a list of files in the given folder
func getFilesFromFolder(folder string) ([]string, error) {
	files, err := filepath.Glob(folder + "/*")
	if err != nil {
		return nil, err
	}
	return files, nil
}

// removeDuplicates removes duplicate elements from a slice
func removeDuplicates[T comparable](input []T) []T {
	return collection.NewSet[T](input).ToSlice()
}

// logError logs an error with the given message
func logError(message string) {
	log.Println("Error:", message)
}

// parseFloat attempts to parse a string to a float64
func parseFloat(str string) (float64, error) {
	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, err
	}
	return val, nil
}

// getFilesInDirectory returns a list of files in the directory given by 'directory'
func getFilesInDirectory(directory string) ([]string, error) {
	files, err := os.ReadDir(directory)
	if err != nil {
		logError("Error reading directory")
		return nil, err
	}
	var fileNames []string
	for _, file := range files {
		if !file.IsDir() {
			fileNames = append(fileNames, file.Name())
		}
	}
	return fileNames, nil
}

// getFilesMatchingPattern returns a list of files matching the given pattern in the given directory
func getFilesMatchingPattern(directory string, pattern string) ([]string, error) {
	files, err := filepath.Glob(filepath.Join(directory, pattern))
	if err != nil {
		logError("Error finding matching files")
		return nil, err
	}
	return files, nil
}

// isFile returns true if the path is a file, false otherwise
func isFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// isDirectory returns true if the path is a directory, false otherwise
func isDirectory(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// calculateMean returns the mean of a list of numbers
func calculateMean(numbers []float64) (float64, error) {
	sum, err := sumNumbers(numbers)
	if err != nil {
		return 0, err
	}
	return sum / float64(len(numbers)), nil
}

// sumNumbers returns the sum of a list of numbers
func sumNumbers(numbers []float64) (float64, error) {
	var sum float64
	for _, num := range numbers {
		sum += num
		if math.IsNaN(num) {
			return 0, errors.New("NaN encountered in numbers")
		}
	}
	return sum, nil
}

// splitString splits a string at the given delimiter
func splitString(str string, delimiter string) (string, string) {
	parts := strings.SplitN(str, delimiter, 2)
	if len(parts) < 2 {
		return "", str
	}
	return parts[0], parts[1]
}