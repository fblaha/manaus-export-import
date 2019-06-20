package config

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadConfigDefault(t *testing.T) {
	config := LoadConfig()
	expected := Conf{
		URL:         "http://localhost:7777",
		Concurrency: runtime.NumCPU(),
		ArchiveFile: config.ArchiveFile,
	}
	require.Equal(t, config, expected)
}

func TestArchiveFile(t *testing.T) {
	file := archiveFile()
	require.Contains(t, file, "export")
	require.Contains(t, file, "zip")
}
