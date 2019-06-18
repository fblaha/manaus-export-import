package rest

import (
	"encoding/json"
	"github.com/pkg/errors"
)

type idLoader struct {
	mnsURL string
	loader urlLoader
}

func newIDLoader(mnsURL string, loader urlLoader) idLoader {
	return idLoader{mnsURL: mnsURL, loader: loader}
}

func (p idLoader) load() ([]string, error) {
	bytes, err := p.loader(p.mnsURL + "/market-ids")
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
