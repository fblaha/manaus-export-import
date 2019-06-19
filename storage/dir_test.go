package storage

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDirectoryWriter(t *testing.T) {
	tempDir, err := ioutil.TempDir(".", "export")
	defer func() {
		require.NoError(t, os.RemoveAll(tempDir))
	}()
	writer := NewDirectoryWriter(tempDir, ".txt")
	err = writer.Write("100", []byte("{}"))
	require.NoError(t, err)
}

func TestDirectoryWriteFailure(t *testing.T) {
	tempDir, err := ioutil.TempDir(".", "export")
	writer := NewDirectoryWriter(tempDir, ".txt")
	require.NoError(t, writer.Purge())
	err = writer.Write("100", []byte("{}"))
	require.Error(t, err)
}

func TestDirectoryWriterNotExist(t *testing.T) {
	require.Panics(t, func() {
		NewDirectoryWriter("/path/does/not/exist", ".txt")
	})
}

func TestDirectoryArchive(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "export")
	require.NoError(t, err)
	defer func() {
		require.NoError(t, os.RemoveAll(tempDir))
	}()
	writer := NewDirectoryWriter(tempDir, ".txt")
	err = writer.Write("100", []byte("{}"))
	require.NoError(t, err)

	archiveFile := path.Join(tempDir, "export.zip")
	err = writer.MakeArchive(archiveFile)
	require.NoError(t, err)
	require.FileExists(t, archiveFile)
}

func TestDirectoryPurge(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "export")
	require.NoError(t, err)
	writer := NewDirectoryWriter(tempDir, ".txt")
	require.DirExists(t, tempDir)
	require.NoError(t, writer.Purge())
}
