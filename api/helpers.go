package machinelearningmodels

import (
	"fmt"
	"log"
	"math"
	"strings"
	"time"

	"github.com/edsrzf/mmap-go"
	"github.com/gocarina/gocsv"
)

func ReadCSV(filePath string) ([]map[string]string, error) {
	file, err := mmap.Map(filePath, mmap.RDONLY, 0)
	if err != nil {
		return nil, err
	}

	defer file.Unmap()

	csvParser := gocsv.NewParser(file)
	reader := csvParser Records()

	var header []string
	var rows []map[string]string

	for {
		row, err := reader.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		if header == nil {
			header = row
		} else {
			rows = append(rows, map[string]string(strings.ToLower(header[0])): row[0]))
		}
	}

	return rows, nil
}

func GetDateRange(start, end time.Time) []time.Time {
	var dates []time.Time

	for i := start; i.Before(end); i = i.AddDate(0, 0, 1) {
		dates = append(dates, i)
	}

	return dates
}

func IsPowerOfTwo(n int) bool {
	return n > 0 && math.Log2(float64(n)) == math.Floor(math.Log2(float64(n)))
}

func GetNearestPowerOfTwo(n int) int {
	pow := 1
	for pow < n {
		pow *= 2
	}

	return pow
}

func Log10(value float64) float64 {
	return math.Log(value) / math.Log(10)
}

func IsNaN(value interface{}) bool {
	switch value.(type) {
	case float64:
		return math.IsNaN(value.(float64))
	case int:
		return false
	default:
		return false
	}
}

func LogMessage(level string, message string, args ...interface{}) {
	switch level {
	case "DEBUG":
		log.Printf("[DEBUG] "+fmt.Sprintf(message, args...))
	case "INFO":
		log.Printf("[INFO] "+fmt.Sprintf(message, args...))
	case "WARNING":
		log.Printf("[WARNING] "+fmt.Sprintf(message, args...))
	case "ERROR":
		log.Printf("[ERROR] "+fmt.Sprintf(message, args...))
	case "CRITICAL":
		log.Fatal("[CRITICAL] "+fmt.Sprintf(message, args...))
	default:
		log.Printf("[UNKNOWN] "+fmt.Sprintf(message, args...))
	}
}