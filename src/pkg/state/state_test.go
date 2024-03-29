package state

import (
	"testing"
)

// Test case:
// 1. Verify inserting a new URL mapping
// 2. Verify looking up a long URL by its short form
// 3. Verify looking up a short URL by its long form
// 4. Verify looking up a non-existent short URL
// 5. Verify looking up a non-existent long URL
func TestState(t *testing.T) {
	state := NewState()
	longURL := "http://example.com"
	shortURL := "Exmpl"
	if err := state.Insert(shortURL, longURL); err != nil {
		t.Fatalf("Insert failed: %v", err)
	}
	resLongURL, err := state.LookupShort(shortURL)
	if err != nil {
		t.Fatalf("LookupShort failed: %v", err)
	}
	if resLongURL.String() != longURL {
		t.Errorf("Expected long URL %s, got %s", longURL, resLongURL)
	}
	resShortURL, err := state.LookupLong(longURL)
	if err != nil {
		t.Fatalf("LookupLong failed: %v", err)
	}
	if resShortURL != shortURL {
		t.Errorf("Expected short URL %s, got %s", shortURL, resShortURL)
	}
	_, err = state.LookupShort("NonExistent")
	if err == nil {
		t.Errorf("Expected error when looking up non-existent short URL, got nil")
	}
	_, err = state.LookupLong("http://nonexistent.com")
	if err == nil {
		t.Errorf("Expected error when looking up non-existent long URL, got nil")
	}
}

func TestURLParsingError(t *testing.T) {
	state := NewState()
	invalidURL := "http:///example.com" // Invalid URL
	if err := state.Insert("short", invalidURL); err == nil {
		t.Errorf("Expected URL parsing error, got nil")
	}
}
