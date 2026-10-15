package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
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

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Request failed:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if *showInfo {
		fmt.Println(resp.Status)

		for key, values := range resp.Header {
			for _, value := range values {
				fmt.Printf("%s: %s\n", key, value)
			}
		}

		fmt.Println()
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not read response:", err)
		os.Exit(1)
	}

	if resp.StatusCode >= 400 {
		fmt.Fprintln(os.Stderr, "HTTP error:", resp.Status)
		os.Exit(1)
	}

	fmt.Print(string(body))
}
