package main

import (
	"flag"
	"fmt"
	"os"
)

type headerFlags []string

func (h *headerFlags) String() string {
	return fmt.Sprintf("%v", *h)
}

func (h *headerFlags) Set(value string) error {
	*h = append(*h, value)
	return nil
}

func main() {
	showInfo := flag.Bool("i", false, "show response status and headers")
	method := flag.String("X", "GET", "HTTP method")
	var headers headerFlags

	flag.Var(&headers, "H", "request header")

	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "Usage: httpal [-i] [-X method] [-H header] <url>")
		os.Exit(1)
	}

	url := args[0]

	body, resp, err := makeRequest(*method, url, headers)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Request failed:", err)
		os.Exit(1)
	}

	if *showInfo {
		printHeaders(resp)
	}

	if resp.StatusCode >= 400 {
		fmt.Fprintln(os.Stderr, "HTTP error:", resp.Status)
		os.Exit(1)
	}

	fmt.Print(body)
}
