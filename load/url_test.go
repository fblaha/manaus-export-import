package load

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoadURL(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, "42")
		require.NoError(t, err)
	}))
	defer ts.Close()

	res, err := loadURL(ts.URL)
	require.NoError(t, err)

	require.Equal(t, "42", string(res))
}
