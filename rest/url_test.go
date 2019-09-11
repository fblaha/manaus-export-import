package rest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadURL(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, "42")
		require.NoError(t, err)
	}))
	defer ts.Close()

	urlLoader := NewURLLoader(http.DefaultClient.Do)
	res, err := urlLoader(ts.URL)
	require.NoError(t, err)

	require.Equal(t, "42", string(res))
}
