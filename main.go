package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	showInfo := flag.Bool("i", false, "show response status and headers")

	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "Usage: httpal [-i] <url>")
		os.Exit(1)
	}

	url := args[0]

	body, resp, err := makeRequest(url)
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
