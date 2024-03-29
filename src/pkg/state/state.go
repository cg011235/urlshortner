// Package state provides functionality to maintain in-memory state of
// URLs mapping using simple methods
package state

import (
	"errors"
	"fmt"
	"net/url"
	"sort"
	"sync"
)

var (
	ErrURLParse      = errors.New("error parsing URL")
	ErrEntryNotFound = errors.New("entry not found")
)

// State manages mappings between long URLs and short URLs, along with domain counts.
type State struct {
	longToShortMap map[url.URL]string
	shortToLongMap map[string]*url.URL
	domainCountMap map[string]int64
	mutex          sync.Mutex
}

// NewState initializes and returns a new State instance with initialized maps.
func NewState() *State {
	return &State{
		longToShortMap: make(map[url.URL]string),
		shortToLongMap: make(map[string]*url.URL),
		domainCountMap: make(map[string]int64),
	}
}

// LookupShort looks up the long URL for a given short URL and returns it if found.
func (s *State) LookupShort(shortUrl string) (*url.URL, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if val, ok := s.shortToLongMap[shortUrl]; ok {
		return val, nil
	}
	return nil, ErrEntryNotFound
}

// LookupLong looks up the short URL for a given long URL and returns it if found.
func (s *State) LookupLong(longUrl string) (string, error) {
	u, err := url.Parse(longUrl)
	if err != nil {
		return "", fmt.Errorf("%w: %v", ErrURLParse, err)
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if val, ok := s.longToShortMap[*u]; ok {
		return val, nil
	}
	return "", ErrEntryNotFound
}

// Insert creates a mapping from a long URL to a short URL and updates domain counts.
func (s *State) Insert(shortURL string, longUrl string) error {
	u, err := url.Parse(longUrl)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrURLParse, err)
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.longToShortMap[*u] = shortURL
	s.shortToLongMap[shortURL] = u
	s.domainCountMap[u.Host]++
	return nil
}

func (s *State) TopDomainsWithCount(n int) []struct {
	Domain string `json:"domain"`
	Count  int64  `json:"count"`
} {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	var domainCounts []struct {
		Domain string `json:"domain"`
		Count  int64  `json:"count"`
	}

	for domain, count := range s.domainCountMap {
		domainCounts = append(domainCounts, struct {
			Domain string `json:"domain"`
			Count  int64  `json:"count"`
		}{
			Domain: domain,
			Count:  count,
		})
	}

	sort.Slice(domainCounts, func(i, j int) bool {
		return domainCounts[i].Count > domainCounts[j].Count
	})

	if len(domainCounts) > n {
		return domainCounts[:n]
	}
	return domainCounts
}
