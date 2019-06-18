package rest

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMarketFetch(t *testing.T) {
	var calledURL string
	loader := newMarketLoader("http://mns/rest", func(url string) (bytes []byte, e error) {
		calledURL = url
		return []byte("{}"), nil
	})
	bytes, err := loader.load("1000")
	require.Equal(t, "http://mns/rest/markets/1000", calledURL)
	require.NoError(t, err)
	require.NotEmpty(t, bytes)
}
