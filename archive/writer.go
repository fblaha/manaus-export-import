package archive

import (
	"github.com/mholt/archiver"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/pkg/errors"
)

// NewWriter constructor
func NewWriter(dir string, suffix string) Writer {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Panic("dir does not exist: ", err)
	}
	return Writer{dir: dir, suffix: suffix}
}

// Writer writes to directory by id
type Writer struct {
	dir    string
	suffix string
}

// Write writes data to file, file name reflects id
func (w Writer) Write(id string, data []byte) error {
	filePath := path.Join(w.dir, id+w.suffix)
	log.Println("writing data to :", filePath)
	if err := ioutil.WriteFile(filePath, data, 0644); err != nil {
		return errors.Wrap(err, "unable to save data to file : "+filePath)
	}
	return nil
}

// MakeArchive makes an archive
func (w Writer) MakeArchive(file string) error {
	infos, err := ioutil.ReadDir(w.dir)
	if err != nil {
		return errors.Wrap(err, "unable to read directory : "+w.dir)
	}
	var files []string
	for _, info := range infos {
		files = append(files, path.Join(w.dir, info.Name()))
	}
	return archiver.Archive(files, file)
}

// Purge purges tmp dir
func (w Writer) Purge() error {
	return os.RemoveAll(w.dir)
}
