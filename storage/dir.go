package storage

import (
	"github.com/mholt/archiver"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/pkg/errors"
)

// NewDirectoryWriter constructor
func NewDirectoryWriter(dir string, suffix string) DirectoryWriter {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Panic("dir does not exist: ", err)
	}
	return DirectoryWriter{dir: dir, suffix: suffix}
}

// DirectoryWriter writes to directory by id
type DirectoryWriter struct {
	dir    string
	suffix string
}

// Write writes data to file, file name reflects id
func (dw DirectoryWriter) Write(id string, data []byte) error {
	filePath := path.Join(dw.dir, id+dw.suffix)
	log.Println("writing data to :", filePath)
	if err := ioutil.WriteFile(filePath, data, 0644); err != nil {
		return errors.Wrap(err, "unable to save data to file : "+filePath)
	}
	return nil
}

// MakeArchive makes an archive
func (dw DirectoryWriter) MakeArchive(file string) error {
	infos, err := ioutil.ReadDir(dw.dir)
	if err != nil {
		return errors.Wrap(err, "unable to read directory : "+dw.dir)
	}
	var files []string
	for _, info := range infos {
		files = append(files, path.Join(dw.dir, info.Name()))
	}
	return archiver.Archive(files, file)
}

// Purge purges tmp dir
func (dw DirectoryWriter) Purge() error {
	return os.RemoveAll(dw.dir)
}
