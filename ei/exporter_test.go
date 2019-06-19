package ei

import (
	"fmt"
	"github.com/fblaha/manaus-export-import/storage"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"testing"
)

type mockIDLoader struct {
	err error
}

func (m mockIDLoader) LoadIDs() ([]string, error) {
	return []string{"1", "2", "3"}, m.err
}

type mockDataLoader struct {
	err error
}

func (m mockDataLoader) Load(id string) ([]byte, error) {
	return []byte(id), m.err
}

type mockDataWriter struct {
	err error
}

func (m mockDataWriter) Write(id string, data []byte) error {
	return m.err
}

func TestExporterExecute(t *testing.T) {
	tempDir, _ := ioutil.TempDir("", "export")
	defer func() {
		require.NoError(t, os.RemoveAll(tempDir))
	}()
	writer := storage.NewDirectoryWriter(tempDir, ".txt")
	exporter := NewExporter(mockIDLoader{}, mockDataLoader{}, writer)
	require.NoError(t, exporter.Execute())
}

func TestExporterExecuteErrorID(t *testing.T) {
	err := fmt.Errorf("id fetch error")
	idLoader := mockIDLoader{err}
	dataLoader := mockDataLoader{}
	dataWriter := mockDataWriter{}
	checkExportError(idLoader, dataLoader, dataWriter, err, t)
}

func TestExporterExecuteErrorDataLoad(t *testing.T) {
	idLoader := mockIDLoader{}
	err := fmt.Errorf("data load error")
	dataLoader := mockDataLoader{err}
	dataWriter := mockDataWriter{}
	checkExportError(idLoader, dataLoader, dataWriter, err, t)
}

func TestExporterExecuteErrorDataWrite(t *testing.T) {
	idLoader := mockIDLoader{}
	dataLoader := mockDataLoader{}
	err := fmt.Errorf("data write error")
	dataWriter := mockDataWriter{err}
	checkExportError(idLoader, dataLoader, dataWriter, err, t)
}

func checkExportError(idLoader mockIDLoader, dataLoader mockDataLoader, dataWriter mockDataWriter, expectedErr error, t *testing.T) {
	exporter := NewExporter(idLoader, dataLoader, dataWriter)
	err := exporter.Execute()
	require.Error(t, err)
	require.Equal(t, expectedErr, errors.Cause(err))
}
