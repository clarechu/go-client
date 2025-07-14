package rest

import "fmt"

type RequestError struct {
	StatusCode int    `json:"StatusCode"`
	Method     string `json:"method"`
	Path       string `json:"path"`
	URL        string `json:"url"`
	Body       []byte `json:"body"`
}

func (e *RequestError) Error() string {
	return fmt.Sprintf("error: %s", string(e.Body))
}

func (e *RequestError) GetURL() string {
	return fmt.Sprintf("url: %s", e.URL)
}

func (e *RequestError) GetPath() string {
	return fmt.Sprintf("path: %s", e.Path)
}
