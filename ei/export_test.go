package ei

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/fblaha/manaus-export-import/config"
	"github.com/stretchr/testify/require"
)

func TestConfigureExport(t *testing.T) {
	loadConfig := config.LoadConfig()
	_, purge, err := configureExport(loadConfig)
	require.NoError(t, err)
	defer purge()
}

func TestExport(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, `[ "1", "2", "3" ]`)
		require.NoError(t, err)
		require.Contains(t, "/market-ids/ /footprints/1 /footprints/2 /footprints/3", r.URL.String())
	}))
	defer ts.Close()

	archive := fmt.Sprintf("export%d.zip", time.Now().Nanosecond())
	defer func() {
		require.NoError(t, os.RemoveAll(archive))
	}()
	conf := config.Conf{
		URL:         ts.URL,
		ArchiveFile: archive,
		Concurrency: 10,
	}
	Export(conf)
}
