package main

import (
	"flag"
	"fmt"
	"os"
	"time"
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
	data := flag.String("d", "", "request body")
	dataFile := flag.String("f", "", "read request body from file")
	outputFile := flag.String("o", "", "write response body to file")
	timeout := flag.Duration("timeout", 10*time.Second, "request timeout")

	var headers headerFlags
	flag.Var(&headers, "H", "request header")

	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		fmt.Fprintln(
			os.Stderr,
			"Usage: httpal [-i] [-X method] [-H header] [-d data] [-f file] [-o file] [-timeout duration] <url>",
		)
		os.Exit(1)
	}

	if *data != "" && *dataFile != "" {
		fmt.Fprintln(os.Stderr, "Error: use either -d or -f, not both")
		os.Exit(1)
	}

	requestData := *data

	if *dataFile != "" {
		fileData, err := os.ReadFile(*dataFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Could not read body file:", err)
			os.Exit(1)
		}

		requestData = string(fileData)
	}

	body, resp, err := makeRequest(
		*method,
		args[0],
		headers,
		requestData,
		*timeout,
	)
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

	if *outputFile != "" {
		err := os.WriteFile(
			*outputFile,
			[]byte(body),
			0644,
		)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Could not write output file:", err)
			os.Exit(1)
		}

		fmt.Printf("Response saved to %s\n", *outputFile)
		return
	}

	fmt.Print(body)
}
