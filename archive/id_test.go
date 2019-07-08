package archive

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadIDs(t *testing.T) {
	path := filepath.Join("testdata", "archive")
	loader := NewIDLoader(path, ".json")
	ids, err := loader.LoadIDs()
	require.NoError(t, err)
	require.Len(t, ids, 3)
	require.Contains(t, ids, "1000856054200016")
}

func TestLoadIDsError(t *testing.T) {
	path := filepath.Join("testdata", "no-such-dir")
	loader := NewIDLoader(path, ".json")
	ids, err := loader.LoadIDs()
	require.Error(t, err)
	require.Nil(t, ids)
}
