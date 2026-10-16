package main

import (
	"fmt"
	"io"
	"net/http"
)

func makeRequest(method string, url string, headers []string) (string, *http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return "", nil, err
	}

	for _, header := range headers {
		var key, value string

		_, err := fmt.Sscanf(header, "%[^:]: %[^\n]", &key, &value)
		if err != nil {
			return "", nil, fmt.Errorf("invalid header format: %q", header)
		}

		req.Header.Set(key, value)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
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
