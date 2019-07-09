package archive

import (
	"path"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWriter(t *testing.T) {
	tempDir, purge, err := CreateTempDir()
	defer purge()
	writer := NewWriter(tempDir, ".txt")
	err = writer.Write("100", []byte("{}"))
	require.NoError(t, err)
}

func TestDirectoryWriteFailure(t *testing.T) {
	tempDir, purge, err := CreateTempDir()
	defer purge()
	writer := NewWriter(tempDir, ".txt")
	purge()
	err = writer.Write("100", []byte("{}"))
	require.Error(t, err)
}

func TestWriterNotExist(t *testing.T) {
	require.Panics(t, func() {
		NewWriter("/path/does/not/exist", ".txt")
	})
}

func TestDirectoryArchive(t *testing.T) {
	tempDir, purge, err := CreateTempDir()
	defer purge()
	require.NoError(t, err)
	writer := NewWriter(tempDir, ".txt")
	err = writer.Write("100", []byte("{}"))
	require.NoError(t, err)

	archiveFile := path.Join(tempDir, "export.zip")
	err = writer.MakeArchive(archiveFile)
	require.NoError(t, err)
	require.FileExists(t, archiveFile)
}
