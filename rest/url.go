package rest

import (
    "fmt"
    "io/ioutil"
    "net/http"

    "github.com/pkg/errors"
)

type httpClient func(req *http.Request) (*http.Response, error)

// URLLoader loads data from given URL
type URLLoader func(url string) ([]byte, error)

// NewURLLoader constructor
func NewURLLoader(httpClient httpClient) URLLoader {
    return func(url string) ([]byte, error) {
        req, err := http.NewRequest("GET", url, nil)
        if err != nil {
            return nil, errors.Wrap(err, "unable to create request")
        }

        res, err := httpClient(req)
        if err != nil {
            return nil, errors.Wrap(err, "http request failed")
        }
        if res == nil || res.StatusCode != 200 {
            return nil, fmt.Errorf("unexpected response: %+v", res)

        }
        return ioutil.ReadAll(res.Body)
    }
}
