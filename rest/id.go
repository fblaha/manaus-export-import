package rest

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type idLoader struct {
	url    string
	loader urlLoader
}

func newIDLoader(url string, loader urlLoader) idLoader {
	return idLoader{url: url, loader: loader}
}

func (p idLoader) LoadIDs() ([]string, error) {
	bytes, err := p.loader(p.url)
	if err != nil {
		return nil, errors.Wrap(err, "unable to load IDs")
	}
	var ids []string
	err = json.Unmarshal(bytes, &ids)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse IDs")
	}
	return ids, nil
}
