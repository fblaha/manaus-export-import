package export

import (
	"github.com/fblaha/manaus-export-import/pool"
	"github.com/pkg/errors"
)

type idLoader interface {
	LoadIDs() ([]string, error)
}

type dataLoader interface {
	Load(id string) ([]byte, error)
}

type dataWriter interface {
	Write(id string, data []byte) error
}

// Exporter does export
type Exporter struct {
	idLoader
	dataLoader
	dataWriter

	dir      string
	url      string
	executor pool.Executor
}

// Execute executes export
func (e Exporter) Execute() error {
	ids, err := e.LoadIDs()
	if err != nil {
		return errors.Wrap(err, "unable to read ids to export")
	}
	// TODO concurrent execution
	for _, id := range ids {
		data, err := e.Load(id)
		if err != nil {
			return errors.Wrap(err, "unable to to load market: "+id)
		}
		err = e.Write(id, data)
		if err != nil {
			return errors.Wrap(err, "unable to to write market: "+id)
		}
	}
	// TODO create archive
	return nil

}
