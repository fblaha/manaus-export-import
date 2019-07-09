package archive

import (
	"io/ioutil"
	"log"
	"path"
)

type layout struct {
	dir    string
	suffix string
}

// DataLoader capable of loading data by ID from archive
type DataLoader struct {
	layout
}

// NewDataLoader constructor
func NewDataLoader(dir string, suffix string) DataLoader {
	return DataLoader{layout: layout{dir: dir, suffix: suffix}}
}

// Load data from archive by ID
func (l DataLoader) Load(id string) ([]byte, error) {
	filePath := path.Join(l.dir, id+l.suffix)
	log.Println("loading data from :", filePath)
	return ioutil.ReadFile(filePath)
}
