package ei

import (
	"github.com/fblaha/manaus-export-import/config"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/mholt/archiver"
)

func TestImport(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "POST", r.Method)
		bytes, err := ioutil.ReadAll(r.Body)
		require.NoError(t, err)
		require.Contains(t, string(bytes), "1000856224250015")

	}))
	defer ts.Close()

	tempDir, err := ioutil.TempDir("", "mns-test")
	defer os.RemoveAll(tempDir)
	require.NoError(t, err)
	content := filepath.Join("testdata", "1000856224250015.json")
	testArchive := filepath.Join(tempDir, "testArchive.zip")
	err = archiver.Archive([]string{content}, testArchive)
	require.NoError(t, err)

	Import(config.Conf{ArchiveFile: testArchive, URL: ts.URL, Concurrency: 10})

}
