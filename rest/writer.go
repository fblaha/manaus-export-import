package rest

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"net/http"
)

// NewWriter constructor
func NewWriter(url string, contentType string) Writer {
	return Writer{url: url, contentType: contentType}
}

// Writer writes to directory by id
type Writer struct {
	url         string
	contentType string
}

// Write do http post for given url
func (w Writer) Write(id string, data []byte) error {
	log.Printf("posting data (id: %s) to %s", id, w.url)
	resp, err := http.Post(w.url, w.contentType, bytes.NewBuffer(data))
	if resp == nil || resp.StatusCode >= 400 {
		return fmt.Errorf("unexpected response: %+v", resp)
	}
	return errors.Wrap(err, "unable to post data")
}
