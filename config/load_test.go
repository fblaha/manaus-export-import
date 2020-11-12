package config

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadConfigDefault(t *testing.T) {
	config := LoadConfig()
	expected := Conf{
		URL:         "http://manaus:8080/",
		Concurrency: runtime.NumCPU(),
		ArchiveFile: config.ArchiveFile,
	}
	require.Equal(t, config, expected)
	require.Equal(t, config.MarketIDsURL(), "http://manaus:8080/market-ids/")
	require.Equal(t, config.FootprintsURL(), "http://manaus:8080/footprints/")
}

func TestArchiveFile(t *testing.T) {
	file := archiveFile()
	require.Contains(t, file, "export")
	require.Contains(t, file, "zip")
}
