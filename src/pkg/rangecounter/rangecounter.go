// Package range provides functionality to maintain a counter within a given range
// and returns the next counter value as a 7-character string.
package rangecounter

import (
	"errors"
	"sync"
)

// We use base 62 A-Z, a-z, 0-9
const baseChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type RangeCounter struct {
	min     int64
	max     int64
	mutex   sync.Mutex
	current int64
}

// NewRangeCounter creates a new RangeCounter with the specified range.
func NewRangeCounter(min, max int64) *RangeCounter {
	return &RangeCounter{
		min:     min,
		max:     max,
		current: min,
	}
}

// Get returns the next counter value as a 7-character string, wrapping around if the max is reached.
func (rc *RangeCounter) Get() (string, error) {
	rc.mutex.Lock()
	defer rc.mutex.Unlock()

	if rc.current > rc.max {
		return "", errors.New("exceeded range")
	}

	numStr := numberToString(rc.current)
	rc.current++

	return numStr, nil
}

// numberToString converts a number to a 7-character string based on baseChars encoding.
func numberToString(num int64) string {
	if num == 0 {
		return string(baseChars[0])
	}

	var result []byte
	base := int64(len(baseChars))
	for num > 0 && len(result) < 7 {
		remainder := num % base
		result = append([]byte{baseChars[remainder]}, result...)
		num /= base
	}

	for len(result) < 7 {
		result = append([]byte{baseChars[0]}, result...)
	}

	return string(result)
}
