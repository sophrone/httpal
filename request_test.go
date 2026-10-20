package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestMakeRequest(t *testing.T) {
	testCases := []struct {
		name           string
		method         string
		headers        []string
		data           string
		expectedMethod string
		expectedHeader string
		expectedBody   string
	}{
		{
			name:           "GET request",
			method:         "GET",
			expectedMethod: "GET",
			expectedBody:   "",
		},
		{
			name:           "POST request with body",
			method:         "POST",
			data:           "hello from httpal",
			expectedMethod: "POST",
			expectedBody:   "hello from httpal",
		},
		{
			name:           "request with custom header",
			method:         "GET",
			headers:        []string{"X-App-Name: httpal"},
			expectedMethod: "GET",
			expectedHeader: "httpal",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != testCase.expectedMethod {
					t.Errorf(
						"expected method %s, got %s",
						testCase.expectedMethod,
						r.Method,
					)
				}

				if testCase.expectedHeader != "" {
					gotHeader := r.Header.Get("X-App-Name")

					if gotHeader != testCase.expectedHeader {
						t.Errorf(
							"expected header %q, got %q",
							testCase.expectedHeader,
							gotHeader,
						)
					}
				}

				requestBody, err := io.ReadAll(r.Body)
				if err != nil {
					t.Fatalf("could not read request body: %v", err)
				}

				if string(requestBody) != testCase.expectedBody {
					t.Errorf(
						"expected request body %q, got %q",
						testCase.expectedBody,
						string(requestBody),
					)
				}

				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte("test response"))
			}))
			defer server.Close()

			body, resp, err := makeRequest(
				testCase.method,
				server.URL,
				testCase.headers,
				testCase.data,
				5*time.Second,
			)

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if resp.StatusCode != http.StatusOK {
				t.Errorf(
					"expected status %d, got %d",
					http.StatusOK,
					resp.StatusCode,
				)
			}

			if body != "test response" {
				t.Errorf(
					"expected response body %q, got %q",
					"test response",
					body,
				)
			}
		})
	}
}

func TestMakeRequestInvalidURL(t *testing.T) {
	_, _, err := makeRequest(
		"GET",
		"://invalid-url",
		nil,
		"",
		5*time.Second,
	)

	if err == nil {
		t.Fatal("expected an error, got nil")
	}
}

func TestMakeRequestNotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	}))
	defer server.Close()

	_, resp, err := makeRequest(
		"GET",
		server.URL,
		nil,
		"",
		5*time.Second,
	)

	if err != nil {
		t.Fatalf("expected no request error, got %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf(
			"expected status %d, got %d",
			http.StatusNotFound,
			resp.StatusCode,
		)
	}
}
