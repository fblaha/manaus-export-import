package ei

import (
	"github.com/fblaha/manaus-export-import/pool"
	"github.com/pkg/errors"
)

// Exporter does export
type Exporter struct {
	IDLoader
	DataLoader
	DataWriter

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
