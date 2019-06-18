package storage

import (
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/pkg/errors"
)

// NewDirectoryWriter constructor
func NewDirectoryWriter(dir string) DirectoryWriter {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Panic("dir does not exist: ", err)
	}
	return DirectoryWriter{dir: dir}
}

// DirectoryWriter writes to directory by id
type DirectoryWriter struct {
	dir string
}

// WriteMarket writes market to file, file name reflects market id
func (dw DirectoryWriter) WriteMarket(id string, data []byte) error {
	filePath := path.Join(dw.dir, id+".json")
	if err := ioutil.WriteFile(filePath, data, 0644); err != nil {
		return errors.Wrap(err, "unable to save market to file: "+filePath)
	}
	return nil
}
