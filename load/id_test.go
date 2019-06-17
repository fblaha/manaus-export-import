package load

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func newFakeIDLoader(data string, err error) idLoader {
	return newIDLoader("http://ignored", func(url string) (bytes []byte, e error) {
		return []byte(data), err
	})
}

func TestLoadIDs(t *testing.T) {
	loader := newFakeIDLoader(`[ "aaa" ]`, nil)
	ids, err := loader.load()
	require.NoError(t, err)
	require.NotEmpty(t, ids)
}

func TestLoadIDsParseError(t *testing.T) {
	loader := newFakeIDLoader(`[ "aaa" `, nil)
	ids, err := loader.load()
	require.Error(t, err)
	require.Contains(t, err.Error(), "unable to parse")
	require.Nil(t, ids)
}

func TestLoadIDsError(t *testing.T) {
	loader := newFakeIDLoader("", fmt.Errorf("some error"))
	ids, err := loader.load()
	require.Error(t, err)
	require.Contains(t, err.Error(), "some error")
	require.Nil(t, ids)
}
