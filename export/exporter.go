package export

import (
	"github.com/fblaha/manaus-export-import/pool"
	"github.com/pkg/errors"
)

type idLoader interface {
	LoadIDs() ([]string, error)
}

type marketLoader interface {
	LoadMarket(id string) ([]byte, error)
}

type marketWriter interface {
	WriteMarket(id string, data []byte) error
}

type Exporter struct {
	idLoader
	marketLoader
	marketWriter

	dir      string
	url      string
	executor pool.Executor
}

func (e Exporter) Run() error {
	ids, err := e.LoadIDs()
	if err != nil {
		return errors.Wrap(err, "unable to read ids to export")
	}
	for _, id := range ids {
		data, err := e.LoadMarket(id)
		if err != nil {
			return errors.Wrap(err, "unable to to load market: "+id)
		}
		err = e.WriteMarket(id, data)
		if err != nil {
			return errors.Wrap(err, "unable to to load market: "+id)
		}
	}
	return nil

}
