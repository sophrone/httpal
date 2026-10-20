# httpal

A small command-line HTTP client written in Go.

## Features

- Make HTTP requests
- Print response bodies
- Display response status and headers
- Send custom headers
- Send request bodies
- Read request bodies from files
- Save responses to files
- Configure request timeouts

## Installation

```bash
go build -o httpal
```

## Usage

```bash
./httpal <url>
```

## Examples

```bash
./httpal [https://example.com](https://example.com)
./httpal -i [https://example.com](https://example.com)
./httpal -H "X-App-Name: httpal" [https://httpbin.org/headers](https://httpbin.org/headers)
./httpal -X POST -d "hello" [https://httpbin.org/post](https://httpbin.org/post)
./httpal -o response.html [https://example.com](https://example.com)
```

## Run tests

```bash
go test ./...
```