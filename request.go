package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func makeRequest(method string, url string, headers []string, data string) (string, *http.Response, error) {
	var body io.Reader

	if data != "" {
		body = strings.NewReader(data)
	}

	req, err := http.NewRequest(method, url, body)
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

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", resp, err
	}

	return string(responseBody), resp, nil
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
