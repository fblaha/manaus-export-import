package archive

import (
	"github.com/mholt/archiver"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtractToTmpDir(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "mns-test")
	defer os.RemoveAll(tempDir)
	require.NoError(t, err)
	content := filepath.Join("testdata", "archive", "1000856054200016.json")
	testArchive := filepath.Join(tempDir, "testArchive.zip")
	err = archiver.Archive([]string{content}, testArchive)
	require.NoError(t, err)

	tmpDir, del, err := ExtractToTmpDir(testArchive)
	defer del()
	require.NoError(t, err)
	require.DirExists(t, tmpDir)
	require.FileExists(t, filepath.Join(tmpDir, "1000856054200016.json"))
}

func TestExtractToTmpDirError(t *testing.T) {
	tmpDir, del, err := ExtractToTmpDir("no-such-archive.zip")
	defer del()
	require.DirExists(t, tmpDir)
	require.NotZero(t, tmpDir)
	require.Error(t, err)
}
