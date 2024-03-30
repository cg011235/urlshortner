package util

import "testing"

func TestValidateLongURL(t *testing.T) {
	tests := []struct {
		url     string
		wantErr bool
	}{
		{"https://www.example.com", false},
		{"http://example.com", false},
		{"ftp://example.com", true},
		{"www.example.com", true},
		{"", true},
	}

	for _, tt := range tests {
		err := ValidateLongURL(tt.url)
		if (err != nil) != tt.wantErr {
			t.Errorf("ValidateLongURL(%q) error = %v, wantErr %v", tt.url, err, tt.wantErr)
		}
	}
}

func TestValidateShortURL(t *testing.T) {
	tests := []struct {
		url     string
		wantErr bool
	}{
		{"Abc1234", false},
		{"abcdefg", false},
		{"1234567", false},
		{"abcd123", false},
		{"1a2B3c4", false},
		{"abcd*12", true},
		{"abc", true},
		{"12345678", true},
		{"", true},
	}

	for _, tt := range tests {
		err := ValidateShortURL(tt.url)
		if (err != nil) != tt.wantErr {
			t.Errorf("ValidateShortURL(%q) error = %v, wantErr %v", tt.url, err, tt.wantErr)
		}
	}
}
