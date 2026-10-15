package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: httpal <url>")
		os.Exit(1)
	}

	url := os.Args[1]

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Request failed:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not read response:", err)
		os.Exit(1)
	}

	if resp.StatusCode >= 400 {
		fmt.Fprintf(os.Stderr, "HTTP error: %s\n", resp.Status)
		os.Exit(1)
	}

	fmt.Print(string(body))
}
