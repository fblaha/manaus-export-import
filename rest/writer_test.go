package rest

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWriter(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bytes, _ := ioutil.ReadAll(r.Body)
		data := string(bytes)
		require.Equal(t, "42", data)
	}))
	defer ts.Close()

	writer := NewWriter(http.DefaultClient.Do, ts.URL, "text/plain")
	err := writer.Write("ignored", []byte("42"))
	require.NoError(t, err)
}

func TestWriterError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer ts.Close()

	writer := NewWriter(http.DefaultClient.Do, ts.URL, "text/plain")
	err := writer.Write("1000", []byte("42"))
	require.Error(t, err)
}
