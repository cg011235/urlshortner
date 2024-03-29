package rangecounter

import (
	"strings"
	"testing"
)

// Test case:
// 1. Initialise the range counter instance.
// 2. Call Get() on range counter.
// 3. Verify that value is within the baseChars range.
// 4. Simulate reaching the max value and ensure proper error is handled.
func TestRangeCounter(t *testing.T) {
	min := int64(1)
	max := int64(100)
	rc := NewRangeCounter(min, max)

	val, err := rc.Get()
	if err != nil {
		t.Fatalf("Failed to get the first counter value: %v", err)
	}
	if len(val) != 7 {
		t.Errorf("Expected the counter value to be 7 characters, got %d", len(val))
	}

	for _, char := range val {
		if !strings.ContainsRune(baseChars, char) {
			t.Errorf("Counter value contains invalid character: %v", char)
		}
	}

	rc.current = max + 1
	_, err = rc.Get()
	if err == nil {
		t.Errorf("Expected an error when exceeding the max counter value")
	}
}

// Test case:
// 1. Convert a number to a 7-character string.
// 2. Verify the resulting string only contains valid characters.
// 3. Verify padding with the first character of baseChars if the result is shorter than 7 characters.
func TestNumberToString(t *testing.T) {
	num := int64(1234567)
	str := numberToString(num)
	if len(str) != 7 {
		t.Errorf("Expected the string to be 7 characters, got %d", len(str))
	}

	for _, char := range str {
		if !strings.ContainsRune(baseChars, char) {
			t.Errorf("The resulting string contains invalid character: %v", char)
		}
	}
	num = int64(1) // A number that will result in a short string.
	str = numberToString(num)
	if len(str) != 7 {
		t.Errorf("Expected the string to be 7 characters after padding, got %d", len(str))
	}
	if str[0] != baseChars[0] {
		t.Errorf("Expected the string to be padded with the first character of baseChars, got %v", str[0])
	}
}
