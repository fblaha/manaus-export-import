package rest

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type httpClient func(req *http.Request) (*http.Response, error)

// URLLoader loads data from given URL
type URLLoader func(url string) ([]byte, error)

// NewURLLoader constructor
func NewURLLoader(httpClient httpClient) URLLoader {
	return func(url string) ([]byte, error) {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, fmt.Errorf("unable to create request: %v", err)
		}

		res, err := httpClient(req)
		if err != nil {
			return nil, fmt.Errorf("http request failed: %v", err)
		}
		if res == nil || res.StatusCode != 200 {
			return nil, fmt.Errorf("unexpected response: %+v", res)

		}
		return ioutil.ReadAll(res.Body)
	}
}
