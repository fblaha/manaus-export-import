package rest

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

// NewWriter constructor
func NewWriter(httpClient httpClient, url string, contentType string) Writer {
	return Writer{httpClient: httpClient, url: url, contentType: contentType}
}

// Writer writes to directory by id
type Writer struct {
	httpClient  httpClient
	url         string
	contentType string
}

// Write do http post for given url
func (w Writer) Write(id string, data []byte) error {
	log.Printf("posting data (id: %s) to %s", id, w.url)
	req, err := http.NewRequest("POST", w.url, bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("unable to create request: %v", err)
	}
	req.Header.Set("Content-Type", w.contentType)
	resp, err := w.httpClient(req)
	if err != nil {
		return fmt.Errorf("unable to send request: %v", err)
	}
	if resp == nil || resp.StatusCode >= 400 {
		return fmt.Errorf("unexpected response: %+v", resp)
	}
	return nil
}
