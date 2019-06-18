package storage

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDirectoryWriter(t *testing.T) {
	tempDir, err := ioutil.TempDir(".", "export")
	defer func() {
		require.NoError(t, os.RemoveAll(tempDir))
	}()
	writer := NewDirectoryWriter(tempDir)
	err = writer.Write("100", []byte("{}"))
	require.NoError(t, err)
}

func TestDirectoryWriterNotExist(t *testing.T) {
	require.Panics(t, func() {
		NewDirectoryWriter("/path/does/not/exist")
	})
}
