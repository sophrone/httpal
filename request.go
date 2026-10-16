package main

import (
	"fmt"
	"io"
	"net/http"
)

func makeRequest(url string) (string, *http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", resp, err
	}

	return string(body), resp, nil
}

func printHeaders(resp *http.Response) {
	fmt.Println(resp.Status)

	for key, values := range resp.Header {
		for _, value := range values {
			fmt.Printf("%s: %s\n", key, value)
		}
	}

	fmt.Println()
}
