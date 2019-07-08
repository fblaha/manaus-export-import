package archive

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/pkg/errors"
)

// IDLoader capable of loading from directory
type IDLoader struct {
	dir    string
	suffix string
}

// NewIDLoader constructor
func NewIDLoader(dir string, suffix string) IDLoader {
	return IDLoader{dir: dir, suffix: suffix}
}

// LoadIDs loads list of IDs from URL
func (l IDLoader) LoadIDs() ([]string, error) {
	log.Println("loading IDs from :", l.dir)
	files, err := ioutil.ReadDir(l.dir)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read dir : "+l.dir)
	}
	var ids []string
	for _, file := range files {
		id := strings.TrimSuffix(file.Name(), l.suffix)
		ids = append(ids, id)
	}
	return ids, nil
}
