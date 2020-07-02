package archive

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// IDLoader capable of loading IDs from files names in the directory
type IDLoader struct {
	layout
}

// NewIDLoader constructor
func NewIDLoader(dir string, suffix string) IDLoader {
	return IDLoader{layout: layout{dir: dir, suffix: suffix}}
}

// LoadIDs extracts IDs from file names in the directory
func (l IDLoader) LoadIDs() ([]string, error) {
	log.Println("loading IDs from:", l.dir)
	files, err := ioutil.ReadDir(l.dir)
	if err != nil {
		return nil, fmt.Errorf("unable to read dir : %s %v", l.dir, err)
	}
	var ids []string
	for _, file := range files {
		fileName := file.Name()
		if strings.HasSuffix(fileName, l.suffix) {
			id := strings.TrimSuffix(fileName, l.suffix)
			ids = append(ids, id)
		}
	}
	return ids, nil
}
