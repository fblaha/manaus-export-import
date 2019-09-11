package rest

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/pkg/errors"
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
	req.Header.Set("Content-Type", w.contentType)
	resp, err := w.httpClient(req)
	if resp == nil || resp.StatusCode >= 400 {
		return fmt.Errorf("unexpected response: %+v", resp)
	}
	return errors.Wrap(err, "unable to post data")
}
