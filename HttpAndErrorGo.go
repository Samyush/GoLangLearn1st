package main

import (
	"fmt"
	"net/http"
)

// will try get method and throw error if any

func contentType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close() //making sure we closed the body

	ctype := resp.Header.Get("Content-Type")
	if ctype == "" { // return error if content-type header not found
		return "", fmt.Errorf("cant find the Content-Type header")
	}
	return ctype, nil
}
