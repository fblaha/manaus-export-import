package storage

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"os"
	"path"
)

type DirectoryWriter struct {
	dir string
}

func NewDirectoryWriter(dir string) DirectoryWriter {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Panic("dir does not exist: ", err)
	}
	return DirectoryWriter{dir: dir}
}

func (dw DirectoryWriter) WriteMarket(id string, data []byte) error {
	filePath := path.Join(dw.dir, id+".json")
	if err := ioutil.WriteFile(filePath, data, 0644); err != nil {
		return errors.Wrap(err, "unable to save market to file: "+filePath)
	}
	return nil
}
