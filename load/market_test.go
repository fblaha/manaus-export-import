package load

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMarketFetch(t *testing.T) {
	var called bool
	loader := newMarketLoader("http://mns/rest", func(url string) (bytes []byte, e error) {
		require.Equal(t, "http://mns/rest/markets/1000" ,url)
		called = true
		return []byte("{}"), nil
	})
	bytes, err := loader.load("1000")
	require.NoError(t, err)
	require.True(t, called)
	require.NotEmpty(t, bytes)
}

