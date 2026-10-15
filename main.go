package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: httpal <url>")
		return
	}

	url := os.Args[1]

	fmt.Println("URL:", url)
}
