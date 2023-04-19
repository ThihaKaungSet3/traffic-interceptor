package api

import (
	"encoding/json"
	"net/http"

	"github.com/patrickmn/go-cache"
)

func SetUpRoutes(cache *cache.Cache) {
	http.HandleFunc("/workings", func(w http.ResponseWriter, r *http.Request) {
		value, found := cache.Get("proxies")
		if !found {
			// Return an error response if the key is not found in the cache
			http.Error(w, "Value not found in cache", http.StatusNotFound)
			return
		}

		// Marshal the value into a JSON-encoded byte slice
		data, err := json.Marshal(value)
		if err != nil {
			// Return an error response if there was a problem marshaling the value
			http.Error(w, "Error marshaling value", http.StatusInternalServerError)
			return
		}

		// Set the Content-Type header to "application/json" and write the response body
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	})
}
