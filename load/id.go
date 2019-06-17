package load

import (
	"encoding/json"
	"github.com/pkg/errors"
)

type idLoader struct {
	mnsUrl string
	loader urlLoader
}

func newIDLoader(mnsUrl string, loader urlLoader) idLoader {
	return idLoader{mnsUrl: mnsUrl, loader: loader}
}

func (p idLoader) load() ([]string, error) {
	bytes, err := p.loader(p.mnsUrl + "/market-ids")
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
