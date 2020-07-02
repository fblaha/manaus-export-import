package rest

import (
	"encoding/json"
	"fmt"
	"log"
)

// IDLoader capable of loading IDs from URL
type IDLoader struct {
	url    string
	loader URLLoader
}

// NewIDLoader constructor
func NewIDLoader(url string, loader URLLoader) IDLoader {
	return IDLoader{url: url, loader: loader}
}

// LoadIDs loads list of IDs from URL
func (p IDLoader) LoadIDs() ([]string, error) {
	log.Println("loading IDs from:", p.url)
	bytes, err := p.loader(p.url)
	if err != nil {
		return nil, fmt.Errorf("unable to load IDs: %v", err)
	}
	var ids []string
	err = json.Unmarshal(bytes, &ids)
	if err != nil {
		return nil, fmt.Errorf("unable to parse IDs: %v", err)
	}
	return ids, nil
}
