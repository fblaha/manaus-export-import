package archive

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadData(t *testing.T) {
	path := filepath.Join("testdata", "archive")
	loader := NewDataLoader(path, ".json")
	data, err := loader.Load("1000856054200016")
	require.NoError(t, err)
	require.Contains(t, string(data), "1000856054200016")
}
func TestLoadDataError(t *testing.T) {
	path := filepath.Join("testdata", "archive")
	loader := NewDataLoader(path, ".json")
	data, err := loader.Load("no-such-id")
	require.Error(t, err)
	require.Nil(t, data)
}
