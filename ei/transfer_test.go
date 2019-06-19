package ei

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/fblaha/manaus-export-import/storage"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
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

func TestTransferExecute(t *testing.T) {
	tempDir, _ := ioutil.TempDir("", "export")
	defer func() {
		require.NoError(t, os.RemoveAll(tempDir))
	}()
	writer := storage.NewDirectoryWriter(tempDir, ".txt")
	transfer := NewTransfer(mockIDLoader{}, mockDataLoader{}, writer)
	require.NoError(t, transfer.Execute())
}

func TestTransferExecuteErrorID(t *testing.T) {
	err := fmt.Errorf("id fetch error")
	idLoader := mockIDLoader{err}
	dataLoader := mockDataLoader{}
	dataWriter := mockDataWriter{}
	checkTransferror(idLoader, dataLoader, dataWriter, err, t)
}

func TestTransferExecuteErrorDataLoad(t *testing.T) {
	idLoader := mockIDLoader{}
	err := fmt.Errorf("data load error")
	dataLoader := mockDataLoader{err}
	dataWriter := mockDataWriter{}
	checkTransferror(idLoader, dataLoader, dataWriter, err, t)
}

func TestTransferExecuteErrorDataWrite(t *testing.T) {
	idLoader := mockIDLoader{}
	dataLoader := mockDataLoader{}
	err := fmt.Errorf("data write error")
	dataWriter := mockDataWriter{err}
	checkTransferror(idLoader, dataLoader, dataWriter, err, t)
}

func checkTransferror(idLoader mockIDLoader, dataLoader mockDataLoader, dataWriter mockDataWriter, expectedErr error, t *testing.T) {
	transfer := NewTransfer(idLoader, dataLoader, dataWriter)
	err := transfer.Execute()
	require.Error(t, err)
	require.Equal(t, expectedErr, errors.Cause(err))
}
