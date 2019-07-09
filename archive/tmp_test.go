package archive

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDirectoryPurge(t *testing.T) {
	tempDir, purge, err := CreateTempDir()
	require.NoError(t, err)
	require.DirExists(t, tempDir)
	purge()
}
