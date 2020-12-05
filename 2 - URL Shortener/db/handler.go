package main

import (
	"github.com/boltdb/bolt"
	"net/http"
)

type yamlRecords struct {
	Path string `yaml:"path"`
	Url string `yaml:"url"`
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if val, ok := pathsToUrls[r.URL.Path]; ok {
			http.Redirect(w, r, val, http.StatusPermanentRedirect)
		} else {
			fallback.ServeHTTP(w, r)
		}
	})
}

func DbHandler(db *bolt.DB, fallback http.Handler) (http.HandlerFunc, error) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		checkKey := r.URL.Path
		keyToValue := DbFetchForKey(db, checkKey)

		if keyToValue != nil {
			http.Redirect(w, r, string(keyToValue), http.StatusFound)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}), nil
}

func buildMap(yaml []yamlRecords) map[string]string {
	builtMap := map[string]string{}

	for _, record := range yaml {
		builtMap[record.Path] = record.Url
	}

	return builtMap
}
