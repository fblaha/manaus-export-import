package ei

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/fblaha/manaus-export-import/archive"
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
	tempDir, purge, err := archive.CreateTempDir()
	require.NoError(t, err)
	defer purge()
	writer := archive.NewWriter(tempDir, ".txt")
	transfer := NewTransfer(mockIDLoader{}, mockDataLoader{}, writer)
	require.NoError(t, transfer.Execute(10))
	infos, _ := ioutil.ReadDir(tempDir)
	require.Len(t, infos, 3)
	for i, info := range infos {
		require.Equal(t, fmt.Sprintf("%d.txt", i+1), info.Name())
	}
}

func TestTransferExecuteErrorID(t *testing.T) {
	err := fmt.Errorf("id fetch error")
	idLoader := mockIDLoader{err}
	dataLoader := mockDataLoader{}
	dataWriter := mockDataWriter{}
	checkTransferError(idLoader, dataLoader, dataWriter, err, t)
}

func TestTransferExecuteErrorDataLoad(t *testing.T) {
	idLoader := mockIDLoader{}
	err := fmt.Errorf("data load error")
	dataLoader := mockDataLoader{err}
	dataWriter := mockDataWriter{}
	checkTransferError(idLoader, dataLoader, dataWriter, err, t)
}

func TestTransferExecuteErrorDataWrite(t *testing.T) {
	idLoader := mockIDLoader{}
	dataLoader := mockDataLoader{}
	err := fmt.Errorf("data write error")
	dataWriter := mockDataWriter{err}
	checkTransferError(idLoader, dataLoader, dataWriter, err, t)
}

func checkTransferError(idLoader mockIDLoader, dataLoader mockDataLoader, dataWriter mockDataWriter, expectedErr error, t *testing.T) {
	transfer := NewTransfer(idLoader, dataLoader, dataWriter)
	err := transfer.Execute(10)
	require.Error(t, err)
	require.Equal(t, expectedErr, errors.Cause(err))
}
