package rest

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func mockClient(mockStatusCode int, mockBody string, mockErr error) httpClient {
	return func(req *http.Request) (*http.Response, error) {
		response := http.Response{
			StatusCode: mockStatusCode,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(mockBody))),
		}
		return &response, mockErr
	}
}

func TestLoadURL(t *testing.T) {
	urlLoader := NewURLLoader(mockClient(200, "42", nil))
	res, err := urlLoader("http://ignored")
	require.NoError(t, err)
	require.Equal(t, "42", string(res))
}

func TestLoadURLError(t *testing.T) {
	urlLoader := NewURLLoader(mockClient(200, "42", fmt.Errorf("some error")))
	res, err := urlLoader("http://ignored")
	require.Error(t, err)
	require.Nil(t, res)
}

func TestLoadURLErrorStatus(t *testing.T) {
	urlLoader := NewURLLoader(mockClient(400, "42", nil))
	res, err := urlLoader("http://ignored")
	require.Error(t, err)
	require.Nil(t, res)
}
