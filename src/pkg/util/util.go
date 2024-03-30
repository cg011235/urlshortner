package util

import (
	"errors"
	"net/url"
	"regexp"
)

// validateLongURL checks if the given URL is well-formed and includes a scheme and host.
// It returns an error if the URL is invalid.
func ValidateLongURL(longURL string) error {
	parsedURL, err := url.ParseRequestURI(longURL)
	if err != nil {
		return err
	}
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return errors.New("URL must have an http or https scheme")
	}
	return nil
}

// validateShortURL checks if the given short URL matches the expected pattern.
func ValidateShortURL(shortURL string) error {
	pattern := `^[a-zA-Z0-9]{7}$`
	re, err := regexp.Compile(pattern)
	if err != nil {
		return errors.New("failed to compile short URL validation pattern")
	}
	if !re.MatchString(shortURL) {
		return errors.New("invalid short URL format")
	}
	return nil
}
