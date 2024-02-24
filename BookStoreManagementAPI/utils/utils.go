package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// ParseBody reads an HTTP request body and unmarshals the JSON content into the provided variable x.
// It returns an error if reading the body or unmarshalling fails.
func ParseBody(r *http.Request, x interface{}) (err error) {
    body, err := io.ReadAll(r.Body)
    defer r.Body.Close() // Ensure the request body is closed
    if err != nil {
        return  // Return the error to the caller
    }

    if err = json.Unmarshal(body, x); err != nil {
        return  // Return unmarshalling error to the caller
    }

    return nil // No error occurred
}