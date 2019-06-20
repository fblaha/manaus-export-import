package rest

import (
	"context"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"time"
)

type urlLoader func(url string) ([]byte, error)

// LoadURL loads bytes from URL
func LoadURL(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create request")
	}

	ctx, cancel := context.WithTimeout(req.Context(), 30*time.Second)
	defer cancel()

	req = req.WithContext(ctx)

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "http request failed")
	}

	return ioutil.ReadAll(res.Body)

}
