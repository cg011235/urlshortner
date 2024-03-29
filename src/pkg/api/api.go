package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"urlshortner/src/pkg/rangecounter"
	"urlshortner/src/pkg/state"

	"github.com/gorilla/mux"
)

var inMemState *state.State
var rangeManager *rangecounter.RangeCounter

// RootHandler handles requests to the root path "/"
func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the URL Shortener API!")
}

// ShortenURLHandler handles requests to shorten a URL
func ShortenURLHandler(w http.ResponseWriter, r *http.Request) {
	type request struct {
		OriginalURL string `json:"originalUrl"`
	}

	type response struct {
		ShortURL string `json:"shortUrl"`
	}

	var req request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	if req.OriginalURL == "" {
		http.Error(w, "Original URL is missing", http.StatusBadRequest)
		return
	}
	log.Println("Recieved request for original URL: ", req.OriginalURL)

	shortURL, err := inMemState.LookupLong(req.OriginalURL)
	if err != nil {
		if errors.Is(err, state.ErrURLParse) {
			http.Error(w, "Error looking up URL mapping", http.StatusInternalServerError)
			return
		} else if errors.Is(err, state.ErrEntryNotFound) {
			shortURL, err = GenerateShortURL()
			if err != nil {
				http.Error(w, "Error generating short URL", http.StatusInternalServerError)
				return
			}
			log.Println("Generated short URL as lookup did not find entry: ", shortURL)
		} else {
			http.Error(w, "Error looking up URL mapping", http.StatusInternalServerError)
			return
		}
	} else {
		log.Println("Found short URL in lookup: ", shortURL)
	}

	if err := inMemState.Insert(shortURL, req.OriginalURL); err != nil {
		http.Error(w, "Error inserting URL mapping", http.StatusInternalServerError)
		return
	}
	resp := response{ShortURL: shortURL}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

// GenerateShortURL is a placeholder for your URL shortening logic.
// Implement this function based on your requirements.
func GenerateShortURL() (string, error) {
	val, err := rangeManager.Get()
	if err != nil {
		log.Fatalf("Failed to get the counter value: %v", err)
		return "", err
	}
	return val, nil
}

// RedirectHandler handles requests to the shortened URL path
func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortURL := vars["shorturl"]
	u, err := inMemState.LookupShort(shortURL)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	log.Println("Redirecting short url : ", shortURL, " to original url : ", u.String())
	http.Redirect(w, r, u.String(), http.StatusFound)
}

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	topDomains := inMemState.TopDomainsWithCount(3)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string][]struct {
		Domain string `json:"domain"`
		Count  int64  `json:"count"`
	}{"topDomains": topDomains}); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func NewRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", RootHandler).Methods("GET")
	r.HandleFunc("/shorten", ShortenURLHandler).Methods("POST")
	r.HandleFunc("/{shorturl}", RedirectHandler).Methods("GET")
	r.HandleFunc("/metrics", MetricsHandler).Methods("GET")
	return r
}

func Init() {
	inMemState = state.NewState()
	rangeManager = rangecounter.NewRangeCounter(0, 1000000)
}
